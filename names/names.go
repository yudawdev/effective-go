package names

// 不会被其他包直接引用
var inner string

// Outer 可直接被其他包使用
var Outer string

// 首字母大写，可以被其他包调用
func OutMethod() {

}

// 首字母小写，对于其他包是不可见的
func innerMethod() {

}
