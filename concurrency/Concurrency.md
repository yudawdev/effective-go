## Go 并发

## Goroutines - 协程

main 函数确实是在一个特殊的 goroutine 中运行的，我们通常称它为"主 goroutine"。

## chan - 通道

### 方向分类

- 双向通道 chan
- 只读通道 <-chan - 只能从通道中读取
- 只写通道 chan<- - 只能写入到通道中

### 有无缓冲分类

- 无缓冲 `make(chan int)`
    - 无缓冲通道，通道必须同时有发送者和消费者，才能完成通信，即发送和接受是同步发生的
- 有缓冲 `make(chan int, 5)`
    - 只有缓冲区满时才阻塞

### 规则

- 双向通道 可以 转化为单向通道
- 单向通道 不可以 转化为双向通道

### 特性

- 线程安全(同步)是通道的核心特征
- 通道可关闭
  ```go 
  close(ch)
  if val, ok := <-ch; ok {
    // 通道开放
  } else {
    // 通道关闭
  }
  ```

## select

- 多路复用，用于处理多个通道操作
- 类似 switch，但是针对通道操作

**基本语法**

```go
select {
case <-ch1:
// 分支 1
case <-ch2:
// 分支 2
}
```

当任意通道（ch1、ch2）有返回后，执行先返回的通道 case

### 特殊情况一：非阻塞 select

```go
select {
case <-ch:
// 处理 channel 数据
default:
// 立即执行，不阻塞
}
```

### 特殊情况二：空的 select

```go
select {} // 永久阻塞
```

### 特殊情况三：多个 channel 同时就绪, 随机选择一个 case 执行

```go
select {
case <-ch1:
// 随机选择执行
case <-ch2:
// 随机选择执行
}
```

## 并发原语 - sync.WaitGroup

含义：等待一组 goroutines 完成

## 并发原语 - sync.Mutex

含义：互斥锁

特点：

- 不可重入

## 典型的并发模式

### 生产者-消费者模式

### 扇出（Fan-out）模式

### 超时控制模式

## 问题

1. 如何在一个子协程中，读取另外协程的变量


## 并发（Concurrency）与并行（Parallelism）的区别

[Concurrency is not parallelism](https://go.dev/blog/waza-talk)

## Go vs Java 的并发模型比较

### Go 并发模型 - 基于消息传递
CSP 模型 - Communicating Sequential Processes（通信顺序进程）
Go 的设计哲学 "Do not communicate by sharing memory; share memory by communicating" （不要通过共享内存来通信，而要通过通信来共享内存）正是 CSP 模型的核心思想的体现。

### Go 协程的调度
- 协程是通过 Go 运行时的调度器调度的
- 调度点
  - 系统调用 
  - channel 操作 
  - time.Sleep()
  - mutex 操作 
  - 主动调用 runtime.Gosched()
- Go 使用 M:N 调度模型:
  - M (Machine): 操作系统线程 
  - P (Processor): 调度上下文,默认等于 CPU 核心数 
  - G (Goroutine): 用户态的协程
  - 在单核场景下: 只有一个 M (系统线程) 和一个 P (调度上下文), 多个 G (goroutine) 会被调度到这个 P 上执行

### Java 并发模型 - 基于共享内存

