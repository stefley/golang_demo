// package main

// import "fmt"

// // 切片切分
// func test1() {
// 	var s1 = []int{1, 2, 3, 4, 5, 6}
// 	s2 := s1[:3]
// 	fmt.Printf("s2: %T, %v\n", s2, s2)
// 	s3 := s2[:]
// 	fmt.Printf("s3: %v\n", s3)
// 	s4 := s1[1:4]
// 	fmt.Printf("s4: %v\n", s4)
// }

// // 数组切分
// func test2() {
// 	var a1 = [...]int{1, 2, 3, 4, 5, 6}
// 	var s1 = a1[:]
// 	fmt.Printf("s1: %T, %v\n", s1, s1)
// 	s2 := a1[2:5]
// 	fmt.Printf("s2: %v\n", s2)
// 	s3 := a1[3:]
// 	fmt.Printf("s3: %v\n", s3)
// }

// func main() {
// 	test1()
// 	test2()
// }
