package main

import (
	"database/sql"
	"fmt"
	snowflak "github.com/bwmarrin/snowflake"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"math/rand"
	"os"
	"time"
)

func GenerateWeightFunc() float32 {
	value := rand.Float32() * 100
	return value
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic("初始化数据库失败")
	}
	defer db.Close()
	stmt, err := db.Prepare("INSERT INTO test_db(id,uid,weight,created_at) values (?,?,?,?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	node, err := snowflak.NewNode(1)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	start := time.Now()

	for i := 1; i <= 1000; i++ {
		for j := 1; j <= 100; j++ {
			weight := GenerateWeightFunc()
			uid := node.Generate()
			_, err := stmt.Exec(i, uid, weight, start)
			if err != nil {
				panic(err)
				fmt.Println("插入过程中产生错误")
			}
		}
	}

}
