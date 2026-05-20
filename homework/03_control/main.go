package main

import "fmt"

/*
第 03 次作业：流程控制（进阶版）

完成要求（必做）：
1) 写一个函数 gradeLevel(score int) string：
   - 90~100: A
   - 80~89: B
   - 70~79: C
   - 60~69: D
   - 其他: F
2) 用 for 循环计算 1~100 的奇数和并输出。
3) 用 switch 把数字 1~7 转成星期几，非法输入返回"无效日期"。

加分挑战（选做）：
- 打印 1~100 里的 FizzBuzz（3 的倍数打印 Fizz，5 的倍数打印 Buzz）。
*/

func gradeLevel(score int) string {
	// TODO: 实现评分规则
	return "TODO"
}

func weekday(n int) string {

	return "TODO"
}

func main() {
	fmt.Println("TODO: 完成 03_control 作业")

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
