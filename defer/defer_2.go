package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("defer_3")
	}()
	defer func() {
		fmt.Println("defer_2")
	}()
	defer func() {
		fmt.Println("defer_1")
	}()
}
