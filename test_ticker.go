// Ticker可以周期执行

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {

// ticker := time.NewTicker(time.Second)
// counter := 1
// // 每秒执行一次
// for _ = range ticker.C {
// 	fmt.Println("ticker")
// 	counter++
// 	if counter >= 5 {
// 		ticker.Stop()
// 		break
// 	}
// }

// 	ticker := time.NewTicker(time.Second)
// 	chanInt := make(chan int)

// 	go func() {
// 		for _ = range ticker.C {
// 			select {
// 			case chanInt <- 1:
// 				fmt.Println("send 1")
// 			case chanInt <- 2:
// 				fmt.Println("send 2")
// 			case chanInt <- 3:
// 				fmt.Println("send 3")
// 			}
// 		}
// 	}()

// 	sum := 0
// 	for v := range chanInt {
// 		fmt.Printf("收到: %v\n", v)
// 		sum += v
// 		if sum >= 10 {
// 			fmt.Printf("sum: %v\n", sum)
// 			break
// 		}
// 	}
// }
