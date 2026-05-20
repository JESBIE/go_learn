package main

import (
	"fmt"
	"sync"
)

/*
第 10 章（高级并发）- 03：Fan-out / Fan-in

模式解释：
- Fan-out：把一个输入流分发给多个 worker 并行处理。
- Fan-in：把多个 worker 的输出合并回一个输出流。

这比“单个 worker”吞吐量更高，常用于流水线处理。
*/

func producer(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
		}
	}()
	return out
}

func squareWorker(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			result := n * n
			fmt.Printf("[worker-%d] %d^2=%d\n", id, n, result)
			out <- result
		}
	}()
	return out
}

// merge 把多个输入 channel 合并成一个输出 channel（fan-in）
func merge(cs ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			out <- n
		}
	}

	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	fmt.Println("=== 03 fan-out / fan-in ===")

	// 输入流
	in := producer(1, 2, 3, 4, 5, 6, 7, 8)

	// fan-out: 这里把同一个 in 交给多个 worker 并发消费。
	// 注意：同一个值只会被其中一个 worker 取到（竞争消费）。
	w1 := squareWorker(1, in)
	w2 := squareWorker(2, in)
	w3 := squareWorker(3, in)

	// fan-in: 合并多个 worker 输出
	out := merge(w1, w2, w3)

	// 消费最终输出
	for v := range out {
		fmt.Println("[main] 收到:", v)
	}

	fmt.Println("流水线结束")
}

