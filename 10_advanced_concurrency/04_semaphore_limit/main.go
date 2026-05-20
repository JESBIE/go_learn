package main

import (
	"fmt"
	"sync"
	"time"
)

/*
第 10 章（高级并发）- 04：Semaphore（并发限流）

需求场景：
- 你有很多任务，但一次最多只允许 N 个并发执行（比如最多同时请求 3 个外部 API）。

实现方式（无第三方库）：
- 用“带缓冲 channel”当信号量：
  sem := make(chan struct{}, N)
  进入任务前写入 sem（占用一个配额）
  任务结束后读出 sem（释放一个配额）
*/

func task(id int, sem chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	// 获取配额：当通道满了会阻塞，达到限流效果
	sem <- struct{}{}
	fmt.Printf("[task-%d] 开始执行（当前并发<=%d）\n", id, cap(sem))

	// 模拟任务耗时
	time.Sleep(300 * time.Millisecond)
	fmt.Printf("[task-%d] 执行完成\n", id)

	// 释放配额
	<-sem
}

func main() {
	fmt.Println("=== 04 semaphore 并发限流 ===")

	totalTasks := 10
	maxConcurrent := 3

	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	wg.Add(totalTasks)

	for i := 1; i <= totalTasks; i++ {
		go task(i, sem, &wg)
	}

	wg.Wait()
	fmt.Println("全部任务完成")
}

