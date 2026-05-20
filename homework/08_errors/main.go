package main

import (
	"errors"
	"fmt"
)

/*
第 08 次作业：错误处理（进阶版）

完成要求（必做）：
1) 定义哨兵错误：ErrInvalidConfig。
2) 写函数 loadConfig(path string) error：
   - path 为空时返回 ErrInvalidConfig
3) 写函数 boot(path string) error：
   - 调用 loadConfig，并用 fmt.Errorf("boot failed: %w", err) 包装
4) 在 main 中用 errors.Is 判断是否是 ErrInvalidConfig。

加分挑战（选做）：
- 自定义结构化错误类型并用 errors.As 解析字段。
*/

var ErrInvalidConfig = errors.New("invalid config")

func loadConfig(path string) error {
	// TODO: 校验参数并返回错误
	return nil
}

func boot(path string) error {
	// TODO: 包装底层错误并返回
	return nil
}

func main() {
	fmt.Println("TODO: 完成 08_errors 作业")
}

