// package main

// import (
// 	"fmt"
// 	"time"
// )

// func showMessage(msg string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("msg: %v\n", msg)
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

// // main主协程  go会启动协程
// func main() {
// 	go showMessage("Java") // go 关键子启动协程
// 	go showMessage("golang")

// 	fmt.Println("main end....") // 主函数（协程）退出，程序结束，其他协程直接结束
// }
