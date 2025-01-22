package functions

import (
	"errors"
	"fmt"
)

// 返回多个结果
func returnMultiResults() (int8, string, []string, error) {
	return 0, "r1", []string{"1"}, errors.New("new error")
}

// 返回多个结果 - 方法定义返回结果的参数名
func returnMultiDefineResults() (r1 int8, r2 string) {
	r1 = 8
	r2 = "r2"
	return
}

func test() {
	results, s, strings, err := returnMultiResults()

	// println 是 go 中比较低级的打印，输出格式无法控制
	// 0 r1 [1/1]0xc000048720 (0x496388,0xc000048730)
	println(results, s, strings, err)

	// results=0, s=r1, strings=[1], err=new error
	fmt.Printf("results=%d, s=%s, strings=%v, err=%v\n", results, s, strings, err)

	// results r1: 8, r2: r2
	r1, r2 := returnMultiDefineResults()
	fmt.Printf("results r1: %d, r2: %s", r1, r2)
}
