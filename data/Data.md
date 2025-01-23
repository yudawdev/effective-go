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






