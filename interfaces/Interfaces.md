## Interfaces

1. 接口定义对象的行为方式,只要实现了特定方法就满足接口要求. Duck Typing
2. Go 常用简单接口(1-2个方法),如 io.Writer
3. 一个类型可以实现多个接口,示例展示了 Sequence 类型同时实现:
   sort.Interface(用于排序)
   Stringer 接口(用于格式化输出)

示例代码展示了 Sequence 类型:

实现排序所需的三个方法:Len()、Less()、Swap()
实现 String() 方法用于打印
包含 Copy() 方法用于创建副本

## Conversions 类型转化

- 显式类型转化 Type(value)
  val := int64(123)

- 类型判断
```go
// switch 方式
switch v := value.(type) {
case int:
// v 是 int 类型
case string:
// v 是 string 类型
}

// 单类型断言
str, ok := value.(string)

```

## 通用性
1. Go 中的接口实现是隐式的
2. 如果一个类型，和某接口定义的方法 signature 是一样的，即实现了该接口。