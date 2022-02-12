// for range遍历数组、切片、字符串、map及通道channel
// 数组、切片、字符串返回索引和值
// map返回键和值
// 通道channel只返回通道内的值
// package main

// import "fmt"

// func main() {
// 	f1()
// 	f2()
// 	f3()
// 	f4()
// }

// func f1() {
// 	var a = [5]int{1, 2, 3, 4, 5}
// 	for i, v := range a {
// 		fmt.Printf("i: %v, v: %v\n", i, v)
// 	}
// }
// func f2() {
// 	var slice = []int{1, 2, 3}
// 	for _, v := range slice {
// 		fmt.Printf("v: %v\n", v)
// 	}
// }
// func f3() {
// 	var m = make(map[string]string, 0)
// 	m["name"] = "tom"
// 	m["age"] = "10"
// 	m["email"] = "zyq@gmail.com"
// 	for k, v := range m {
// 		fmt.Printf("k: %v, v: %v\n", k, v)
// 	}
// }

// func f4() {
// 	s := "hello world"
// 	for i, v := range s {
// 		fmt.Printf("i: %v, v: %c\n", i, v)
// 	}
// }
