// package main

// import (
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{ /* SkipDefaultTransaction: true */ }) // 全局禁用事务
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = d
// }

// func test1() {
// 	// session级别禁用事务
// 	db.Session(&gorm.Session{SkipDefaultTransaction: true})
// }

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      int
// 	Birthday time.Time
// 	Active   bool
// }

// // 测试事务操作 没有使用事务控制
// func test2() {
// 	user := User{
// 		Name: "tom",
// 	}
// 	db.Create(&user)
// 	db.Create(nil)
// }

// // 测试事务操作 使用事务控制
// func test3() {
// 	user := User{
// 		Name: "tom",
// 	}
// 	//手动控制
// 	tx := db.Begin()
// 	tx.Create(&user)
// 	err := tx.Create(nil).Error
// 	if err != nil {
// 		tx.Rollback()
// 	} else {
// 		tx.Commit()
// 	}
// }

// func test4() {
// 	user := User{
// 		Name: "tom",
// 	}
// 	db.Transaction(func(tx *gorm.DB) error {
// 		if err := tx.Create(&user).Error; err != nil {
// 			return err
// 		}
// 		if err := tx.Create(nil).Error; err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// func main() {

// }
