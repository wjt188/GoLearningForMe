package main

import (
	"fmt"
	"os"
)

// 通过Create方法来创建文件
func main() {
	file, err := os.Create("./File/test.txt")
	if err != nil {
		fmt.Println("文件创建失败")
		return
	}
	file.Close()
}
