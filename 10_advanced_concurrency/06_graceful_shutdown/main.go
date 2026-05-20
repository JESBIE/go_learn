package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

/*
第 10 章（高级并发）- 06：优雅停机（graceful shutdown）

目标：
1) 用 signal.NotifyContext 监听退出信号（SIGINT / SIGTERM）。
2) 收到信号后，停止接新任务，但让已在执行中的任务完成。
3) 用 WaitGroup 等待 worker 收尾，避免“暴力退出”丢数据。

为什么重要：
- 在线服务关闭时，最怕“正在处理的数据丢失”。
- 优雅停机是高并发服务的基本功。
*/

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("[worker-%d] 开始处理 job=%d\n", id, job)
		time.Sleep(250 * time.Millisecond) // 模拟任务耗时
		fmt.Printf("[worker-%d] 完成 job=%d\n", id, job)
	}
	fmt.Printf("[worker-%d] jobs 关闭，退出\n", id)
}

func main() {
	fmt.Println("=== 06 graceful shutdown ===")

	// 监听 Ctrl+C / 终止信号
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	jobs := make(chan int, 8)
	var wg sync.WaitGroup

	// 启动 worker 池
	workerCount := 3
	wg.Add(workerCount)
	for i := 1; i <= workerCount; i++ {
		go worker(i, jobs, &wg)
	}

	// 演示方便：2 秒后向自身发送一次中断信号（相当于你按 Ctrl+C）
	go func() {
		time.Sleep(2 * time.Second)
		_ = syscall.Kill(os.Getpid(), syscall.SIGINT)
	}()

	// 生产任务：一旦收到退出信号，停止继续投递
produceLoop:
	for j := 1; j <= 30; j++ {
		select {
		case <-ctx.Done():
			fmt.Println("[main] 收到退出信号，停止接收新任务")
			break produceLoop
		case jobs <- j:
		}
	}

	close(jobs) // 关键：不再接新任务，通知 worker 读完退出
	wg.Wait()   // 关键：等待 worker 把已接任务做完
	fmt.Println("[main] 所有 worker 收尾完成，程序退出")
}

