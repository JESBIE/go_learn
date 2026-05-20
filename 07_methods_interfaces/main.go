package main

import (
	"fmt"
	"math"
)

// ========== 方法：给类型加「行为」 ==========

type Student struct {
	Name  string
	Score int
}

// 值接收者：不修改原 struct（改的是副本）
func (s Student) PassOrFail() string {
	if s.Score >= 60 {
		return "及格"
	}
	return "不及格"
}

// 指针接收者：可以修改原 struct
func (s *Student) AddBonus(bonus int) {
	s.Score += bonus
}

// ========== 接口：只规定「能做什么」 ==========

// 任何实现了 Area() 和 Name() 的类型，都自动满足 Shape 接口（无需 implements）
type Shape interface {
	Area() float64
	Name() string
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Name() string {
	return "圆形"
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Name() string {
	return "矩形"
}

// 接受接口类型 → 多种 struct 都能传进来（多态）
func printShape(s Shape) {
	fmt.Printf("%s 面积=%.2f\n", s.Name(), s.Area())
}

// ========== 类型断言：从 interface{} 里取出具体类型 ==========

func describe(v any) {
	fmt.Print("值: ", v)
	if s, ok := v.(string); ok {
		fmt.Printf(" → 是 string，长度 %d\n", len(s))
	} else {
		fmt.Println()
	}
}

func main() {
	fmt.Println("--- 方法 ---")

	stu := Student{Name: "小明", Score: 55}
	fmt.Println(stu.PassOrFail()) // 值接收者，stu 和 &stu 都能调

	stu.AddBonus(10) // 指针接收者，会改原数据
	fmt.Println("加分后:", stu.Score, stu.PassOrFail())

	// 值接收者 vs 指针接收者 怎么选？
	fmt.Println("\n--- 接收者选择 ---")
	fmt.Println("要改字段 → 用指针接收者 (*T)")
	fmt.Println("struct 很大 → 用指针接收者（少拷贝）")
	fmt.Println("小 struct、只读 → 值接收者 (T) 即可")
	fmt.Println("拿不准 → 同一类型的方法统一用指针接收者")

	// ========== 接口 ==========
	fmt.Println("\n--- 接口 ---")

	shapes := []Shape{
		Circle{Radius: 3},
		Rectangle{Width: 4, Height: 5},
	}

	for _, s := range shapes {
		printShape(s) // Circle、Rectangle 都满足 Shape
	}

	// 接口变量的零值是 nil
	var s Shape
	fmt.Println("nil 接口 == nil?", s == nil)

	// ========== any（空接口）==========
	fmt.Println("\n--- any / interface{} ---")

	describe("Go")
	describe(42)
	describe(Circle{Radius: 1})

	// 标准库 fmt.Println 的参数就是 ...any
	fmt.Println("混合类型:", 1, "hello", true)
}
