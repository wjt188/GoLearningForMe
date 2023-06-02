package main

import "fmt"

func main() {
	i := 0
	p := &i
	defer fmt.Println("i:", i)
	*p++
	fmt.Println("i:", i)
}
