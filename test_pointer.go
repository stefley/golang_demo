// Go语言中的函数传参都是值拷贝，当我们想要修改某个变量的时候，我们可以创建一个指向该变量地址的指针变量。传递数据使用指针，而无需拷贝数据
// 类型指针不能进行偏移和运算
// Go语言的指针操作非常简单，只需要记住两个符号：&(取地址)和*(根据地址取值)
/*
指针地址和指针类型
每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。
Go语言中的值类型(int、float、bool、string、array、struct)都有对应的指针类型，如*int、*int64、*string等
*/
// package main

// import "fmt"

// func main() {
// 	var ip *int
// 	fmt.Printf("ip: %v\n", ip) // nil
// 	fmt.Printf("ip: %T\n", ip) // *int

// 	var i int = 100
// 	ip = &i
// 	fmt.Printf("ip: %v\n", ip)  // 0xc0000120f0
// 	fmt.Printf("ip: %v\n", *ip) // 100

// 	var sp *string
// 	var s = "hello"
// 	sp = &s
// 	fmt.Printf("sp: %v\n", sp)  // 0xc000044230
// 	fmt.Printf("sp: %v\n", *sp) // "hello"

// 	var bp *bool
// 	var b bool = true
// 	bp = &b
// 	fmt.Printf("bp: %v\n", bp) // 0xc0000aa098
// 	fmt.Printf("bp: %v\n", *bp) // true
// }
