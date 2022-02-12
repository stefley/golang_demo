// map是一种key:value的键值对数据结构容器。map内部实现是哈希表（hash） map是无序的
// package main

// import "fmt"

// func test1() {
// 	m1 := make(map[string]string)
// 	m1["name"] = "tom"
// 	m1["age"] = "10"
// 	fmt.Printf("m1: %v\n", m1)
// 	// 类型声明
// 	var m2 map[string]int
// 	// 初始化
// 	m2 = make(map[string]int)
// 	m2["age"] = 10
// 	m2["tel"] = 110
// 	fmt.Printf("m2: %v\n", m2)
// }

// //初始化map
// func test2() {
// 	var m1 = map[string]string{"name": "tom", "age": "22", "email": "tom@gmail.com"}
// 	fmt.Printf("m1: %v\n", m1)
// 	m2 := make(map[string]string)
// 	m2["name"] = "jack"
// 	m2["age"] = "20"
// 	fmt.Printf("m2: %v\n", m2)
// }

// // 访问map
// func test3() {
// 	m1 := map[string]string{"name": "tom", "age": "20", "email": "tom@gmail.com"}
// 	var key = "name"
// 	var vlaue = m1[key]
// 	fmt.Printf("vlaue: %v\n", vlaue)
// }

// // 判断某个key是否存在 value,ok := map[key]
// func test4() {
// 	var m1 = map[string]string{"name": "tom", "age": "22", "email": "tom@gmail.com"}
// 	var k1 = "name"
// 	var k2 string = "age1"
// 	v, ok := m1[k1]
// 	fmt.Printf("v: %v, ok: %v\n", v, ok) // tom true
// 	fmt.Println("------------------")
// 	v, ok = m1[k2]
// 	fmt.Printf("v: %v, ok: %v\n", v, ok) //  false
// }

// // 遍历map
// func test5() {
// 	m := map[string]string{"name": "tom", "age": "20", "tel": "155xxxxxxx"}
// 	for key, value := range m {
// 		fmt.Printf("key: %v value: %v\n", key, value)
// 	}
// }
// func main() {
// 	test1()
// 	test2()
// 	test3()
// 	test4()
// 	test5()
// }
