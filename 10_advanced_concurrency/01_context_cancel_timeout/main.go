package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
第 10 章（高级并发）- 01：context 取消与超时

你需要掌握的点：
1) 父协程创建 context，然后传给所有子协程。
2) 子协程在循环里监听 <-ctx.Done()，一旦取消就尽快退出。
3) 用 WaitGroup 保证主协程等待所有子协程“收尾”完成。

为什么这很重要：
- 真实项目里，HTTP 请求超时、用户取消、服务关闭都需要“整条链路一起停”。
- 只靠 goroutine 不够，必须有“可取消协议”，context 就是这个协议。
*/

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	// ticker 模拟“周期任务”
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			// context 取消时，这里会被触发。
			// 任何资源释放动作（关闭连接/写日志）都应该放在退出前做。
			fmt.Printf("[worker-%d] 收到取消信号，原因: %v，准备退出\n", id, ctx.Err())
			return
		case t := <-ticker.C:
			fmt.Printf("[worker-%d] 正在工作，tick=%s\n", id, t.Format("15:04:05.000"))
		}
	}
}

func main() {
	fmt.Println("=== 01 context 取消与超时 ===")

	// WithTimeout: 到时间后自动 cancel
	ctx, cancel := context.WithTimeout(context.Background(), 1200*time.Millisecond)
	defer cancel() // 好习惯：函数退出前确保释放

	var wg sync.WaitGroup
	workerCount := 3
	wg.Add(workerCount)

	for i := 1; i <= workerCount; i++ {
		go worker(ctx, &wg, i)
	}

	// 等待所有 worker 优雅退出
	wg.Wait()
	fmt.Println("所有 worker 已停止，主程序退出")
}
