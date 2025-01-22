link: https://go.dev/doc/effective_go#control-structures

## If 说明

1. if 语句中，强制使用花括号
2. if 和 switch 可以接受初始化语句，通常设置一个局部变量。
   ```go
   if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
    }
   ```
3. 建议提前返回，不要深层嵌套

```go
f, err := os.Open(name)
if err != nil {
return err
}
d, err := f.Stat()
if err != nil {
f.Close()
return err
}
codeUsing(f, d)
```

## For 说明

1. Go 中 `for` 统一了所有循环形式，并没有 `do-while` 形式
2. 遍历使用 `range` 关键字, 对字符串的 range 遍历会自动处理 UTF-8 编码
3. `++` & `--` 是语句非表达式
4. 使用平行赋值替代逗号运算符
5. 跳出循环一般使用 `break`，如果内部有嵌套或者 `switch`，使用循环别名跳出

## Switch

1. switch 后的表达式：可以没有、不仅仅是常量或者整数。没有即为 true，直到符合 case 后的表达式
2. switch 可以代替 if else if else
3. switch 匹配类型
