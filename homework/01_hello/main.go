package main

import (
	"fmt"
	"os"
)

/*
第 01 次作业：Hello 与程序入口（进阶版）

完成要求（必做）：
1) 打印三行内容：姓名、学习目标、今天日期（日期先写死字符串）。
2) 用 fmt.Printf 打印一句完整介绍，包含姓名和当前章节编号。
3) 在注释里写出 go run . 与 go build 的区别（至少两点）。

加分挑战（选做）：
- 用 os.Args 读取命令行参数，把姓名改成可传入（默认值可自定义）。
*/

func main() {
	// TODO: 在这里完成本次作业
	fmt.Println("TODO: 完成 01_hello 作业")
	fmt.Println("姓名：wjj")
	fmt.Println("学习目标：学习go语言")
	fmt.Println("今天日期：2026-05-18")
	fmt.Printf("我是%s，今天学习了第%d章\n", "wjj", 1)
	fmt.Println("go run . 与 go build 的区别：")
	fmt.Println("go run . 是运行程序，go build 是编译程序")
	fmt.Println("go run . 是运行程序，go build 是编译程序")
	name := os.Args[1]
	fmt.Println("name: ", name)

}
