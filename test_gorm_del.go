// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      int
// 	Birthday time.Time
// 	Active   bool
// }

// func delete1() {
// 	var user User
// 	db.First(&user)
// 	db.Delete(&user)
// }

// func delete2() {
// 	var user User
// 	db.First(&user)
// 	db.Where("name=?", "jiji").Delete(&user)
// 	// DELETE from users where id=1 AND name="jiji"
// }

// func delete3() {
// 	db.Delete(&User{}, 10)
// 	// DELETE FROM users WHERE id=10
// 	db.Delete(&User{}, []int{1, 2, 3})
// 	// DELETE FROM users WHERE id IN (1,2,3)
// }

// // delete hook
// func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
// 	fmt.Println("before delete")
// 	return
// }

// func delete4() {
// 	db.Where("name LIKE ?", "%ji%").Delete(User{})
// 	// DELETE from users where name LIKE "%ji%"
// 	db.Delete(User{}, "name LIKE ?", "%ji%")
// }
// func main() {
// 	delete1()
// }

// var db *gorm.DB

// func init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Panicf("database connect failed: %v\n", err)
// 	}
// 	db = d
// }
