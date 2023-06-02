package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{2, 4, 5, 1, 0, -2, 3, 9}
	fmt.Println("原始数组：", a)
	sort.Ints(a)
	fmt.Println("排序后的数组：", a)
}
