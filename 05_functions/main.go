package main

import (
	"errors"
	"fmt"
)

// 基础函数：参数在前，返回类型在后
func add(a int, b int) int {
	return a + b
}

// 同类型参数可缩写
func multiply(a, b int) int {
	return a * b
}

// 多返回值（Go 的招牌特性）
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

// 命名返回值 + 裸 return（简短函数才用）
func split(sum int) (a, b int) {
	a = sum / 2
	b = sum - a
	return // 裸 return，返回 a 和 b
}

// defer：函数结束时执行（类似 finally）
func readFile(path string) {
	fmt.Println("打开文件:", path)
	defer fmt.Println("关闭文件:", path) // 函数返回前一定执行
	fmt.Println("读取数据...")
}

// 可变参数
func join(sep string, words ...string) string {
	result := ""
	for i, w := range words {
		if i > 0 {
			result += sep
		}
		result += w
	}
	return result
}

func main() {
	fmt.Println("add:", add(3, 5))
	fmt.Println("multiply:", multiply(4, 7))

	// 接收多返回值
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Println("错误:", err)
	} else {
		fmt.Println("divide:", quotient)
	}

	// 不想用的返回值用 _ 丢弃
	q, _ := divide(9, 3)
	fmt.Println("divide(丢弃err):", q)

	a, b := split(10)
	fmt.Printf("split(10): a=%d, b=%d\n", a, b)

	readFile("hello.txt")
	fmt.Println("join:", join(", ", "Go", "Python", "Rust"))
}
