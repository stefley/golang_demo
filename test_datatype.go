// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	// var name string = "tom"
// 	// age := 20
// 	// b := true
// 	// fmt.Printf("%T\n", name)
// 	// fmt.Printf("%T\n", age)
// 	// fmt.Printf("%T\n", b)

// 	// a := 100
// 	// p := &a
// 	// fmt.Printf("%T\n", p)
// 	// fmt.Printf("p: %v\n", p)

// 	// ...省略长度
// 	// a := [...]int{1, 2}
// 	// fmt.Printf("%T\n", a)

// 	// a := []int{1, 2, 3}
// 	// fmt.Printf("%T\n", a)
// 	// b1 := false
// 	// b2 := true
// 	// var b3 bool = false
// 	// var (
// 	// 	b4 = true
// 	// )
// 	// fmt.Printf("b1: %v\n", b1)
// 	// fmt.Printf("b2: %v\n", b2)
// 	// fmt.Printf("b3: %v\n", b3)
// 	// fmt.Printf("b4: %v\n", b4)
// 	// age := 16
// 	// if age >= 18 {
// 	// 	fmt.Println("成年")
// 	// } else {
// 	// 	fmt.Println("未成年")
// 	// }
// 	// count := 10
// 	// for i := 0; i < count; i++ {
// 	// 	fmt.Printf("i: %v\n", i)
// 	// }

// 	// age := 18
// 	// gender := "男"

// 	// if age >= 18 && gender == "男" {
// 	// 	fmt.Println("成年男子")
// 	// }

// 	// var i8 int8
// 	// var i16 int16
// 	// var i32 int32
// 	// var i64 int64
// 	// var u8 uint8
// 	// var u16 uint16
// 	// var u32 uint32
// 	// var u64 uint64

// 	// fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
// 	// fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
// 	// fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
// 	// fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)

// 	// fmt.Printf("%T %dB %v~%v\n", u8, unsafe.Sizeof(u8), 0, math.MaxUint8)
// 	// fmt.Printf("%T %dB %v~%v\n", u16, unsafe.Sizeof(u16), 0, math.MaxUint16)
// 	// fmt.Printf("%T %dB %v~%v\n", u32, unsafe.Sizeof(u32), 0, math.MaxUint32)
// 	// fmt.Printf("%T %dB %v~%v\n", u64, unsafe.Sizeof(u64), 0, uint64(math.MaxUint64))

// 	// var f32 float32
// 	// var f64 float64
// 	// fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
// 	// fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)

// 	// var ui uint
// 	// ui = uint(math.MaxUint64) // 再加1会导致overflows错误
// 	// fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, ui)

// 	// var imax, imin int
// 	// imax = int(math.MaxInt64)
// 	// imin = int(math.MinInt64)

// 	// fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)

// 	// var a int = 10
// 	// fmt.Printf("%d \n", a)  // 十进制
// 	// fmt.Printf("%b \n", a)  // 二进制
// 	// fmt.Printf("%o \n", a)  // 八进制
// 	// fmt.Printf("%0x \n", a) // 十六进制

// 	// var b int = 077 //八进制以0开头
// 	// fmt.Printf("%o \n", b)

// 	// var c int = 0xff // 十六进制以0x开头
// 	// fmt.Printf("%x \n", c)
// 	// fmt.Printf("%X \n", c)

// 	// 浮点 flaot
// 	fmt.Printf("%f \n", math.Pi)
// 	fmt.Printf("%.2f \n", math.Pi)

// 	// 复数 complex64 complex128
// 	var c1 complex64
// 	c1 = 1 + 2i
// 	var c2 complex128
// 	c2 = 2 + 3i

// 	fmt.Println(c1)
// 	fmt.Println(c2)
// }
