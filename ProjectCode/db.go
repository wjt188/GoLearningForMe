package main

import (
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//db, err := sql.Open("sqlite3", "test.db")
	//if err != nil {
	//	panic("连接失败")
	//}
	//defer db.Close()
	//node, err := snowflak.NewNode(1)
	//if err != nil {
	//	println(err.Error())
	//	os.Exit(1)
	//}
	//
	//for i := 0; i < 20; i++ {
	//	id := node.Generate()
	//
	//	fmt.Printf("int64 ID: %d\n", id)
	//	fmt.Printf("string ID: %s\n", id)
	//	fmt.Printf("base2 ID: %s\n", id.Base2())
	//	fmt.Printf("base64 ID: %s\n", id.Base64())
	//	fmt.Printf("ID time: %d\n", id.Time())
	//	fmt.Printf("ID node: %d\n", id.Node())
	//	fmt.Printf("ID step: %d\n", id.Step())
	//	fmt.Println("--------------------------------")
	//}
	var username string

	for {
		fmt.Scanf("%s", &username)
		fmt.Println(username)
	}

}
