// // 缓冲区是具有读取和写入方法的可变大小的字节缓冲区。Buffer的零值是准备使用的空缓冲区
// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"io"
// )

// func testVar() {
// 	var b bytes.Buffer                     // 定义以一个buffer变量，不用初始化，可以直接使用
// 	bn := new(bytes.Buffer)                // 使用new返回BUffer变量指针
// 	b1 := bytes.NewBuffer([]byte("hello")) //从一个[]byte切片，构造一个Buffer
// 	b2 := bytes.NewBufferString("hello world")

// 	fmt.Printf("b: %v\n", b)
// 	fmt.Printf("bn: %T\n", bn)
// 	fmt.Printf("b1: %v\n", b1)
// 	fmt.Printf("b2: %v\n", b2)
// }

// func test1() {
// 	var b bytes.Buffer
// 	// n, _ := b.WriteString("hello")
// 	// fmt.Printf("n: %v\n", n) // 5
// 	// fmt.Printf("b: %s\n", b.Bytes()[:n])
// 	// b.Write([]byte("hi"))
// 	// fmt.Printf("b: %v\n", string(b.Bytes()))
// 	b.WriteByte('.')
// 	fmt.Printf("b: %v\n", string(b.Bytes()))
// }

// func test2() {
// 	var b = bytes.NewBufferString("hello world")
// 	b1 := make([]byte, 2)
// 	for {
// 		n, err := b.Read(b1)
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Printf("b1: %T %s\n", b1, b1[0:n])
// 		fmt.Printf("n: %v\n", n)
// 	}

// }

// func main() {
// 	test2()
// }
