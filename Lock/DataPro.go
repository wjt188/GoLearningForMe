package main

import (
	"fmt"
	"sync"
	"time"
)

var count int
var lock sync.Mutex

func main() {
	go CountPlus(10000)
	go CountPlus(10000)
	time.Sleep(time.Second)
	fmt.Println(count)
}
func CountPlus(times int) {
	for i := 0; i < times; i++ {
		lock.Lock()
		count++
		lock.Unlock()
	}
}
