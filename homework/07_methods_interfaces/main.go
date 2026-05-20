package main

import (
	"fmt"
	"math"
)

/*
第 07 次作业：方法与接口（进阶版）

完成要求（必做）：
1) 定义接口 Shape：
   - Area() float64
   - Name() string
2) 实现 Rectangle 和 Circle 两个类型，满足 Shape 接口。
3) 写函数 printShapeInfo(s Shape)，打印图形名与面积（保留两位小数）。
4) 使用 []Shape 存多个图形并遍历输出。

加分挑战（选做）：
- 为 Rectangle 增加 Scale(f float64) 方法（指针接收者）。
*/

type Shape interface {
	Area() float64
	Name() string
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 { return 0 } // TODO
func (r Rectangle) Name() string  { return "TODO" }

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { return 0 } // TODO
func (c Circle) Name() string  { return "TODO" }

func printShapeInfo(s Shape) {
	fmt.Printf("%s area=%.2f\n", s.Name(), s.Area())
}

func main() {
	_ = math.Pi
	fmt.Println("TODO: 完成 07_methods_interfaces 作业")
}

