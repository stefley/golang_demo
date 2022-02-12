// 切片可以理解为可变长度的数组
// package main

// import "fmt"

// func test1() {
// 	// 切片是引用类型，可以使用make函数来创建
// 	var s1 []int
// 	var s2 []string
// 	fmt.Printf("s1: %T\n", s1)
// 	fmt.Printf("s2: %T\n", s2)
// }

// // make
// func test2() {
// 	var s1 = make([]int, 3)
// 	fmt.Printf("s1: %v\n", s1)
// 	s1[0] = 100
// 	fmt.Printf("s1: %v\n", s1)
// }

// // 切片长度和容量
// func test3() {
// 	var s1 = []int{1, 2, 3}
// 	fmt.Printf("len(s1): %v\n", len(s1))
// 	fmt.Printf("cap(s1): %v\n", cap(s1))
// 	fmt.Printf("s1[0]: %v\n", s1[0])
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// }
