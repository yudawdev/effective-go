package functions

import (
	"fmt"
	"time"
)

// 匿名函数

// 基本匿名函数
func anonymousBasic() {

	// 匿名函数 - 定义及立即执行
	func() {
		fmt.Println("Hello from anonymous!")
		time.Sleep(1 * time.Second)
	}()

	// 匿名函数 - 定义及执行分开
	f := func() {
		fmt.Println("Hello from anonymous!")
		time.Sleep(1 * time.Second)
	}

	// 立即执行
	f()

	// 使用协程执行匿名函数
	go f()

	// 开启协程调用匿名函数
	go func() {
		fmt.Println("Hello from goroutines!")
		time.Sleep(1 * time.Second)
	}()

}

// 带普通参数的匿名函数
func anonymousNormalParams() {
	name := "Alice"

	// 匿名函数 - 定义及立即执行
	func(username string) {
		fmt.Printf("Hello, %s!\n", username)
	}(name) // 传入参数 name

	// 开启协程调用匿名函数
	go func(username string) {
		fmt.Printf("Hello, %s! from goroutines\n", username)
	}(name) // 传入参数 name

	time.Sleep(1 * time.Second)
}

// 带多个参数的匿名函数
func anonymousMultiParams() {
	age := 25
	location := "Beijing"

	func(userAge int, userCity string) {
		fmt.Printf("User is %d years old and lives in %s\n", userAge, userCity)
	}(age, location) // 传入多个参数

	go func(userAge int, userCity string) {
		fmt.Printf("User is %d years old and lives in %s from goroutines\n", userAge, userCity)
	}(age, location) // 传入多个参数

	time.Sleep(1 * time.Second)
}

// 带返回值的匿名函数
func anonymousWithResult() {
	i := func(x, y int) int {
		result := x + y
		return result
	}(10, 20)

	fmt.Printf("Result: %d\n", i)

	go func(x, y int) int {
		result := x + y
		fmt.Printf("Result: %d\n", result)
		return result
	}(10, 20)

	time.Sleep(1 * time.Second)
}

// 匿名函数实际的例子 - 固定时间执行
func anonymousFixedDuration() {
	// 每隔一段时间执行任务
	go func(interval time.Duration) {
		ticker := time.NewTicker(interval)
		for range ticker.C {
			fmt.Println("执行定时任务...")
		}
	}(2 * time.Second)
}
