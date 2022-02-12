// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var i int = 100

// var wg sync.WaitGroup
// var addCount int
// var subCount int

// var lock sync.Mutex // 互斥锁

// func add() {
// 	time.Sleep(time.Millisecond * 10)
// 	defer wg.Done()
// 	lock.Lock()
// 	i += 1
// 	addCount++
// 	fmt.Printf("i++: %v\n", i)
// 	lock.Unlock()
// }

// func sub() {
// 	time.Sleep(time.Millisecond * 2)
// 	defer wg.Done()
// 	lock.Lock()
// 	i--
// 	subCount++
// 	fmt.Printf("i--: %v\n", i)
// 	lock.Unlock()
// }
// func main() {
// 	for i := 0; i < 100; i++ {
// 		wg.Add(1)
// 		go add()
// 		wg.Add(1)
// 		go sub()
// 	}

// 	wg.Wait()
// 	fmt.Printf("addCount: %v\n", addCount)
// 	fmt.Printf("subCount: %v\n", subCount)
// 	fmt.Printf("end i: %v\n", i)
// }
