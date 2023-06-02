package main

import "fmt"

func main() {
	{
		fmt.Println("block ends")
		defer println("defer ends")
	}
	fmt.Println("main ends")
	//for i := 0; i < 5; i++ {
	//	defer fmt.Println(i)
	//}

}
