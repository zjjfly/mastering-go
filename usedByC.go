//必须是main包
package main

//C语言调用go代码

import "C"
import "fmt"

//函数之前以export开头的这一行注释告诉编译器把这个函数导出供C调用
//export PrintMessage
func PrintMessage() {
	fmt.Println("A Go function!")
}

//export Multiply
func Multiply(a, b int) int {
	return a * b
}

//main方法是必须的
func main()  {

}
