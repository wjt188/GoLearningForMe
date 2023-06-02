package main

import (
	"fmt"
	"reflect"
	"time"
)

func IntChanRecvListener(intChan chan int) {
	intValue := <-intChan
	fmt.Println(intValue, reflect.TypeOf(intValue))
}

func main() {
	var intChan = make(chan int)
	fmt.Println(intChan, reflect.TypeOf(intChan))
	go IntChanRecvListener(intChan)
	intChan <- 100
	close(intChan)
	time.Sleep(time.Second)
}
