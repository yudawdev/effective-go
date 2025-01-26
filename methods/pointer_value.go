package methods

import "fmt"

// 获取指针 : *Type | &实例
// s: Hello, s Type: string
// p: 0xc000106050, p Type: *string
func getPointer() {
	s := "Hello"
	p := &s

	fmt.Printf("s: %s, s Type: %T\n p: %v, p Type: %T", s, s, p, p)
}

// 获取值 : *引用
// v : [hello], v Type: []string
// v1 : [1 2 3       ], v1 Type: [10]string
func getValue() {
	// 定义引用类型
	var p *[]string

	// 赋值引用
	p = &[]string{"hello"}

	// 通过引用获取值
	v := *p
	fmt.Printf("v : %v, v Type: %T", v, v)

	fmt.Println("")

	// 这行打印是额外想观察下：slice 和 array 打印出来的类型有何不同
	v1 := [10]string{"1", "2", "3"}
	fmt.Printf("v1 : %v, v1 Type: %T", v1, v1)
}

type People struct {
	Name string
	Age  int
}

func (p People) String() string {
	return fmt.Sprintf("People( %v, %d)", p.Name, p.Age)
}

func test() {

	// 直接修改 - 通过值类型变量
	peopleValue := People{
		Name: "Value",
		Age:  10,
	}

	peopleValue.Name = "Value1"
	(&peopleValue).Age = 11

	fmt.Println("peopleValue: ", peopleValue)

	// 通过指针修改
	peoplePointer := &People{
		Name: "Pointer",
		Age:  20,
	}

	peoplePointer.Name = "Value2"
	(*peoplePointer).Age = 30

	fmt.Println("peoplePointer: ", peoplePointer)

	// 误区：通过解引用变量修改 !!!
	value := *peoplePointer // 需要注意： value 是一块新的内存空间，将 *peoplePointer 复制了一份
	value.Name = "VVV"
	fmt.Println("* Reference: ", value)
	fmt.Println("ori : ", peoplePointer)
}

func test2() {
	people := &People{
		Name: "Value",
		Age:  10,
	}

	v := *people

	fmt.Println(v == *people)
}
