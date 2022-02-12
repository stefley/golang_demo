// // sort 包主要针对[]int []float64 []string 以及其他自定义切片的排序
// /*
// 	type IntSlice struct
// 	type Float64Slice []float64
// 	type StringSlice []string
// */
// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func test1() {
// 	s := []int{2, 4, 1, 3}
// 	sort.Ints(s)
// 	fmt.Printf("s: %v\n", s)
// 	b := sort.IntsAreSorted(s)
// 	fmt.Printf("b: %v\n", b)
// }

// // 二分查找元素下标
// func test2() {
// 	s := []int{3, 5, 1, 4, 2, 6, 7}
// 	i := sort.SearchInts(s, 6)
// 	fmt.Printf("i: %v\n", i)
// }

// func test3() {
// 	// type NewInts []uint
// 	// 	func (n NewInts) Len() int {
// 	// 		return len(n)
// 	// 	}
// 	// 	func (n NewInts) Less(i, j int) bool {
// 	// 		fmt.Println(i,j, n[i] < n[j], n)
// 	// 		return n[i] < n[j]
// 	// 	}
// 	// 	func (n NewInts) Swap(i, j int) {
// 	// 		n[i],n[j] = n[j],n[i]
// 	// 	}

// 	// 	n := []uint{1,3,2}
// 	// 	sort.Sort(NewInts(n))
// 	// fmt.Println(n)
// }

// func test4() {
// 	f := []float64{1.2, 1.3, 1.0, 1.1}
// 	sort.Float64s(f)
// 	fmt.Printf("f: %v\n", f)
// }

// func test5() {
// 	// 字符串排序，先比较高位，相同再比较低位
// 	ls := sort.StringSlice{"100", "22", "33", "55", "11"}
// 	fmt.Println("ls:", ls)
// 	sort.Strings(ls)
// 	fmt.Printf("ls: %v\n", ls)

// 	ls1 := sort.StringSlice{
// 		"啊",
// 		"波",
// 		"的",
// 		"次",
// 		"特",
// 	}
// 	for _, v := range ls1 {
// 		fmt.Printf("v: %v\n", []byte(v))
// 	}
// 	sort.Strings(ls1)
// 	fmt.Printf("ls1: %v\n", ls1)
// }

// func test6() {
// 	// 2维切片
// 	// type testSlice [][]int
// 	// func (l testSlice) Len() int { return len(l)}
// 	// func (l testSlice) Swap(i, j int) { l[i],l[j] = l[j],l[i]}
// 	// func (l testSlice) Less(i, j int) bool { return l[i][1] < l[j][1]}

// }

// func test7() {
// 	type testSlice []map[string]float64

// 	// 实现Len Less Swap方法
// }

// func test8() {
// 	type People struct {
// 		Name string
// 		Age  int
// 	}
// 	type testSlice []People
// 	// 实现Len Less Swap方法
// }
// func main() {
// 	// test4()
// 	test5()
// }
