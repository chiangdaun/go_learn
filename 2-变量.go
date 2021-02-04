
/*
Go语言的变量命名方式和C语言一样，变量名由字母数字下划线组成，首字母不能为数字。
声明变量的一般形式是使用var关键字

var identifier type
可以一次声明多个变量：
var identifier1, identifier2 type

*/
//package main
//import "fmt"
//
//func main() {
//	var a string = "Runoob"
//	fmt.Println(a)
//
//	var b, c int = 1, 2
//	fmt.Println(b, c)
//}
/*
变量声明：
1. 指定变量类型，如果没有初始化，则变量默认为零值
零值就是变量没有做初始化时系统默认设置的值
*/
//package main
//import "fmt"
//func main() {
//
//	// 声明一个变量并初始化
//	var a = "RUNOOB"
//	fmt.Println(a)
//
//	// 没有初始化就为零值
//	var b int
//	fmt.Println(b)
//
//	// bool 零值为 false
//	var c bool
//	fmt.Println(c)
//}

//package main
//
//import "fmt"
//
//func main() {
//	var i int
//	var f float64
//	var b bool
//	var s string
//	fmt.Printf("%v %v %v %q\n", i, f, b, s)
//}

/*
变量声明：
2. 根据值自行判定变量类型
*/
//package main
//
//import "fmt"
//
//func main() {
//	var d = true
//	fmt.Println(d)
//}

/*
变量声明：
3. 省略var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：
	var intVal int
	intVal :=1 // 这时候会产生编译错误
	intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
*/

package main
import "fmt"
func main() {
	f := "Runoob" // var f string = "Runoob"

	fmt.Println(f)
}