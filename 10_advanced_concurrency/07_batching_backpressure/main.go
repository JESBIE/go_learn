package main

import (
	"context"
	"fmt"
	"time"
)

/*
第 10 章（高级并发）- 07：批处理（batching）与降压

目标：
1) 高频小消息输入时，不要每条都立刻落库/发请求。
2) 用“按条数 or 按时间”两种触发条件聚合为批次。
3) 通过批处理降低下游系统压力，提高吞吐。

场景：
- 日志批量写入、埋点上报、消息批量持久化。
*/

func producer(ctx context.Context, total int, interval time.Duration) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= total; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
				time.Sleep(interval)
			}
		}
	}()
	return out
}

func batcher(ctx context.Context, in <-chan int, batchSize int, flushEvery time.Duration) <-chan []int {
	out := make(chan []int)
	go func() {
		defer close(out)

		ticker := time.NewTicker(flushEvery)
		defer ticker.Stop()

		batch := make([]int, 0, batchSize)
		flush := func() bool {
			if len(batch) == 0 {
				return true
			}
			// 深拷贝一份再发送，避免后续复用底层数组导致数据被改写。
			copyBatch := append([]int(nil), batch...)
			select {
			case <-ctx.Done():
				return false
			case out <- copyBatch:
				batch = batch[:0]
				return true
			}
		}

		for {
			select {
			case <-ctx.Done():
				_ = flush() // 尽量把最后残留的数据冲掉
				return
			case v, ok := <-in:
				if !ok {
					_ = flush()
					return
				}
				batch = append(batch, v)
				if len(batch) >= batchSize {
					if !flush() {
						return
					}
				}
			case <-ticker.C:
				if !flush() {
					return
				}
			}
		}
	}()
	return out
}

func consumeBatch(batch []int) {
	// 模拟批量写入外部系统
	time.Sleep(120 * time.Millisecond)
	fmt.Printf("[consumer] 批量写入 count=%d data=%v\n", len(batch), batch)
}

func main() {
	fmt.Println("=== 07 batching backpressure ===")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	in := producer(ctx, 25, 40*time.Millisecond)
	batches := batcher(ctx, in, 6, 250*time.Millisecond)

	for b := range batches {
		consumeBatch(b)
	}

	fmt.Println("批处理流程结束")
}

