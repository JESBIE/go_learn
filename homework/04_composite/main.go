package main

import "fmt"

/*
第 04 次作业：复合类型（进阶版）

完成要求（必做）：
1) 定义结构体 Student：
   - Name string
   - Scores []int
2) 写函数 avg(scores []int) float64，计算平均分（空切片返回 0）。
3) 使用 map[string]Student 存 3 个学生，输出每位同学平均分。
4) 找出平均分最高的学生并打印。

加分挑战（选做）：
- 把 map 遍历结果整理成切片并按平均分降序输出。
*/

type Student struct {
	Name   string
	Scores []int
}

func avg(scores []int) float64 {

	if len(scores) == 0 {
		return 0
	}
	// TODO: 计算平均值，注意空切片
	result := 0.0

	for _, score := range scores {
		result += float64(score)
	}

	return result / float64(len(scores))

}

func main() {
	fmt.Println("TODO: 完成 04_composite 作业")

	students := map[string]Student{
		"小明": Student{Name: "小明", Scores: []int{90, 80, 70}},
		"小红": Student{Name: "小红", Scores: []int{85, 88, 90}},
		"小刚": Student{Name: "小刚", Scores: []int{70, 75, 80}},
	}

	for name, student := range students {
		fmt.Println(name, student.Name)
		fmt.Println(name, student.Scores)
		fmt.Println(name, avg(student.Scores))
	}

}
