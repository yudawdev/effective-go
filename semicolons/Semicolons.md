link: https://go.dev/doc/effective_go#semicolons

1. 词法分析器在扫描时会使用一个简单的规则自动插入分号，所以源代码中基本上看不到分号。
2. 分号插入规则的一个重要结果是：你不能把控制结构（if、for、switch 或 select）的左花括号放在下一行。