package main

import "fmt"

/*
第 3 课：close + range

关键知识：
1) 通常由“发送方”关闭 channel：close(ch)。
2) 接收方可以用 for v := range ch 持续读取，直到 channel 关闭且数据读空。
3) 关闭后不能再发送，否则 panic。
*/

func producer(out chan<- int) {
	defer close(out) // 发送方负责关闭

	for i := 1; i <= 5; i++ {
		out <- i * 10
	}
}

func main() {
	fmt.Println("=== 03 close and range ===")

	ch := make(chan int)
	go producer(ch)

	// range 会自动在 channel 关闭后退出循环
	for v := range ch {
		fmt.Println("[main] 收到:", v)
	}

	fmt.Println("channel 已关闭，循环结束")
}

