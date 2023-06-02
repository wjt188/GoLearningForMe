package main

import (
	"fmt"
	"unsafe"
)

type V struct {
	i int32
	j int64
}

func (this V) PutI() {
	fmt.Printf("i=%d\n", this.i)
}
func (this V) PutJ() {
	fmt.Printf("j=%d\n", this.j)
}
func main() {
	var v = new(V)
	var i = (*int32)(unsafe.Pointer(v))
	*i = int32(98)
	var j = (*int64)(unsafe.Pointer(uintptr(unsafe.Pointer(v)) + uintptr(unsafe.Sizeof(int32(0)))))
	*j = int64(763)
	v.PutI()
	v.PutJ()
}
