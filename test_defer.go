// defer语句会将其后面跟随的语句进行延迟处理.在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
// 先被defer的语句最后执行，最后defer的最先执行
/**
特性
1.关键字defer用于注册延迟调用
2.这些调用知道return前才执行，因此可以用来做资源清理
3.多个defer语句，按先进后出的方式执行
4.defer语句中的变量，在defer声明时就决定了
用途
1.关闭文件句柄
2.锁资源释放
3.数据库连接释放
*/
// package main

// import "fmt"

// 4 3 2 1 start
// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println("step1")
// 	defer fmt.Println("step2")
// 	defer fmt.Println("step3")
// 	defer fmt.Println("step4")
// }
