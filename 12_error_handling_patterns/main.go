package main

import (
	"errors"
	"fmt"
	"os"
)

func loadConfig(path string) error {
	_, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("读取配置失败: %w", err)
	}
	return nil
}
func main() {
	err := loadConfig("config.yaml")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("配置文件不存在", err)

		}
		if errors.Is(err, os.ErrPermission) {
			fmt.Println("没有权限读取配置文件", err)

		}
		fmt.Println("读取配置文件失败", err)

	}

	var pathErr *os.PathError
	if errors.As(err, &pathErr) {
		fmt.Println("底层路径错误:", "op=", pathErr.Op, "path=", pathErr.Path)
	}
	fmt.Println("Config loaded successfully")

}
