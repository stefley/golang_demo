// package main

// import (
// 	"fmt"
// 	"time"
// )

// func test1() {
// 	now := time.Now()
// 	fmt.Printf("now: %v\n", now) //now: 2022-01-24 22:17:16.0939576 +0800 CST m=+0.007616401

// 	i := now.Year()
// 	m := now.Month()
// 	i2 := now.Day()
// 	i3 := now.Hour()
// 	i4 := now.Minute()
// 	i5 := now.Second()
// 	fmt.Printf("i: %v\n", i)
// 	fmt.Printf("m: %v\n", m)
// 	fmt.Printf("i2: %v\n", i2)
// 	fmt.Printf("i3: %v\n", i3)
// 	fmt.Printf("i4: %v\n", i4)
// 	fmt.Printf("i5: %v\n", i5)
// 	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", i, m, i2, i3, i4, i5) // %02d 2个宽度不足补0
// }

// func test2() {
// 	// 时间戳 1970.1.1 08:00:00GMT
// 	now := time.Now()
// 	// TimeStamp type:int64, TimeStamp:1643034232
// 	fmt.Printf("TimeStamp type:%T, TimeStamp:%v", now.Unix(), now.Unix())
// 	// 纳秒数
// 	fmt.Printf("now.UnixNano(): %v\n", now.UnixNano())
// }

// func test3() {
// 	// 时间戳与时间格式转换
// 	timestamp := time.Now().Unix()
// 	timeObj := time.Unix(timestamp, 0)
// 	fmt.Printf("timeObj: %v\n", timeObj)
// 	fmt.Printf("timeObj.Year(): %v\n", timeObj.Year())
// }

// // func add(h, m, s, mls, msc, ns, time.Duration) {
// // 	now := time.Now()
// // 	fmt.Printf("now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls + time.Microsecond*msc + time.Nanosecond*ns): %v\n", now.Add(time.Hour*h+time.Minute*m+time.Second*s+time.Millisecond*mls+time.Microsecond*msc+time.Nanosecond*ns))
// // }

// func testFormat() {
// 	// 2006 1234
// 	t := time.Now()
// 	// t.Format("2006-01-02: 15:04"): 2022-01-24: 22:37
// 	fmt.Printf("t.Format(\"2006-01-02: 15:04\"): %v\n", t.Format("2006-01-02: 15:04"))
// }

// func testParse() {
// 	t := time.Now()
// 	fmt.Printf("t: %v\n", t)
// 	// 加载时区
// 	l, err := time.LoadLocation("Asia/Shanghai")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	t2, err2 := time.ParseInLocation("2006/01/02 15:04:05", "2022/01/24 22:37:22", l)
// 	if err2 != nil {
// 		fmt.Println(err2)
// 		return
// 	}
// 	fmt.Printf("t2: %v\n", t2)
// 	fmt.Printf("t2.Sub(t): %v\n", t2.Sub(t))
// }
// func main() {

// }
