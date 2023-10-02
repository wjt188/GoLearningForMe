package main

import (
	"fmt"
	"runtime"
)

func main() {
	var a, b int
	a = 10
	b = 20
	var ptra *int
	var ptrb = &b

	ptra = &a
	fmt.Printf("a 变量的地址是: %x\n", &a)
	fmt.Printf("b 变量的地址是: %x\n", &b)

	// 指针变量的存储地址
	fmt.Printf("ptra 变量的存储地址: %x\n", ptra)
	fmt.Printf("ptrb 变量的存储地址: %x\n", ptrb)

	// 使用指针访问值
	fmt.Printf("*ptra 变量的值: %d\n", *ptra)
	fmt.Printf("*ptrb 变量的值: %d\n", *ptrb)
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}
