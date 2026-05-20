package main

import "fmt"

// ---- 值传递：函数里改的是副本 ----
func addOneValue(n int) {
	n++
}

// ---- 指针传递：函数里改的是原变量 ----
func addOnePtr(n *int) {
	*n++ // * 解引用：拿到指针指向的值
}

// ---- 用指针改 struct ----
type Student struct {
	Name  string
	Score int
}

func bumpScore(s Student) {
	s.Score += 10 // 改的是副本，main 里不变
}

func bumpScorePtr(s *Student) {
	s.Score += 10 // 等价于 (*s).Score，Go 自动帮你解引用
}

func main() {
	fmt.Println("--- 指针基础 ---")

	x := 42
	p := &x // & 取地址：p 的类型是 *int

	fmt.Println("x 的值:", x)
	fmt.Println("p 存的地址:", p)
	fmt.Println("*p 解引用:", *p) // * 读指针指向的值

	*p = 100 // 通过指针改 x
	fmt.Println("改 *p 后 x:", x)

	// ---- nil 指针 ----
	fmt.Println("\n--- nil 指针 ---")

	var p2 *int // 零值是 nil（没有指向任何变量）
	fmt.Println("p2 == nil?", p2 == nil)

	// *p2 = 1  // 取消注释会 panic：不能解引用 nil
	p2 = &x
	fmt.Println("*p2:", *p2)

	// ---- 值传递 vs 指针传递 ----
	fmt.Println("\n--- 值传递 vs 指针传递 ---")

	n := 10
	addOneValue(n)
	fmt.Println("值传递后 n:", n) // 还是 10

	addOnePtr(&n)
	fmt.Println("指针传递后 n:", n) // 11

	// ---- struct 与指针 ----
	fmt.Println("\n--- struct 指针 ---")

	s1 := Student{Name: "小明", Score: 80}
	bumpScore(s1)
	fmt.Println("值传递:", s1.Score) // 80

	bumpScorePtr(&s1)
	fmt.Println("指针传递:", s1.Score) // 90

	// 两种创建指针的写法
	s2 := &Student{Name: "小红", Score: 85} // 取 struct 字面量的地址
	s3 := new(Student)                      // 分配零值 Student，返回 *Student
	s3.Name = "小刚"
	s3.Score = 72
	fmt.Println("s2:", s2)
	fmt.Println("s3:", s3)

	// ---- 什么时候用指针？ ----
	fmt.Println("\n--- 何时用指针 ---")
	fmt.Println("1. 需要在函数里修改调用方的变量")
	fmt.Println("2. struct 较大，避免拷贝（性能）")
	fmt.Println("3. 需要表示「可能没有值」→ 用 nil 指针（进阶）")
	fmt.Println("注意：Go 没有指针运算（不能像 C 那样 p++）")
}
