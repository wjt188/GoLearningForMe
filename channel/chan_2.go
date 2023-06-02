package main

import "fmt"

//判断channel是否关闭

func main() {
	initchan := make(chan int, 10)
	initchan <- 1
	initchan <- 2
	initchan <- 3
	close(initchan)
	for {
		initvalue, isOpen := <-initchan
		if !isOpen {
			fmt.Println("channel 已经关闭")
			break
		}
		fmt.Println(initvalue)
	}

}
