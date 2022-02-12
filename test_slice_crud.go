// package main

// import "fmt"

// // add  append
// func test1() {
// 	var s1 = []int{}
// 	s1 = append(s1, 100)
// 	s1 = append(s1, 200)
// 	fmt.Printf("s1: %v\n", s1)
// 	s2 := make([]string, 3)
// 	s2[1] = "tom"
// 	s2 = append(s2, "jaja")
// 	fmt.Printf("s2: %v\n", s2)
// 	fmt.Printf("len(s2): %v\n", len(s2))
// 	fmt.Printf("cap(s2): %v\n", cap(s2))
// }

// // delete
// func test2() {
// 	var s1 = []int{1, 2, 3, 4}
// 	// 删除i append(s[0,i], s[i+1:]...)
// 	s1 = append(s1[:2], s1[3:]...)
// 	fmt.Printf("s1: %v\n", s1)
// }

// // update
// func test3() {
// 	var s1 = []string{"z", "y", "q"}
// 	fmt.Printf("s1: %v\n", s1)
// 	s1[1] = "z"
// 	fmt.Printf("s1: %v\n", s1)

// }

// // query
// func test4() {
// 	var s1 = []int{1, 2, 3, 4}
// 	var key = 2
// 	for i, v := range s1 {
// 		if v == key {
// 			fmt.Printf("i: %v\n", i)
// 		}
// 	}
// }

// // copy
// func test5() {
// 	var s1 = []int{1, 2, 3, 4}
// 	// 引用类型
// 	var s2 = s1
// 	s2[0] = 100
// 	fmt.Printf("s1: %v\n", s1)
// 	fmt.Printf("s2: %v\n", s2)
// 	var s3 = make([]int, 4)
// 	copy(s3, s1)
// 	fmt.Printf("s3: %v\n", s3)
// 	s3[2] = 200
// 	fmt.Printf("s3: %v\n", s3)
// 	fmt.Printf("s1: %v\n", s1)
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// 	test4()
// 	test5()
// }
