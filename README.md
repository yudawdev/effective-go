# Effective Go Learning Guide

这是一个面向 Java 开发者的 Go 语言学习指南，基于 [Effective Go](https://go.dev/doc/effective_go) 官方文档，结合实践经验整理而成。

## 项目说明

- Go 版本: 1.23.4
- 在线练习环境: [Go Playground](https://go.dev/play/)

## 目录结构

```
effective-go/
├── blankidentifier/    # 空白标识符使用
├── commentary/         # 注释规范
├── concurrency/        # 并发编程
├── controlstructures/  # 控制结构
├── data/              # 数据类型和结构
├── error/             # 错误处理
├── formatting/        # 代码格式化
├── functions/         # 函数
├── interfaces/        # 接口
├── methods/           # 方法
└── names/             # 命名规范
```

## 核心内容

### 1. 代码组织与格式化

- 使用 `gofmt` 进行代码格式化
- 包的组织和命名规范
- 代码注释规范
- 分号的使用规则

### 2. 数据结构与类型

- 基本数据类型
- 数组和切片
- Map 的使用
- 结构体
- 指针操作

### 3. 函数与方法

- 函数定义和多返回值
- defer 机制
- 方法定义
- 接口实现
- init 函数特性

### 4. 并发编程

- Goroutines
- Channels
- Select 机制
- 并发模式
  - 生产者-消费者模式
  - 扇出模式
  - 超时控制
- sync 包的使用
  - WaitGroup
  - Mutex

### 5. 错误处理

- error 接口
- 错误处理最佳实践
- 自定义错误类型
- 错误包装

### 6. Go vs Java 比较

- 并发模型对比
  - Go: CSP 模型 (通过通信来共享内存)
  - Java: 共享内存模型
- 协程调度机制
- 错误处理方式
- 接口实现方式

## 最佳实践

1. 命名规范
   - 可见性由首字母大小写控制
   - 包名使用小写
   - 接口命名推荐使用 method + er
   - 使用驼峰命名法

2. 错误处理
   - 错误是值，不是异常
   - 错误应包含上下文信息
   - 提前返回错误，避免嵌套

3. 并发处理
   - 使用通道进行goroutine间通信
   - 合理使用select实现多路复用
   - 正确使用互斥锁和WaitGroup

4. 代码组织
   - 合理使用包的组织
   - 遵循"少即是多"的设计原则
   - 保持接口的简单性

## 工具使用

- `go fmt`: 代码格式化
- `go vet`: 代码静态分析
- `go mod`: 依赖管理

## 参考资源

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Documentation](https://go.dev/doc/)
- [Go Playground](https://go.dev/play/)

## 注意事项

1. 本项目主要面向有 Java 开发经验的开发者
2. 重点关注 Go 与 Java 的差异
3. 建议配合官方文档阅读
4. 代码示例均可在 Go Playground 中运行
