package othernames

import (
	"effective-go/names"
	"fmt"
)

// 使用 names 包中的暴露的方法和变量
func test() {
	names.OutMethod()
	outer := names.Outer
	fmt.Println("outer: ", outer)
}
