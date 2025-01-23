## new 关键字

- 只进行内存分配，并且将内存置为 0
- 不进行初始化

## make

## Array 初始化

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