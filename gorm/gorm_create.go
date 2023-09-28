package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// 创建全局结构体
type Student struct {
	Id   int    `json:"id"` //默认主键
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var GlobleDB *gorm.DB

func main() {
	//链接数据库
	conn, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8")
	if err != nil {
		fmt.Println("gorm Open err ", err)
		return
	}
	defer conn.Close()
	GlobleDB = conn
	////defer conn.Close()
	//conn.DB().SetMaxIdleConns(10)
	//conn.DB().SetMaxOpenConns(10)
	////借助gorm，创建数据库表
	//fmt.Println(conn.AutoMigrate(new(Student)).Error)
	//Insert()
	//Serch()
	Update()
}

//func Insert() {
//	var stu Student
//	var age int
//	age = 10
//	for i := 1; i < 10; i++ {
//		stu.Name = fmt.Sprintf("kol-%d", i)
//		stu.Age = age + i
//		GlobleDB.Create(&stu)
//	}
//
//}

//	func Serch() {
//		var stu []Student
//		GlobleDB.Find(&stu)
//		fmt.Println(stu)
//	}
func Update() {
	GlobleDB.Model(new(Student)).Where("name = ?", "kol").Update("name", "curry")
}
