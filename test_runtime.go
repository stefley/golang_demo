// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"time"
// )

// func show(msg string) {
// 	for i := 0; i < 2; i++ {
// 		fmt.Printf("msg: %v\n", msg)
// 	}
// }

// // Goexit 退出协程
// func show2() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 		if i >= 5 {
// 			runtime.Goexit()
// 		}
// 	}
// }

// func a() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("a: %v\n", i)
// 	}
// }
// func b() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("b: %v\n", i)
// 	}
// }
// func main() {

// Gosched
// go show("java")

// for i := 0; i < 2; i++ {
// 	runtime.Gosched() // 让出cpu执行权限
// 	fmt.Println("main....")
// }

// Goexit
// go show2()
// time.Sleep(time.Second)
// fmt.Println("end....")

// GOMAXPROCS
// 	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
// 	runtime.GOMAXPROCS(2)
// 	go a()
// 	go b()
// 	time.Sleep(time.Second)
// }
