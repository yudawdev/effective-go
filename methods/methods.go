package methods

import "fmt"

type ArrayListStr []string

func (a ArrayListStr) add(s []string) []string {
	return append(a, s...)
}

// arr : [1 2 3]
// arr : [1 2 3 4 5 6]
func userCase1() {
	arr := ArrayListStr{"1", "2", "3"}
	o := []string{"4", "5", "6"}

	arr.add(o)
	fmt.Printf("arr : %v", arr)

	fmt.Println()

	arr = arr.add(o)
	fmt.Printf("arr : %v", arr)
}

// 想实现一个方法：实例添加完成后，不需要对原变量进行重新复制
func (a *ArrayListStr) add2(s []string) {
	value := *a                // 将切片描述符复制一份，赋值给变量 value
	str := append(value, s...) // 创建新的切片
	*a = str                   // 将 *a 的描述更改为 str 切片的描述符
}

// 这么写不起作用
func (a *ArrayListStr) add3(s []string) {
	value := *a
	str := append(value, s...)
	a = &str // a 是一个局部变量，其实是外部 a 指针的复制品, 并不会改变外部 a 指针的值
}

// *a = append(*a, s...)
func userCase2() {
	arr := ArrayListStr{"1", "2", "3"}
	o := []string{"4", "5", "6"}

	arr.add2(o)
	fmt.Printf("arr : %v", arr)
}
