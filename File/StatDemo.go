package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./File/testDir/test.txt")
	if err != nil {
		if os.IsExist(err) {
			fmt.Println("文件存在")
		} else {
			fmt.Println("文件不存在")
		}
	} else {
		fmt.Println("文件存在")
		fmt.Println("文件大小：", fileInfo.Size())
		fmt.Println("文件权限：", fileInfo.Mode())
		fmt.Println("文件修改日期：", fileInfo.ModTime())
		fmt.Println("文件名：", fileInfo.Name())

	}

}
