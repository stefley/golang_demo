// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// var db *sql.DB

// func initDB() (err error) {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
// 	db, err = sql.Open("mysql", dsn)
// 	if err != nil {
// 		return err
// 	}
// 	err = db.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }

// // 插入数据
// func insertData() {
// 	fmt.Println("开始插入数据")
// 	sqlStr := "insert into user_tbl(username,password) values(?,?)"
// 	ret, err := db.Exec(sqlStr, "张三", "zs123")
// 	if err != nil {
// 		fmt.Printf("insert failed,err: %v\n", err)
// 		return
// 	}
// 	theID, err := ret.LastInsertId() //新插入的数据id
// 	if err != nil {
// 		fmt.Printf("get lastInsertId failed,err: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("insert success,theID: %v\n", theID)
// }

// type User struct {
// 	id       int
// 	username string
// 	password string
// }

// // 查询一条
// func queryOneRow() {
// 	s := "select * from user_tbl where id=?"
// 	var u User
// 	err := db.QueryRow(s, 4).Scan(&u.id, &u.username, &u.password)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("u: %v\n", u)
// 	}
// }

// // 查询多行
// func queryManyRow() {
// 	s := "select * from user_tbl"
// 	r, err := db.Query(s)
// 	var u User
// 	defer r.Close()
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		for r.Next() {
// 			r.Scan(&u.id, &u.username, &u.password)
// 			fmt.Printf("u: %v\n", u)
// 		}
// 	}
// }

// // 更新
// func update() {
// 	sqlStr := "update user_tbl set username=?, password=? where id=?"
// 	r, err := db.Exec(sqlStr, "bigtom", "tto", 4)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		i, _ := r.RowsAffected()
// 		fmt.Printf("i: %v\n", i)
// 	}
// }

// // 删除
// func delete() {
// 	s := "delete from user_tbl where id=?"
// 	r, err := db.Exec(s, 4)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		i, _ := r.RowsAffected()
// 		fmt.Printf("i: %v\n", i)
// 	}
// }
// func main() {
// 	err := initDB()
// 	if err != nil {
// 		fmt.Printf("初始化失败！, err: %v \n", err)
// 	} else {
// 		fmt.Println("初始化成功")
// 	}
// 	insertData()
// 	queryOneRow()
// 	queryManyRow()
// 	update()
// 	delete()
// }
