package error

import "fmt"

// there's a panic function.
// enter defer function.
// catch the panic : panic string
func main() {
	defer func() {
		fmt.Println("enter defer function.")
		if r := recover(); r != nil {
			fmt.Printf("catch the panic : %v", r)
		}
	}()

	panicFunc("GG")

	// 安全调用其他函数
	safeGo("worker 1", func() {
		panicFunc("GG")
	})
}

func panicFunc(s string) {
	fmt.Printf("there's a panic function name %s.\n", s)
	panic("panic string")
}

// ------ recover 函数调用时机 ------

// 错误调用方式： defer testRecover()  ❌ recover 并不会生效
func testFailRecover() (any, error) {
	if r := recover(); r != nil {
		return "", fmt.Errorf("catch panic %v", r)
	}
	return nil, nil
}

// 错误调用方式： defer recover() ❌
func testFailRecover2() {
	defer recover()
}

// ------ 安全调用其他函数 ------
func safeGo(name string, f func()) {

	// defer
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("catch panic: %v", r)
		}
	}()

	// 执行函数
	f()
}
