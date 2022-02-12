// // bytes包提供了对字节切片进行读写操作的一系列函数，基本处理函数、比较函数、
// // 后缀检查函数、索引函数、分割函数、大小写处理函数和子切片处理函数
// package main

// import (
// 	"bytes"
// 	"fmt"
// 	"strings"
// )

// func testTransform() {
// 	var i int = 100
// 	var b byte = 10
// 	b = byte(i)
// 	i = int(b)

// 	var s string = "hello world"
// 	b1 := []byte{1, 2, 3}
// 	s = string(b1)
// 	b1 = []byte(s)

// 	fmt.Printf("b1: %v\n", b1)
// }

// func testContains() {
// 	s := "golang.com"
// 	b := []byte(s)
// 	b1 := []byte("golang")
// 	b2 := []byte("GOLANG")

// 	b3 := bytes.Contains(b, b1)
// 	fmt.Printf("b3: %v\n", b3)

// 	b4 := bytes.Contains(b, b2)
// 	fmt.Printf("b4: %v\n", b4)

// 	strings.Contains("hello world", "hello")
// }

// func testCount() {
// 	s := []byte("helloooooo")
// 	sep1 := []byte("h")
// 	sep2 := []byte("l")
// 	sep3 := []byte("o")

// 	fmt.Printf("bytes.Count(s, sep1): %v\n", bytes.Count(s, sep1))
// 	fmt.Printf("bytes.Count(s, sep2): %v\n", bytes.Count(s, sep2))
// 	fmt.Printf("bytes.Count(s, sep3): %v\n", bytes.Count(s, sep3))
// }

// func testRepeat() {
// 	b := []byte("hi")
// 	fmt.Printf("bytes.Repeat(b, 1): %v\n", bytes.Repeat(b, 1)) // hi
// 	fmt.Printf("bytes.Repeat(b, 2): %v\n", bytes.Repeat(b, 2)) // hihi
// }

// func testReplace() {
// 	b := []byte("hello world")
// 	old := []byte("o")
// 	news := []byte("ee")

// 	fmt.Printf("bytes.Replace(b, old, news, 0): %s\n", bytes.Replace(b, old, news, 0))   // hello world
// 	fmt.Printf("bytes.Replace(b, old, news, 1): %s\n", bytes.Replace(b, old, news, 1))   // hellee world
// 	fmt.Printf("bytes.Replace(b, old, news, 2): %s\n", bytes.Replace(b, old, news, 2))   // hellee weerld
// 	fmt.Printf("bytes.Replace(b, old, news, -1): %s\n", bytes.Replace(b, old, news, -1)) // hellee weerld
// }

// func testRunes() {
// 	b := []byte("你好世界！")
// 	r := bytes.Runes(b)
// 	fmt.Printf("len(r): %v\n", len(r)) // 5
// 	fmt.Printf("len(b): %v\n", len(b)) // 15
// }

// func testJoin() {
// 	s1 := [][]byte{[]byte("你好"), []byte("世界")}
// 	sep4 := []byte(",")
// 	fmt.Printf("bytes.Join(s1, sep4): %s\n", bytes.Join(s1, sep4))
// 	sep5 := []byte("#")
// 	fmt.Printf("bytes.Join(s1, sep5): %s\n", bytes.Join(s1, sep5))
// }
// func main() {
// 	testJoin()
// }
