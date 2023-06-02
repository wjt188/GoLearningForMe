package main

import (
	"fmt"
)

func Send(chan1 chan int) {
	for i := 0; i < 10; i++ {
		chan1 <- i
		//time.Sleep(1 * time.Second)
	}
}
func recvFunc(chan1 chan int) {
	for {
		select {
		case intvalue1 := <-chan1:
			fmt.Println("", intvalue1)
		}
	}
}

func main() {
	ch := make(chan int)
	signal := make(chan int)
	go func() {
		for item := range ch {
			fmt.Println(item)
		}
		signal <- 2
	}()
	for i := 0; i <= 10; i++ {
		ch <- i

	}
	close(ch)
	<-signal
	//go Send(ch)
	//go recvFunc(ch)
	//time.Sleep(5 * time.Second)

}
