// package main

// func main() {
// timer := time.NewTimer(time.Second * 2)
// fmt.Printf("time.Now(): %v\n", time.Now())

// t1 := <-timer.C // 阻塞的，直到时间到了
// fmt.Printf("t1: %v\n", t1)

// timer := time.NewTimer(time.Second*2)
// <-timer.C // 阻塞

// time.Sleep(time.Second*2) // 睡眠阻塞

// <-time.After(time.Second * 2)

// timer := time.NewTimer(time.Second)

// go func() {
// 	<-timer.C
// 	fmt.Println("func....")
// }()

// s := timer.Stop() // 停掉定时
// if s {
// 	fmt.Println("timer stop")
// }

// fmt.Println("before...")
// timer := time.NewTimer(time.Second * 5)
// timer.Reset(time.Second * 2) // 重新设置定时为2s
// <-timer.C
// fmt.Println("after....")
// }
