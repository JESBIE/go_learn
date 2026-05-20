package main

import (
	"fmt"
	"sync"
)

// 练习1：并发计数器
// 目标：100 个 goroutine 各加 1，最终应该等于 100。
// 这里给出两个版本：先看不安全版，再看加锁后的安全版。

func unsafeCounter() {
	var wg sync.WaitGroup
	counter := 0

	workers := 100
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			// 并发读写同一个变量，但没有加锁。
			// 这里会有 data race（数据竞争）。
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("[不安全版] counter =", counter)
}

func safeCounter() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0

	workers := 100
	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func() {
			defer wg.Done()

			// 临界区加锁：同一时刻只允许一个 goroutine 修改 counter。
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("[安全版]   counter =", counter)
}

func main() {
	fmt.Println("===== 并发计数器示例 =====")
	unsafeCounter()
	safeCounter()

	// 运行建议：
	// 1) go run .
	// 2) go run -race .  // 观察不安全版的 data race 报告
}
