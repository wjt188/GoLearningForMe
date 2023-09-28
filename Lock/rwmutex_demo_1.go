package main

import (
	"fmt"
	"sync"
	"time"
)

var countTest int
var locker sync.RWMutex

func main() {
	for i := 1; i <= 3; i++ {
		go Write(i)
	}
	for i := 1; i <= 3; i++ {
		go Read(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("countTest的值为：", countTest)
}
func Read(i int) {
	fmt.Println("读操作：", i)
	locker.RLock()
	fmt.Println(i, "读countTest的值为：", countTest)
	time.Sleep(1 * time.Second)
	locker.RUnlock()
}
func Write(i int) {
	fmt.Println("写操作", i)
	locker.Lock()
	countTest++
	fmt.Println(i, "写countTest的值为：", countTest)
	time.Sleep(1 * time.Second)
	locker.Unlock()

}
