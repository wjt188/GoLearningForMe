package main

import (
	"fmt"
	snowflak "github.com/bwmarrin/snowflake"
	"os"
	"reflect"
)

func main() {
	n, err := snowflak.NewNode(2)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	slice := make([]int, 0)
	for i := 0; i < 3; i++ {
		id := n.Generate()
		slice = append(slice, int(id))
		fmt.Println("id:", id)
		fmt.Println(reflect.TypeOf(id))
		fmt.Println("Node:", id.Node())
		fmt.Println("Step:", id.Step())
		fmt.Println("Time:", id.Time())
	}
	fmt.Println(slice)

}
