// print系列(单纯打印) panic系列(打印日志，抛出panic异常) fatal系列(打印日志，强制结束程序os.Exit(1), defer函数不会执行)

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func test1() {
// 	log.Print("my log")          // 2022/01/21 23:10:04 my log
// 	log.Printf("my log:%d", 100) // 2022/01/21 23:12:19 my log:100
// 	name := "tom"
// 	age := 20
// 	log.Println(name, " ", +age)
// }

// func test2() {
// 	// 	2022/01/21 23:20:58 my log
// 	// 2022/01/21 23:20:58 my panic
// 	// panic 结束后执行
// 	// panic: my panic
// 	defer fmt.Println("panic 结束后执行")
// 	log.Print("my log")
// 	log.Panic("my panic")

// 	fmt.Println("end....")
// }

// func test3() {
// 	// 	2022/01/21 23:22:43 my log
// 	// 2022/01/21 23:22:43 fatal....
// 	// exit status 1
// 	defer fmt.Println("defer ....")
// 	log.Print("my log")
// 	log.Fatal("fatal....")

// 	fmt.Println("end ....")
// }

// func test4() {
// 	config := log.Flags()
// 	fmt.Printf("config: %v\n", config) // 3
// }

// var logger *log.Logger

// func init() {
// 	/* 	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
// 	   	// 设置日志前缀
// 	   	log.SetPrefix("Mylog: ")
// 	   	// 日志输出位置
// 	   	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
// 	   	if err != nil {
// 	   		log.Fatal("日志文件错误")
// 	   	}
// 	   	log.SetOutput(f) */

// 	// 自定义logger
// 	f, err := os.OpenFile("a.log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0664)
// 	if err != nil {
// 		log.Fatal("日志错误")
// 	}
// 	logger = log.New(f, "MYLOG: ", log.Ldate|log.Ltime|log.Lshortfile)
// }
// func main() {
// test1()
// test2()
// test3()
// test4()
// log.Print("my log ...")

// logger.Print("logger log.....")
// logger.Printf("logger printf: %v", "test")
// }
