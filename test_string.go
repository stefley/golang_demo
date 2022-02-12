// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	// 字符串字面量
// 	// var s string = "hello world"
// 	// var s1 = "hello world"
// 	// s3 := "hello world"

// 	// fmt.Printf("s: %v\n", s)
// 	// fmt.Printf("s1: %v\n", s1)
// 	// fmt.Printf("s3: %v\n", s3)

// 	// s4 := `
// 	// 	line1
// 	// line2
// 	// line3
// 	// `

// 	// fmt.Printf("s4: %v\n", s4)

// 	// 字符串连接
// 	// s1 := "tom"
// 	// s2 := "20"
// 	// msg := s1 + s2
// 	// fmt.Printf("msg: %v\n", msg)
// 	// msg2 := fmt.Sprintf("%s:%s \n", s1, s2)
// 	// fmt.Printf("msg2: %v\n", msg2)
// 	// // strings.Join
// 	// msg3 := strings.Join([]string{s1, s2}, ",")
// 	// fmt.Printf("msg3: %v\n", msg3)

// 	// // buffer.WriteString
// 	// var buffer bytes.Buffer
// 	// buffer.WriteString("jay")
// 	// buffer.WriteString(":")
// 	// buffer.WriteString("20")
// 	// fmt.Printf("buffer.String(): %v\n", buffer.String())

// 	// 转移字符
// 	// \r 回车
// 	// \n 换行 下行同位置
// 	// \t 制表符
// 	// \' 单引号
// 	// \" 双引号
// 	// \\ 反斜杠

// 	// 切片操作
// 	str := "hello world"
// 	n := 3
// 	m := 5
// 	fmt.Println(str[n])              // 获取n位置的原始字节 108
// 	fmt.Printf("s[a]: %c\n", str[n]) // l
// 	fmt.Println(str[n:m])            // 截取n到m-1字符串  lo
// 	fmt.Println(str[n:])             // 截取n到len(s)-1的字符串 lo world
// 	fmt.Println(str[:m])             // 截取0到m-1 hello

// 	// 字符串方法
// 	// len
// 	length := len(str)
// 	fmt.Printf("length: %v\n", length)
// 	// strings.Split
// 	var strArr []string = strings.Split(str, " ")
// 	fmt.Printf("strArr: %v\n", strArr)
// 	// strings.Contains 是否包含
// 	var hasHello bool = strings.Contains(str, "hello")
// 	fmt.Printf("hasHello: %v\n", hasHello)
// 	// strings.ToLower 转小写
// 	str2 := "HELLO WORLD"
// 	fmt.Printf("strings.ToLower(str2): %v\n", strings.ToLower(str2))
// 	// strings.ToUpper 转大写
// 	fmt.Printf("strings.ToUpper(str): %v\n", strings.ToUpper(str))
// 	// strings.HasPrefix 是否有前缀
// 	fmt.Printf("strings.HasPrefix(str2, \"HE\"): %v\n", strings.HasPrefix(str2, "HE"))
// 	// strings.HasSuffix 是否有后缀
// 	fmt.Printf("strings.HasSuffix(str, \"ld\"): %v\n", strings.HasSuffix(str, "ld"))
// 	// strings.Index 查找字符串中位置
// 	fmt.Printf("strings.Index(str, \"ll\"): %v\n", strings.Index(str, "ll"))
// 	// strings.LastIndex 最后出现位置
// 	fmt.Printf("strings.LastIndex(str, \"l\"): %v\n", strings.LastIndex(str, "l"))
// }
