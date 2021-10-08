package ch2

import (
	"fmt"
	"os"
	"testing"
)

//对panic进行恢复
func TestPanicRecover(t *testing.T) {
	fmt.Println("Inside TestPanicRecover()")
	//从panic恢复的核心是一个defer的匿名函数,里面调用recover方法
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside TestPanicRecover()!")
		}
	}()
	fmt.Println("About to call b()")
	b()
	//下面这两行不会执行
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()!")
	fmt.Println("Exiting b()")
}

func TestPanic(t *testing.T) {
	if len(os.Args) == 1 {
		//panic会终止程序,所以一般会和recover搭配使用
		panic("Not enough arguments!")
	}
	fmt.Println("Thanks for the argument(s)!")
}

