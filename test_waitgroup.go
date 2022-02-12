// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func showMsg(i int) {
// 	// 每执行完一个协程任务，等待组-1
// 	// defer wg.Add(-1)
// 	defer wg.Done()
// 	fmt.Printf("i: %v\n", i)
// }
// func main() {
// 	for i := 0; i < 10; i++ {
// 		// 启动一个协程来执行
// 		go showMsg(i)
// 		// 每启动一个协程，等待组+1
// 		wg.Add(1)
// 	}

// 	wg.Wait()
// 	// 主协程
// 	fmt.Println("end....")
// }
