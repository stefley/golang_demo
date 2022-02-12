// // golong 只有for循 可以通过break、goto、return、panic语句强制退出循环
// package main

// import "fmt"

// func main() {
// 	// for1()
// 	// for2()
// 	// for3()
// 	for4()
// }

// func for1() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}
// }
// func for2() {
// 	i := 1
// 	for ; i <= 10; i++ {
// 		fmt.Printf("i: %v\n", i)
// 	}
// }

// func for3() {
// 	i := 1
// 	for i <= 10 {
// 		fmt.Printf("i: %v\n", i)
// 		i++
// 	}
// }

// // 永真循环
// func for4() {
// 	for {
// 		fmt.Println("run.....")
// 	}
// }
