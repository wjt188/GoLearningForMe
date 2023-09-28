package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	var locker sync.Mutex
	fmt.Println("Locking (G0)")
	locker.Lock()
	fmt.Println("locked (G0)")

	for i := 1; i < 4; i++ {
		go func(i int) {
			fmt.Printf("Locking (G%d)\n", i)
			locker.Lock()
			fmt.Printf("locked (G%d)\n", i)
			time.Sleep(time.Second * 2)
			locker.Unlock()
			fmt.Printf("Unlocked (G%d)\n", i)
			wg.Done()
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println("ready unlock (G0)")
	locker.Unlock()
	fmt.Println("Unlocked (G0)")
	wg.Wait()
}
