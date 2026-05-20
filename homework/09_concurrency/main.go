package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
第 09 次作业：并发（进阶版）

完成要求（必做）：
1) 使用 WaitGroup 启动 10 个 goroutine，打印 worker 编号并等待全部结束。
2) 写一个共享计数器，分别实现：
   - 无锁版本（观察可能的问题）
   - Mutex 加锁版本（结果稳定）
3) 运行 go run -race .，理解 data race 报告。

4) 使用 context.WithTimeout 实现超时任务：
   - 任务每 100ms 打印一次进度
   - 超时后应退出并打印取消原因

加分挑战（选做）：
- 用 channel 改写为 3 个 worker 消费 10 个 job。
*/

func runWorkers() {
	// TODO: WaitGroup + 10 goroutine
}

func unsafeCounter() int {
	// TODO: 故意不加锁，观察并发问题
	return 0
}

func safeCounter() int {
	// TODO: 使用 Mutex 修复并发问题
	return 0
}

func runWithTimeout() {
	// TODO: 用 context.WithTimeout 控制任务超时退出
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	_ = ctx
}

func main() {
	_ = sync.Mutex{}
	fmt.Println("TODO: 完成 09_concurrency 作业")
}

