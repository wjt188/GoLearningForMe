package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("./File/testDir", os.ModePerm)
	if err != nil {
		fmt.Println("目录创建失败")
		return
	}

}
