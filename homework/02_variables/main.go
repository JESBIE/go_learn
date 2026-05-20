package main

import "fmt"

/*
第 02 次作业：变量、类型与转换（进阶版）

完成要求（必做）：
1) 同时使用 var 和 := 声明变量，包含：
   - string、int、float64、bool
2) 编写函数 classifyAge(age int) string：
   - <18: 未成年
   - 18~59: 成年
   - >=60: 长者
3) 演示一次类型转换：int -> float64，并保留两位小数输出。

加分挑战（选做）：
- 用 const 定义税率，写一个函数计算含税价格。
*/

func classifyAge(age int) string {
	// TODO: 按规则返回年龄分类
	result := ""
	if age < 18 {
		result = "未成年"
	} else if age >= 18 && age <= 59 {
		result = "成年"
	} else {
		result = "长者"
	}
	return result
}

func main() {
	fmt.Println("TODO: 完成 02_variables 作业")
	fmt.Println(classifyAge(10))
	fmt.Println(classifyAge(20))
	fmt.Println(classifyAge(60))

	age := 20
	fmt.Println(float64(age) / 3.0)
}
