package main

import (
	"fmt"
	"time"
)

/*
第 4 课：select + timeout

关键知识：
1) select 可以同时等待多个 channel 事件。
2) time.After(d) 会在 d 后返回一个可接收的 channel，用于超时控制。
3) 这是写“避免永久阻塞”代码的基础手法。
*/

func slowOperation() <-chan string {
	out := make(chan string, 1)
	go func() {
		// 模拟慢任务：1 秒后才有结果
		time.Sleep(1 * time.Second)
		out <- "任务完成"
	}()
	return out
}

func main() {
	fmt.Println("=== 04 select with timeout ===")

	resultCh := slowOperation()

	select {
	case res := <-resultCh:
		fmt.Println("收到结果:", res)
	case <-time.After(50000 * time.Millisecond):
		fmt.Println("超时：任务太慢，先返回")
	}
}
