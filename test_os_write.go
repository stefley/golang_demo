// package main

// import "os"

// // r4 w2 x1
// func write() {
// 	// os.O_RDWR 读写 os.O_APPEND 添加 os.O_TRUNC 覆盖
// 	// f, _  := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND)
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
// 	f.Write([]byte("麻了!"))
// 	f.Close()
// }
// func writeString() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR|os.O_APPEND, 0755)
// 	f.WriteString("hello golang")
// 	f.Close()
// }
// func writeAt() {
// 	f, _ := os.OpenFile("a.txt", os.O_RDWR, 0755)
// 	f.WriteAt([]byte("zzz..."), 10)
// 	f.Close()
// }
// func main() {
// 	write()
// 	writeString()
// 	writeAt()
// }
