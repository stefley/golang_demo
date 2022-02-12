// package main

// func main() {
// // 获得所有环境变量
// s := os.Environ()
// fmt.Printf("s: %v\n", s)
// // 获得某个环境变量
// s2 := os.Getenv("GOPATH")
// fmt.Printf("s2: %v\n", s2)

// s3, b := os.LookupEnv("GOPATH")
// if b {
// 	fmt.Printf("s3: %v\n", s3)
// }

// 设置环境变量
// os.Setenv("env1", "env1")

// 替换
// os.Setenv("TESTNAME", "gopher")
// os.Setenv("BURROW", "/usr/gopher")

// fmt.Printf("os.ExpandEnv(\" lives in .\"): %v\n", os.ExpandEnv(" lives in ."))

// 清除所有环境变量
// os.Clearenv()
// }
