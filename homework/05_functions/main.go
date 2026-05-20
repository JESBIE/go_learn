package main

import (
	"errors"
	"fmt"
)

/*
第 05 次作业：函数（进阶版）

完成要求（必做）：
1) 写函数 calc(a, b int, op string) (int, error)：
   - op 支持 + - * /
   - / 时处理 b == 0 的错误
2) 写可变参数函数 sum(nums ...int) int。
3) 写函数 splitTotal(total int) (partA int, partB int)，
   规则：partA 占 30%，partB 占 70%（整数即可）。

加分挑战（选做）：
- 写一个高阶函数 repeat(n int, fn func(int))，循环调用 fn。
*/

func calc(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("不支持的运算符")
	}
}

func sum(nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

func splitTotal(total int) (partA int, partB int) {
	// TODO: 拆分总数
	return 0, 0
}

func main() {
	fmt.Println("TODO: 完成 05_functions 作业")
}
