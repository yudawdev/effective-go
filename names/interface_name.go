package names

import "fmt"

// Sender 接口的命名建议以方法名称+ `-er`
type Sender interface {
	send()
}

// MyReader 如果你的类型实现了一个方法，其含义与某个广为人知的类型上的方法相同，请保持你定义的方法及签名和标准的保持一致。
// MyReader 正确的使用方式：实现标准的 Read 方法
type MyReader struct{}

func (r *MyReader) Read(p []byte) (n int, err error) {
	// 实现读取逻辑
	return len(p), nil
}

// BadReader 错误的使用方式：名字相同但签名不同
type BadReader struct{}

func (r *BadReader) Read() string { // 不要这样做！
	return "data"
}

// Person 正确的字符串转换方法命名
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string { // 好的命名，遵循标准接口
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ToString 不推荐的命名方式
func (p Person) ToString() string { // 不要这样做！
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}
