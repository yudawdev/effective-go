## new 关键字

- 只进行内存分配，并且将内存置为 0
- 不进行初始化

## make

## Array

 ```go
p := [10]int{}
```

## slice

### slice 路径

runtime/slice.go

### slice 初始化

```go
s := make([]string, 10)
```

### runtime/slice.go vs  go/types/slice.go

runtime/slice.go: 实现 slice 的运行时操作(扩容、拷贝等)
go/types/slice.go: 实现编译期的类型检查和类型推导

### slice 遍历去除元素

- 建议使用：直接申请的切片，可读性高，但是使用了额外的资源
- 正向遍历，移动数组元素，修复
- 逆向遍历，移动数组元素

### 常用场景

- 判断某个对象否存在？ 使用 `slices` 包
- 简单排序： 使用 sort 中的方法, eg: `sort.Ints(ints)`
- 按照某个属性排序
- 去重
- groupBy

## Map

### Map 初始化

```go
// map 初始化
m := make(map[string]string)

// map 字面量初始化
m := map[string]string{
"k1": "v1"
"k2": "v2"
}

```

### Map 删除元素

```go
m := make(map[string]string)
delete(m, "key1")

```

### Map 常用场景

- 判断某个 key，在不在 map 中

```go
    // value 对应的值， ok 代表是有存在
value, ok := m[key]

```

- 对 map 进行排序，需要先对 key slice 排序

## Print

### %v 打印

- %v 是一个通用打印占位符
- %+v 打印出属性名称
- %#v 打印 go 代码格式
- %T 打印值的类型

### struct 自定义打印

需要实现 String 方法

### 时间

- Time

## init 函数

1. 执行时机：

- 在包级变量初始化后执行
- 在导入的包都初始化完成后执行
- 在 main 函数之前执行

2. 特性：

- 每个文件可以有多个 init 函数
- 无参数无返回值
- 自动执行，不能手动调用

3. 常见用途：

- 初始化无法通过变量声明实现的状态
- 验证程序状态
- 修复程序配置

示例代码展示了 init 的典型应用：检查环境变量、设置默认值、处理命令行参数。

Java 开发者需注意：Go 的 init 机制不同于 Java 的静态初始化块，更适合于包级别的初始化工作。

## 常量及枚举

1. 关键字 `const`

- 定义单个 `const name = "Go"`
- 定义多个

```go 
const (
PF = "PF"
F = "F"
)
```

2. 枚举

- 可以和关键字 `iota` 联动
- 对于复杂的业务枚举：大致分为两个部分：定义结构体 & 声明枚举实例
