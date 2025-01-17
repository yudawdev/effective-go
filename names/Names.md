link: https://go.dev/doc/effective_go#names

1. 可见可进行受首字母大写影响 - 首字母大写，可被其他包可见
2. 包名通常：全部小写、无下划线、无混合大小写
3. 包名是其源目录的基本名称。 例如：src/encoding/base64 中的包被导入为 "encoding/base64"，但其包名为 base64，而不是 encoding_base64 或 encodingBase64。
4. 建议不要使用 `import .
5. 建议命名时尽量简短，长名称并不会使其更易懂。
   - `once.Do；once.Do(setup)` vs  `once.DoOrWaitUntilDone(setup)`
6. Go 不自动提供 getters 和 setters 方法。如果你有一个不暴露的属性 `owner`, 建议提供 Getter 方法名为 `owner`, Setter 方法名为 `setOwner`
7. 对于接口的命名，如果接口中只有一个方法，那么接口建议的名称为 `methodName + er`, 例如：`Reader`
8. 如果你的类型实现了一个方法，其含义与某个广为人知的类型上的方法相同，请保持你定义的方法及签名和标准的保持一致。包括但不限于：Read、Write、Close、Flush、String 这些广为人知的接口及方法。
9. Go 中使用的是驼峰命名规则