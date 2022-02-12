// package main

// import (
// 	"fmt"
// 	"log"
// 	"time"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var db *gorm.DB

// func init() {
// 	dsn := "root:123456@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
// 	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	db = d
// }

// type User struct {
// 	gorm.Model
// 	Name     string
// 	Age      uint
// 	Birthday time.Time
// }

// func createTable() {
// 	db.AutoMigrate(&User{})
// }

// func inset1() {
// 	user := User{
// 		Name:     "mahuhu",
// 		Age:      25,
// 		Birthday: time.Now(),
// 	}
// 	db.Create(&user)
// }
// func inset2() {
// 	user := User{
// 		Name:     "chenxiaojun",
// 		Age:      28,
// 		Birthday: time.Now(),
// 	}
// 	db.Select("Name", "Age", "CreateAt").Create(&user)
// 	fmt.Printf("user: %v\n", user)
// }

// func inset3() {
// 	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
// 	db.Create(&users)

// 	for _, user := range users {
// 		fmt.Printf("user.ID: %v\n", user.ID)
// 	}
// }

// // 创建钩子
// func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
// 	fmt.Println("before create")
// 	// u.UUID = uuid.New()

// 	// if u.Role == "admin" {
// 	// 	return errors.New("invalid role")
// 	// }
// 	return
// }

// // 根据map类型创建
// func insert4() {
// 	// db.Model(&User{}).Create(map[string]interface{}{
// 	// 	"Name": "fg", "Age": 20,
// 	// })
// 	db.Model(&User{}).Create([]map[string]interface{}{
// 		{"Name": "liujie", "Age": 20},
// 		{"Name": "fanggang", "Age": 28, "Birthday": time.Now()},
// 	})
// }
// func main() {
// 	createTable()
// 	inset1()
// 	inset2()
// 	inset3()
// 	insert4()
// }
