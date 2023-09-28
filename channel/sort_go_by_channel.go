package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan struct{})
	c2 := make(chan struct{})
	c3 := make(chan struct{})
	go func() {
		fmt.Println("goroutine 1")
		close(c1)
	}()
	go func() {
		<-c1
		fmt.Println("goroutine 2")
		close(c2)
	}()
	go func() {
		<-c2
		fmt.Println("goroutine 3")
		close(c3)
	}()
	time.Sleep(time.Second)
}
