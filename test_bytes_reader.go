// // Reader 实现了io.Reader,io.ReaderAt, io.WriteTo, io.Seeker, io.ByteScanner, io.RuneScanner接口，Reader是只读的，可以seek

// package main

// import (
// 	"bytes"
// 	"fmt"
// )

// func testReader() {
// 	data := "123456789"
// 	// 通过[]byte创建Reader
// 	re := bytes.NewReader([]byte(data))
// 	// 返回未读取部分的长度
// 	fmt.Printf("re.Len(): %v\n", re.Len())
// 	// 返回底层数据的长度
// 	fmt.Printf("re.Size(): %v\n", re.Size())
// 	fmt.Println("---------------")

// 	buf := make([]byte, 2)
// 	for {
// 		// 读取数据
// 		n, err := re.Read(buf)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Printf("n: %v\n", string(buf[:n]))
// 		// 返回未读取部分的长度
// 		fmt.Printf("re.Len(): %v\n", re.Len())
// 		// 返回底层数据的长度
// 		fmt.Printf("re.Size(): %v\n", re.Size())
// 	}
// 	fmt.Println("-------------------")

// 	// 设置偏移量，因为上面的操作已经修改了读取位置等信息
// 	re.Seek(0, 0)
// 	// 返回未读取部分的长度
// 	fmt.Printf("re.Len(): %v\n", re.Len())
// 	// 返回底层数据的长度
// 	fmt.Printf("re.Size(): %v\n", re.Size())
// 	for {
// 		// 一个一个字节的读
// 		b, err := re.ReadByte()
// 		if err != nil {
// 			break
// 		}
// 		fmt.Printf("b: %v\n", string(b))
// 	}
// 	fmt.Println("-------------------")

// 	re.Seek(0, 0)
// 	off := int64(0)
// 	for {
// 		// 指定偏移量
// 		n, err := re.ReadAt(buf, off)
// 		if err != nil {
// 			break
// 		}
// 		off += int64(n)
// 		fmt.Println(off, string(buf[:n]))
// 	}
// }
// func main() {
// 	testReader()
// }
