### Background

本文是以为大纲 [effective go](https://go.dev/doc/effective_go)

### Init go mod

```bash
go mod init effective-go
```

go version - 1.23.4

### go playground

[go playground](https://go.dev/play/)

### 类型

在 package builtin 中即 builtin/builtin.go 文件, 会有一些自定义类型，比如 rune

#### rune 类型

```go
type rune = int32
```

rune 类型是 Go 语言的一种特殊数字类型。 和 int32 类型等价， Go 语言通过 rune 处理中文，支持国际化多语言

### 指针和类型
- *Type 在类型声明中表示"指向 Type 的指针"
- *variable 在表达式中才表示解引用操作

### Go vet 分析
常见的检查项
- Printf 类函数的格式字符串 
- 结构体标签使用是否正确 
- 死代码检测 
- 锁（Mutex）的复制问题 
- range 循环变量的使用 
- 未使用的错误返回值 
- 不当的错误处理模式