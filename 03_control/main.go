package main

import "fmt"

func main() {
	// Go 只有 for 一种循环
	fmt.Println("--- for 循环 ---")

	// 经典 C 风格
	for i := 1; i <= 3; i++ {
		fmt.Println("i:", i)
	}

	// 类似 while
	count := 0
	for count < 3 {
		fmt.Println("count:", count)
		count++
	}

	// range 遍历 slice/array/map/string
	names := []string{"小明", "小红", "小刚"}
	for idx, name := range names {
		fmt.Printf("第%d个：%s\n", idx, name)
	}

	// --- if ---
	fmt.Println("\n--- if ---")

	score := 85
	// if 可以带短声明（只在 if 块内有效）
	if grade := "B"; score >= 90 {
		grade = "A"
		fmt.Println("优秀", grade)
	} else if score >= 60 {
		fmt.Println("及格", grade)
	} else {
		fmt.Println("不及格", grade)
	}

	// --- switch ---
	fmt.Println("\n--- switch ---")

	season := "夏"
	switch season {
	case "春":
		fmt.Println("春天")
	case "夏":
		fmt.Println("夏天")
		// Go 的 switch 不用 break，自动终止！
	case "秋", "冬": // 可以多值
		fmt.Println("秋天或冬天")
	default:
		fmt.Println("未知季节")
	}

	// 无表达式 switch = 替代 if-else 链
	num := 75
	switch { // 不跟表达式
	case num >= 90:
		fmt.Println("A")
	case num >= 60:
		fmt.Println("B")
	default:
		fmt.Println("C")
	}

	
}
