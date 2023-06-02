package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
	"math"
	"os"
)

// 四舍五入取整函数
func Round(v float64) int {
	return int(math.Floor(v + 0.5))
}

// 计算逻辑函数
func GetCalcResult(w int) int {
	//start := time.Now()
	var fre = 0     //运费
	var pre float64 //保险费
	var sum = 18    //总费用
	if w > 100 {
		fmt.Println("超出最大总量")
		os.Exit(1)
	}
	if w <= 1 {
		return sum
	}
	if w > 1 {
		pre = float64(sum) * 0.01
		for i := 0; i < (w - 1); i++ {
			fre += 5
			pre = (pre + float64(fre)) * 0.01
			sum += Round(pre) + fre
			fre = 0
		}
	}
	return sum
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	id     int
	uid    string
	weight float32
}

func main() {
	db, err := sql.Open("sqlite3", "./data.db")
	CheckError(err)
	defer db.Close()
	str := "select id,uid ,weight from test_db where id=?"
	var u User
	var id int
	for {
		fmt.Print("请输入用户id:")
		fmt.Scanf("%d", &id)
		r, _ := db.Query(str, id)

		defer r.Close()
		for r.Next() {
			err = r.Scan(&u.id, &u.uid, &u.weight)
			CheckError(err)

			fmt.Println(" id  uid  weight  快递费用:", u.id, u.uid, u.weight, GetCalcResult(int(u.weight)))

		}

	}

}
