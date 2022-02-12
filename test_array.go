// 数组是相同数据类型的一组数据的集合，数组一旦定义长度不能修改，数组可以通过下标来访问元素
// package main

// import "fmt"

// func test1() {
// 	var a1 [2]int
// 	var a2 [3]string
// 	var a3 [2]bool
// 	fmt.Printf("a1: %T\n", a1) // [2]int
// 	fmt.Printf("a2: %T\n", a2) // [3]string
// 	fmt.Printf("a3: %T\n", a3) // [2]bool
// 	fmt.Printf("a1: %v\n", a1) // [0 0]
// 	fmt.Printf("a2: %v\n", a2) // [  ]
// 	fmt.Printf("a3: %v\n", a3) // [false false]
// }

// // 数组初始化
// func test2() {
// 	var a1 = [2]int{1, 2}
// 	fmt.Printf("a1: %v\n", a1)
// 	var a2 = [2]string{"hello", "world"}
// 	fmt.Printf("a2: %v\n", a2)
// 	var a3 = [3]bool{true, false}
// 	fmt.Printf("a3: %v\n", a3)
// }

// // 数组初始化, 省略数组长度,自动推断长度
// func test3() {
// 	var a1 = [...]int{2, 3, 4}
// 	fmt.Printf("a1: %v\n", a1)
// 	fmt.Printf("a1: %T\n", a1) // [3]int
// }

// // 指定索引值的方式来初始化
// func test4() {
// 	var a1 = [...]int{0: 1, 2: 2}
// 	fmt.Printf("a1: %v\n", a1)                     // [1 0 2]
// 	var a2 = [...]bool{2: true, 5: false, 3: true} // [false false true true false false]
// 	fmt.Printf("a2: %v\n", a2)
// }

// // 修改数组
// func test5() {
// 	var a1 = [...]int{1, 2, 3}
// 	fmt.Printf("a1: %v\n", a1)
// 	a1[0] = 100
// 	fmt.Printf("a1: %v\n", a1)
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// 	test4()
// 	test5()
// }
