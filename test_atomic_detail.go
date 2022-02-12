// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// )

// func test_add_sub() {
// 	var i int32 = 100

// 	atomic.AddInt32(&i, 1) // 原子操作+1
// 	fmt.Printf("i: %v\n", i)
// 	atomic.AddInt32(&i, -1) // 原子操作-1
// 	fmt.Printf("i: %v\n", i)

// 	var j int64 = 200
// 	atomic.AddInt64(&j, 2)
// 	fmt.Printf("j: %v\n", j)
// }

// func test_load_store() {
// 	var i int32 = 100
// 	atomic.LoadInt32(&i) // 原子操作 读取/载入
// 	fmt.Printf("i: %v\n", i)

// 	atomic.StoreInt32(&i, 200) // 原子操作 存储
// 	fmt.Printf("i: %v\n", i)
// }

// func test_cas() {
// 	// cas
// 	var i int32 = 100
// 	b := atomic.CompareAndSwapInt32(&i, 100, 200)
// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("i: %v\n", i)
// }
// func main() {
// 	// 直接交换
// 	var i int32 = 100
// 	atomic.SwapInt32(&i, 200)
// 	fmt.Printf("i: %v\n", i)
// }
