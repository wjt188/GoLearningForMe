package main

import (
	"fmt"
	"math/rand"
)

func GenerateWeight() float32 {
	value := rand.Float32() * 100
	return value
}

func main() {
	value := GenerateWeight()
	fmt.Println(value)
}
