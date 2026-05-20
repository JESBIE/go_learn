package main

import "fmt"

func main() {
	// ---- Slice（切片）：动态数组 ----
	fmt.Println("--- Slice ---")

	// 字面量创建
	fruits := []string{"苹果", "香蕉", "橘子"}
	fmt.Println("fruits:", fruits)
	for pos, fruit := range fruits {
		fmt.Println(pos, fruit)
	}

	// append 追加
	fruits = append(fruits, "西瓜")
	fruits = append(fruits, "草莓")
	fmt.Println("append后:", fruits, "长度:", len(fruits))

	// make 创建：make([]T, length, capacity)
	nums := make([]int, 3, 5) // 长度3，容量5
	fmt.Println("nums:", nums, "len:", len(nums), "cap:", cap(nums))

	// 切片操作 [start:end]（左闭右开）
	fmt.Println("前两个:", fruits[0:2])

	// ---- Map：字典 ----
	fmt.Println("\n--- Map ---")

	// 字面量
	scores := map[string]int{
		"小明": 90,
		"小红": 85,
	}
	fmt.Println("scores:", scores)

	// 插入
	scores["小刚"] = 72

	// 取值 + 检测 key 是否存在（comma ok 惯用法）
	score, exists := scores["小刚"]
	fmt.Println("小刚:", score, "存在?", exists)

	score, exists = scores["不存在"]
	fmt.Println("不存在:", score, "存在?", exists) // 默认零值 0

	// ---- Struct：结构体 ----
	fmt.Println("\n--- Struct ---")

	type Person struct {
		Name string
		Age  int
	}

	// 创建方式
	p1 := Person{Name: "小明", Age: 20}
	p2 := Person{"小红", 18} // 按字段顺序（不推荐，可读性差）
	var p3 Person          // 零值
	p3.Name = "小刚"
	p3.Age = 22

	fmt.Println(p1, p2, p3)

	numbers := []int{1, 2, 3, 4, 5}
	for idx, number := range numbers {
		fmt.Printf("第%d个：%d\n", idx, number)
	}
}
