package data

type File struct {
	name   string
	nepipe int8
}

// Composite Literals 复合字面量
func newFileMultiMethods() {
	// new 并 初始化
	f := new(File) // 指针类型
	f.name = ""
	f.nepipe = 0

	// 等价于 new(File)
	f = &File{}
	f.name = ""
	f.nepipe = 0

	// 会报错 f 是指针类型
	//f = File{
	//	name:   "",
	//	nepipe: 0,
	//}

	// 复合字面量初始化1
	f = &File{
		name:   "",
		nepipe: 0,
	}

	// 复合字面量初始化2
	f = &File{"", 0}
}
