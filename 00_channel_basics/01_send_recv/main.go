package main

import (
	"fmt"
	"time"
)

/*
第 1 课：channel 发送与接收（无缓冲）

关键知识：
1) ch := make(chan int) 创建无缓冲 channel。
2) ch <- v 发送；v := <-ch 接收。
3) 无缓冲 channel 的发送和接收必须“配对”才能继续执行。
*/

func worker(ch chan int) {
	time.Sleep(3000 * time.Millisecond)
	for v := range ch {

		fmt.Println("received:", v)
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("worker done")
}

func main() {
	fmt.Println("=== 01 send / recv (unbuffered) ===")

	ch := make(chan int, 2)

	go worker(ch)

	// 接收方放到 goroutine 里，模拟“另一个协程消费消息”。
	// go func() {
	// 	// 睡眠只是为了更容易观察输出顺序，不是必须。
	// 	time.Sleep(300 * time.Millisecond)
	// 	v := <-ch
	// 	fmt.Println("[receiver] 收到:", v)
	// }()

	fmt.Println("[main] 准备发送 42")
	ch <- 42 // 这里会阻塞，直到上面的 goroutine 开始接收
	fmt.Println("[main] 准备发送 43")
	ch <- 43
	fmt.Println("[main] 准备发送 44")
	ch <- 44

	close(ch)
	time.Sleep(2500 * time.Millisecond)
	fmt.Println("[main] 发送完成")
}
