// package main

// import (
// 	"fmt"
// 	"os"
// )

// // 创建文件
// func createFile() {
// 	f, err := os.Create("a.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f.Name(): %v\n", f.Name())
// 	}
// }

// // 创建目录
// func createDir() {
// 	err := os.Mkdir("eg", os.ModePerm) // 最高权限777创建目录
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}

// 	err2 := os.MkdirAll("a/b/c", os.ModePerm)
// 	if err2 != nil {
// 		fmt.Printf("err2: %v\n", err2)
// 	}
// }

// // 删除目录或者文件
// func remove() {
// 	err := os.Remove("a.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}

// 	err2 := os.RemoveAll("a")
// 	if err2 != nil {
// 		fmt.Printf("err2: %v\n", err2)
// 	}
// }

// // 获得和修改当前工作目录
// func wd() {
// 	dir, _ := os.Getwd()
// 	fmt.Printf("dir: %v\n", dir)

// 	// 修改当前工作目录
// 	os.Chdir("d:/")
// 	dir, _ = os.Getwd()
// 	fmt.Printf("dir: %v\n", dir)

// 	// 临时目录
// 	s := os.TempDir()
// 	fmt.Printf("s: %v\n", s)
// }

// // 重命名
// func rename() {
// 	err := os.Rename("test.txt", "test2.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	}
// }

// // 读文件
// func readFile() {
// 	b, _ := os.ReadFile("test2.txt")
// 	fmt.Printf("b: %v\n", b)            // b: [104 105]
// 	fmt.Printf("b: %v\n", string(b[:])) // hi
// }

// // 写文件
// func writeFile() {
// 	os.WriteFile("test2.txt", []byte("hi"), os.ModePerm)
// }
// func main() {
// 	createFile()
// 	createDir()
// 	remove()
// 	wd()
// 	rename()
// 	readFile()
// 	writeFile()
// }