// package main

// import "fmt"

// func test1() {
// 	var a1 [2]int
// 	a1[0] = 100
// 	a1[1] = 120
// 	fmt.Printf("a1[0]: %v\n", a1[0])
// 	fmt.Printf("a1[1]: %v\n", a1[1])
// 	fmt.Println("-----------------------")

// 	a1[0] = 1000
// 	a1[1] = 2000
// 	fmt.Printf("a1: %v\n", a1)
// 	fmt.Printf("a1[1]: %v\n", a1[1])

// 	// 数组长度越界
// 	// a1[3] = 3000  error

// }

// // 数组长度
// func test2() {
// 	var a1 = [3]int{1, 2, 3}
// 	fmt.Printf("len(a1): %v\n", len(a1))
// 	var a2 = [...]int{1, 2, 3, 4}
// 	fmt.Printf("len(a2): %v\n", len(a2))
// }

// // 数组遍历
// func test3() {
// 	// 根据长度和下标
// 	var a1 = [3]int{1, 2, 3}
// 	for i := 0; i < len(a1); i++ {
// 		fmt.Printf("a1[i]: %v, %v\n", a1[i], i)
// 	}
// }
// func test4() {
// 	// for range 遍历
// 	var a1 = [...]int{1, 2, 3}
// 	for i, v := range a1 {
// 		fmt.Printf("i: %v, v: %v\n", i, v)
// 	}
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// 	test4()
// }
