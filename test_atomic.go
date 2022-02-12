// package main

// import (
// 	"fmt"
// 	"sync/atomic"
// 	"time"
// )

// var i = 100
// var lock sync.Mutex // 互斥锁实现协程同步

// func add() {
// 	lock.Lock()
// 	i++
// 	lock.Unlock()
// }

// func sub() {
// 	lock.Lock()
// 	i--
// 	lock.Unlock()
// }

// var i int32 = 100

// cas compare and swap 比较并且交换 old new

// func add() {
// 	atomic.AddInt32(&i, 1)
// }
// func sub() {
// 	atomic.AddInt32(&i, -1)
// }
// func main() {

// 	for i := 0; i < 100; i++ {
// 		go add()
// 		go sub()
// 	}

// 	time.Sleep(time.Second * 2)
// 	fmt.Printf("i: %v\n", i)
// }
