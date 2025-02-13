## Go 中 Error

### 错误理念

- GO 中的多值方法返回
- Error 是一种值，而不是控制流
- 错误应尽量使用 error 而不是 panic

### error 类型接口

```go
type error interface {
Error() string
}
```

### 自定义 error

```go
type PathError struct {
Op   string // 操作类型
Path string // 文件路径
Err  error  // 底层错误
}
```

这种设计展示了 Go 错误处理的一个重要理念：

- 错误应该包含上下文信息
- 错误可以包装其他错误
- 错误信息应该对调试有帮助

### 最佳实践

- 错误处理最佳实践

```go
file, err := os.Create(filename)
if err != nil {
// 处理错误
return err
}
```

- 类型断言进行错误检查

```go
if e, ok := err.(*os.PathError); ok {
// 处理特定类型的错误
}
```

## Panic 函数

```go
func panic(v any)
```

### 使用 Panic 的场景

- 程序初始化阶段（init 函数中比如配置错误，必要组件启动失败）
- 真正的致命错误（如耗尽内存）
- 开发阶段的断言（确保不可能发生的情况）
-

### 不应该使用 panic 的场景

- 业务校验异常
- API 调用失败
- 数据库操作错误

## recover 函数

```go
func recover() any
```

### recover 函数的作用

- 捕获当前 goroutines 中 panic(arguments) 函数中的 arguments,类似于 Java 中的 try-catch 机制 
- 只能在 defer 函数中生效，这是与 Java 不同的重要特点 
- 可以获取 panic 的错误信息，并使程序恢复正常执行

### recover 调用规则

- 只能在 defer 后的匿名函数中调用

```go
// 正确示例
func main()  {
    defer func() {
		if r := revcover() ; r != nil {
            fmt.Printf("catch the panic : %v", r)
        }   
    }
}
```
- 只对当前 goroutine 有效
```go
func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("不会捕获到其他 goroutine 的 panic")
        }
    }()

    // 此协程有 panic，会导致整个程序崩溃，并不会走到上面的 recover() 
    go func() {
        panic("另一个 goroutine 的 panic")  // 主 goroutine 无法捕获
    }()

    time.Sleep(time.Second)
}
```










