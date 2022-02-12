// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"strings"
// )

// /*
// 	ReadAll 读取数据返回读到的字节slice
// 	ReadDir 读取一个目录，返回目录入口数组[]os.FileInfo
// 	ReadFile 读取一个文件，返回文件内容（字节slice）
// 	WriteFile 根据文件路径，写入字节slice
// 	TempDir 在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径
// 	TempFile 在一个目录中创建指定前缀名的临时文件，返回os.File
// */

// func readAll() {
// 	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

// 	b, err := ioutil.ReadAll(r)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("string(b): %v\n", string(b))
// 	fmt.Printf("%s\n", b)

// 	// 读取文件，返回File类型，File已实现Read
// 	f, _ := os.Open("a.txt") // File
// 	defer f.Close()
// 	b2, _ := ioutil.ReadAll(f)
// 	fmt.Printf("b2: %s\n", b2)
// }

// func readDir() {
// 	r, err := ioutil.ReadDir("D:\\works_jyd")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("r: %\n", r) // r: [0xc00010a8c0 0xc00010a000 0xc00010a070 ...]

// 	files, err2 := ioutil.ReadDir(".")
// 	if err2 != nil {
// 		log.Fatal(err2)
// 	}
// 	for _, file := range files {
// 		fmt.Println(file.Name())
// 	}
// }

// func readFile() {
// 	b, err := ioutil.ReadFile("a.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Printf("string(b): %v\n", string(b))
// 	fmt.Printf("b: %s\n", b)
// }

// func writeFile() {
// 	ioutil.WriteFile("a.txt", []byte("hello world"), 0664)
// }

// func tempFile() {
// 	content := []byte("temporary file's content")
// 	tempfile, err := ioutil.TempFile("", "example")

// 	fmt.Printf("tempfile.Name(): %v\n", tempfile.Name()) // tempfile.Name(): C:\Users\zyq\AppData\Local\Temp\example664339275
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer os.Remove(tempfile.Name())

// 	if _, err := tempfile.Write(content); err != nil {
// 		log.Fatal(err)
// 	}
// 	if err := tempfile.Close(); err != nil {
// 		log.Fatal(err)
// 	}
// }
// func main() {
// 	readAll()
// 	readDir()
// 	readFile()
// 	writeFile()
// 	tempFile()
// }
