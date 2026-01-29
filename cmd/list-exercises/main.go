// list-exercises 列出所有可运行的练习（包含 main.go 的目录）。
// 用法：go run ./cmd/list-exercises
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	root := "."
	if len(os.Args) > 1 {
		root = os.Args[1]
	}
	fmt.Println("可运行的练习（go run ./路径）：")
	fmt.Println()
	var count int
	_ = filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil
		}
		if d.IsDir() {
			mainPath := filepath.Join(path, "main.go")
			if info, err := os.Stat(mainPath); err == nil && !info.IsDir() {
				rel, _ := filepath.Rel(root, path)
				// 跳过 cmd 下的工具（如本命令）
				if rel == "cmd" || strings.HasPrefix(rel, "cmd"+string(filepath.Separator)) {
					return nil
				}
				fmt.Printf("  go run ./%s\n", path)
				count++
			}
		}
		return nil
	})
	fmt.Println()
	fmt.Printf("共 %d 个练习\n", count)
}
