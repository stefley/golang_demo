// package main

// import "fmt"

// // append
// func test1() {
// 	s := []int{}
// 	i := append(s, 10)
// 	fmt.Printf("i: %v\n", i)
// 	s1 := []int{4, 5, 6}
// 	i2 := append(s, s1...)
// 	fmt.Printf("i2: %v\n", i2)
// }

// // len
// func test2() {
// 	s := "Hello world"
// 	fmt.Printf("len(s): %v\n", len(s))
// }

// // print println
// func test3() {
// 	name := "tom"
// 	age := 20
// 	print(name, " ", age, "\n")
// 	fmt.Println("-------------")
// 	println(name, " ", age)
// 	println(name, " ", age)
// }

// // panic 发生panic异常前如果有defer先执行defer
// func test4() {
// 	defer fmt.Println("panic defer")
// 	panic("panic...")
// 	fmt.Println("end ....")
// }

// // new make
// /*
// 	只能分配和初始化类型为slice、map、chan的数据，new可以分配任意类型的数据
// 	new分配返回的是指针：*T,make 返回引用：T
// 	new分配的空间被清零，make分配后进行初始化
// */
// func test5() {
// 	b := new(bool)
// 	fmt.Printf("b: %T %v\n", b, *b) // b: *bool false
// 	i := new(int)
// 	fmt.Printf("i: %T %v\n", i, *i) // i: *int 0

// 	s := new(string)
// 	fmt.Printf("s: %T %v\n", s, *s) // s: *string ""
// }
// func test6() {
// 	var p *[]int = new([]int)
// 	var v []int = make([]int, 100)

// 	fmt.Printf("p: %v\n", p)
// 	fmt.Printf("v: %v\n", v)
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// 	test4()
// 	test5()
// 	test6()
// }
