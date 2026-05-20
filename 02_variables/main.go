package main

import "fmt"

func main() {
	// 方式1：完整声明 var 变量名 类型 = 值
	var name string = "小明"

	// 方式2：类型推导，省略类型
	var age = 25

	// 方式3：短声明（最常用），只能在函数内用
	height := 175.5

	// 方式4：先声明再赋值
	var city string
	city = "北京"

	// 多变量声明
	x, y := 10, 20

	fmt.Println(name, age, height, city, x, y)

	// Go 的基本类型
	var (
		i   int     = -42
		u   uint    = 42
		f   float64 = 3.14
		b   bool    = true
		s   string  = "你好"
		r   rune    = '中' // rune = Unicode 码点（字符）
	)
	fmt.Println(i, u, f, b, s, r)
}
