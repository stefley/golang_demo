// package main

// import (
// 	"fmt"
// 	"os"
// )

// func openClose() {
// 	f, err := os.Open("a.txt")
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err) // The system cannot find the file specified.
// 	} else {
// 		fmt.Printf("f.Name(): %v\n", f.Name())
// 		f.Close()
// 	}

// 	f, err = os.OpenFile("a1.txt", os.O_RDWR|os.O_CREATE, 755) // 文件不存在会创建
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 	} else {
// 		fmt.Printf("f.Name(): %v\n", f.Name())
// 		f.Close()
// 	}

// }

// // 读
// func readOps() {
// 	// 循环读取
// 	// f, _ := os.Open("a.txt")
// 	// for {
// 	// 	buf := make([]byte, 10)
// 	// 	n, err := f.Read(buf)
// 	// 	if err == io.EOF {
// 	// 		break
// 	// 	}
// 	// 	fmt.Printf("n: %v\n", n)
// 	// 	fmt.Printf("string(buf): %v\n", string(buf))
// 	// }
// 	// f.Close()

// 	// 从某个位置开始读取
// 	// f, _ := os.Open("a.txt")
// 	// buf := make([]byte, 3)
// 	// n, _ := f.ReadAt(buf, 3)
// 	// fmt.Printf("n: %v\n", n)
// 	// fmt.Printf("string(buf): %v\n", string(buf))

// 	// 读取目录 返回目录切片
// 	// de, _ := os.ReadDir("a")
// 	// for _, value := range de {
// 	// 	fmt.Printf("value.IsDir(): %v\n", value.IsDir())
// 	// 	fmt.Printf("value.Name(): %v\n", value.Name())
// 	// }

// 	// 定位
// 	f, _ := os.Open("a.txt")
// 	f.Seek(3, 0)
// 	buf := make([]byte, 10)
// 	n, _ := f.Read(buf)
// 	fmt.Printf("n: %v\n", n)
// 	fmt.Printf("string(buf): %v\n", string(buf))
// 	f.Close()
// }

// func main() {
// 	openClose()
// 	readOps()
// }
