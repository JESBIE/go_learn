package main

import "fmt"

/*
第 2 课：有缓冲 channel

关键知识：
1) make(chan T, n) 创建容量为 n 的缓冲队列。
2) 队列没满时，发送不会立刻阻塞。
3) 队列满时，再发送会阻塞；队列空时，接收会阻塞。
*/

func main() {
	fmt.Println("=== 02 buffered channel ===")

	ch := make(chan string, 2)
	fmt.Printf("capacity=%d len=%d\n", cap(ch), len(ch))

	ch <- "task-1"
	ch <- "task-2"
	fmt.Printf("发送两个后 len=%d\n", len(ch))

	// 第三个发送会阻塞（因为容量只有 2），所以先接收一个再发。
	// fmt.Println("开始接收:", <-ch)
	ch <- "task-3"

	for i := 0; i < 2; i++ {
		fmt.Println("接收:", <-ch)
	}
}
