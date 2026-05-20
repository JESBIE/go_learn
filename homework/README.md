# Go 复习作业清单

使用方法：
- 每完成一题，把 `[ ]` 改成 `[x]`
- 每章尽量独立写在对应章节目录中
- 提交前至少运行一次 `go run .`

---

## 01_hello

- [ ] 打印三行内容：你的名字、学习目标、今天日期
- [ ] 使用 `fmt.Printf` 打印一句自我介绍
- [ ] 用 1-2 句话写出 `go run .` 和 `go build` 的区别

## 02_variables

- [ ] 定义 `name`、`age`、`height`、`isStudent` 四个变量并输出
- [ ] 同一组变量分别用 `var` 和 `:=` 写一遍
- [ ] 写一个判断年龄是否成年的函数

## 03_control

- [ ] 用 `if` 判断成绩是否及格
- [ ] 用 `for` 打印 1 到 10
- [ ] 用 `switch` 输入 1~7 输出星期几

## 04_composite

- [ ] 创建一个 `[]int` 切片，计算总和和平均值
- [ ] 创建一个 `map[string]int`，保存 3 个同学成绩并遍历输出
- [ ] 定义 `Book{Title, Price}` 并创建两个实例输出

## 05_functions

- [ ] 编写 `add(a, b)` 和 `sub(a, b)` 函数
- [ ] 编写 `divide(a, b) (int, error)`，处理除数为 0
- [ ] 编写可变参数函数，计算任意个整数总和

## 06_pointers

- [ ] 编写 `increase(n *int)`，让传入值加 1
- [ ] 对比传值调用和传指针调用的输出差异
- [ ] 用一句话写出“何时用指针参数”

## 07_methods_interfaces

- [ ] 定义 `Rectangle` 并实现方法 `Area()`
- [ ] 定义接口 `Shape`，再实现 `Circle`
- [ ] 编写 `printArea(s Shape)`，分别传入 `Rectangle` 和 `Circle`

## 08_errors

- [ ] 写一个会返回错误的函数（如参数校验失败）
- [ ] 用 `fmt.Errorf("...: %w", err)` 包装错误
- [ ] 用 `errors.Is` 或 `errors.As` 做错误判断

## 09_concurrency

- [ ] 启动 10 个 goroutine 打印编号，并用 `WaitGroup` 等待结束
- [ ] 共享计数器先写不加锁版本，再改为 `Mutex` 版本
- [ ] 运行一次 `go run -race .`，记录你的观察

---

## 自评（每章完成后打分）

- [ ] 01_hello（0-5）：__
- [ ] 02_variables（0-5）：__
- [ ] 03_control（0-5）：__
- [ ] 04_composite（0-5）：__
- [ ] 05_functions（0-5）：__
- [ ] 06_pointers（0-5）：__
- [ ] 07_methods_interfaces（0-5）：__
- [ ] 08_errors（0-5）：__
- [ ] 09_concurrency（0-5）：__

总分：__/45
