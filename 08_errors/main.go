package main

import (
	"errors"
	"fmt"
)

// ---- 自定义错误类型 ----

var ErrNotFound = errors.New("记录不存在")

type ValidationError struct {
	Field string
	Msg   string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// ---- 业务函数：返回 error ----

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("除数不能为0")
	}
	return a / b, nil
}

func findUser(id int) (string, error) {
	if id <= 0 {
		return "", ErrNotFound
	}
	return "小明", nil
}

func setAge(age int) error {
	if age < 0 || age > 150 {
		return ValidationError{Field: "age", Msg: "年龄不合法"}
	}
	return nil
}

func main() {
	fmt.Println("--- 基础 error 处理 ---")

	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("divide 失败:", err)
	} else {
		fmt.Println("divide 结果:", result)
	}

	// ---- 包装错误：%w ----
	fmt.Println("\n--- 错误包装 ---")

	err = loadConfig("bad.yaml")
	fmt.Println("原始 err:", err)

	// errors.Is：判断错误链里是否包含某个哨兵错误
	if errors.Is(err, ErrNotFound) {
		fmt.Println("→ 是「记录不存在」类错误")
	}

	// errors.As：把错误链里的自定义类型取出来
	var ve ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("→ 是 ValidationError: field=%s msg=%s\n", ve.Field, ve.Msg)
	}

	// ---- 习惯用法 ----
	fmt.Println("\n--- Go 错误习惯 ---")

	if err := setAge(200); err != nil {
		fmt.Println("setAge:", err)
	}

	name, err := findUser(-1)
	if err != nil {
		fmt.Println("findUser:", err)
	} else {
		fmt.Println("用户:", name)
	}

	fmt.Println("\n记住：")
	fmt.Println("1. 错误也是值，用 if err != nil 处理")
	fmt.Println("2. 向上返回时用 fmt.Errorf(\"...: %w\", err) 包装")
	fmt.Println("3. 判断用 errors.Is / errors.As，不要比字符串")
}

func loadConfig(path string) error {
	// 模拟：底层先发现校验失败
	if err := setAge(200); err != nil {
		return fmt.Errorf("加载配置 %s 失败: %w", path, err)
	}
	return nil
}
