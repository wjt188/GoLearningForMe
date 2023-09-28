package main

//select 的用法：select结构会同时监听一个或多个通道，简化处理多个通道的流程。
import (
	"fmt"
	"time"
)

func sendFunc1(chan1 chan int) {
	for i := 0; i < 5; i++ {
		chan1 <- i
		time.Sleep(1 * time.Second)
	}
}
func sendFunc2(chan2 chan int) {
	for i := 10; i >= 5; i-- {
		chan2 <- i
		time.Sleep(1 * time.Second)
	}
}
func recvFunc(chan1 chan int, chan2 chan int) {
	for {
		select {
		case intvalue1 := <-chan1:
			fmt.Println("接收到chan1通道的值：", intvalue1)
		case intvalue2 := <-chan2:
			fmt.Println("接收到chan2通道的值：", intvalue2)

		}
	}
}

func main() {
	chan1 := make(chan int, 5)
	chan2 := make(chan int, 5)
	go recvFunc(chan1, chan2)
	go sendFunc1(chan1)
	go sendFunc2(chan2)
	time.Sleep(time.Second * 5)
	fmt.Println("main程序结束")
}
