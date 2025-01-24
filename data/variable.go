package data

import (
	"flag"
	"fmt"
	"os"
)

var (
	home   = os.Getenv("HOME")   // 获取 HOME 环境变量
	user   = os.Getenv("USER")   // 获取 USER 环境变量
	gopath = os.Getenv("GOPATH") // 获取 GOPATH 环境变量
)

// 环境变量：读取 和 获取
func getEnv() {
	err := os.Setenv("USER", "gold")
	if err != nil {
		fmt.Println("Error setting env:", err)
		return
	}

	user = os.Getenv("USER") // 更新 user 变量值

	fmt.Printf("home: %s\n", home)
	fmt.Printf("user: %s\n", user)
	fmt.Printf("gopath: %s\n", gopath)

	// 无法手动调用 init()
	//init()
}

//  init 函数是 Go 的特殊函数

// 执行时机：
// - 在包级变量初始化后执行
// - 在导入包都初始化完成后执行
// - 在 main 函数之前执行

// 特性
// - 每个文件可以有多个 init 函数
// - 无参数无返回值
// - 自动执行，不能手动调用
func init() {
	if user == "" {
		user = "gold"
	}

	if home == "/root" {
		home = "/home/" + user
	}

	if gopath == "" {
		gopath = home + "/go"
	}

	// gopath may be overridden by --gopath flag on command line.
	flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}

// 可以设置多个 init 函数
func init() {
	os.Setenv("USER", "goldInit")

	gopath = "gopathINIT"
}
