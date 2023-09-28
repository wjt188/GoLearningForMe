package main

import (
	"fmt"
	"sync"
)

func main() {
	var m sync.Map
	m.Store("name", "curry")
	m.Store("address", "gold state")
	m.Store("id", 101678)
	v, ok := m.LoadOrStore("name1", "nick")
	fmt.Println(v, ok)
	v, ok = m.LoadOrStore("name", "kd")
	fmt.Println(v, ok)
	f := func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	}
	m.Range(f)
	m.Delete("name1")
	fmt.Println(m.Load("name1"))
}
