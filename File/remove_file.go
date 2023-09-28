package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("./File/test.txt")
	if err != nil {
		fmt.Println("文件删除失败")
		return
	}
}
