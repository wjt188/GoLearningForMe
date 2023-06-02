package main

import (
	"fmt"
	"time"
)

func IntChanRecListener(intChan chan int) {
	intValue := <-intChan
	//fmt.Println(intValue, reflect.TypeOf(intValue))
	fmt.Println("成功接收第第1个value,通道元素个数，当前缓冲区大小：", intValue, len(intChan), cap(intChan))
	intValue = <-intChan
	//fmt.Println(intValue, reflect.TypeOf(intValue))
	fmt.Println("成功接收第第2个value，通道元素个数，当前缓冲区大小：", intValue, len(intChan), cap(intChan))
}

func main() {
	var intChan2 = make(chan int, 2)
	intChan2 <- 100
	fmt.Println(len(intChan2), cap(intChan2))
	intChan2 <- 200
	fmt.Println(len(intChan2), cap(intChan2))
	go IntChanRecListener(intChan2)
	time.Sleep(time.Second * 2)
	close(intChan2)

}
