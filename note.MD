### basic.go 变量定义
**使用var关键字** 
- var a, b, c bool
- var s1, s2 string = "hello", "world"
- 可放在函数内或者直接放在包内
- 使用var()集中定义变量

**让编译器自动定义类型**
- var a, b, i, s1, s2 = true, false, 3, "hello", "world"

**使用:=定义变量**
- a, b, i, s1, s2 := true, false, 3, "hello", "world"

### 内建变量类型
- bool, string
- (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr(指针)
- byte(8位), rune(字符型，32位)
- float32, float64, complex64, complex128

### 强制类型转换
- 类型转换是强制的

### 常亮定义

### 枚举
- 普通枚举
- 自增枚举

## 回顾

- 变量类型写在变量名之后
- 编译器可以推断变量类型
- 没有char,只有rune
- 原生支持复数类型

### if条件判断
- if的条件里可以赋值
- if的条件里赋值的变量作用域就在这个if语句里

### switch循环
### for循环

- for if 后面的条件没有括号
- if 条件里面也可以定义变量
- 没有while判断

### 函数语法要点回顾
- 返回值类型写在最后面
- 可以返回多个值
- 函数作为参数
- 没有默认参数，可选参数