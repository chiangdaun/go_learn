package main

import "fmt"

func main()  {
	fmt.Println("Hello World!")
}
/*
第一行代码package main定义了包名。必须在源文件中非注释的第一行指明这个文件属于哪个包，如：package main。
package main 表示一个可独立执行的程序，每一个Go应用程序都必须包含一个名为main的包。

import "fmt"告诉Go编译器这个程序需要使用fmt包（函数或者其他因素）。fmt包实现了格式化IO的函数

func main()是程序开始执行的函数。main函数是每一个可执行程序所必须包含的，
	一般来说都是在启动后第一个执行的函数（如果有init函数会先执行init函数）。

fmt.Println()可以将字符串输出到控制台，并在最后增加换行符\n;
使用fmt.Print("....\n")可以达到相同的效果

Print 和 Println 这两个函数也支持使用变量，如：fmt.Println(arr)。如果没有特别指定，它们会以默认的打印格式将变量 arr 输出到控制台

需要注意的是 { 不能单独放在一行，所以以下代码在运行时会产生错误：
ackage main

import "fmt"

func main()
{  // 错误，{ 不能在单独的行上
    fmt.Println("Hello, World!")
}

*/