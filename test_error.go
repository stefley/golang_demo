// package main

// import (
// 	"errors"
// 	"fmt"
// 	"time"
// )

// /*

// 	type error interface {
// 		Error() string
// 	}
// */

// func check(s string) error {
// 	if s == "" {
// 		return errors.New("字符串不能为空")
// 	} else {
// 		return nil
// 	}
// }

// // 自定义错误
// type MyErr struct {
// 	When time.Time
// 	What string
// }

// func (e MyErr) Error() string {
// 	return fmt.Sprintf("%v: %v", e.When, e.What)
// }
// func oops() error {
// 	return MyErr{
// 		time.Date(1989, 3, 15, 22, 30, 0, 0, time.UTC),
// 		"the file sytem has gone away",
// 	}
// }

// func test1() {

// 	err := check("")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// func test2() {
// 	err := oops()
// 	fmt.Printf("err: %T %v\n", err, err)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }
// func main() {
// 	test2()
// }
