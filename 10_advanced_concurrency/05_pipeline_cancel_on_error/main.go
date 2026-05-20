package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

/*
第 10 章（高级并发）- 05：Pipeline + 首错取消（非常常用）

目标：
1) 搭建 3 段流水线：source -> transform -> sink。
2) 任意阶段出现错误时，取消整个流水线，避免 goroutine 泄漏。
3) 所有 channel 发送都要配合 ctx.Done()，保证可以被中断。

为什么它是“高级并发”的关键：
- 大多数真实任务（爬虫、ETL、日志处理、消息处理）都是流水线。
- 只要一段挂了，不取消上下游就会出现阻塞和资源浪费。
*/

func source(ctx context.Context, total int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := 1; i <= total; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("[source] 收到取消，停止生产")
				return
			case out <- i:
				time.Sleep(80 * time.Millisecond)
			}
		}
	}()
	return out
}

func transform(ctx context.Context, cancel context.CancelFunc, in <-chan int, errCh chan<- error) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for n := range in {
			// 模拟：遇到特定数据时报错
			if n == 7 {
				reportErr(errCh, fmt.Errorf("transform 失败: 遇到非法数据 n=%d", n))
				cancel() // 触发全链路取消
				return
			}

			msg := fmt.Sprintf("item-%02d", n*n)
			select {
			case <-ctx.Done():
				fmt.Println("[transform] 收到取消，停止处理")
				return
			case out <- msg:
			}
		}
	}()
	return out
}

func sink(ctx context.Context, in <-chan string) []string {
	var results []string
	for {
		select {
		case <-ctx.Done():
			fmt.Println("[sink] 收到取消，提前结束收集")
			return results
		case item, ok := <-in:
			if !ok {
				return results
			}
			results = append(results, item)
		}
	}
}

func reportErr(errCh chan<- error, err error) {
	// 只保留第一个错误，避免多个 goroutine 同时写 errCh 造成阻塞。
	select {
	case errCh <- err:
	default:
	}
}

func main() {
	fmt.Println("=== 05 pipeline cancel on error ===")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	in := source(ctx, 12)
	out := transform(ctx, cancel, in, errCh)
	results := sink(ctx, out)

	// 查看是否有错误
	select {
	case err := <-errCh:
		if !errors.Is(err, context.Canceled) {
			fmt.Println("流水线失败:", err)
		}
	default:
		fmt.Println("流水线完成，无错误")
	}

	fmt.Println("已收集结果条数:", len(results))
	fmt.Println("结果内容:", results)
}

