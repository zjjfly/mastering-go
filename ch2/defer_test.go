package ch2

//defer会在函数返回之前执行,有多个defer的时候,执行顺序是LIFO(先进后出)
//它的作用是可以在打开一个资源的时候马上说明需要在函数结束之后关闭它,这就可以避免忘记释放资源

import (
	"fmt"
	"log"
	"os"
	"testing"
)

var LOG_FILE = "./mGo.log"

func TestDefer1(t *testing.T) {
	for i := 3; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
	//1 2 3
}

func TestDefer2(t *testing.T) {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	//0 0 0
}

//下面这种使用defer的做法是最受推荐的
func TestDefer3(t *testing.T) {
	for i := 3; i > 0; i-- {
		defer func(i int) {
			fmt.Print(i, " ")
		}(i)
	}
	//1 2 3
}

//defer用于记录日志,这样可以增加日志可读性
func logDefer1(aLog *log.Logger) {
	aLog.Println("-- FUNCTION one ------")
	defer aLog.Println("-- FUNCTION one ------")
	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

//和上面的类似,只有打印的信息的格式有区别,看个人喜好
func logDefer2(aLog *log.Logger) {
	aLog.Println("------ FUNCTION one")
	defer aLog.Println("FUNCTION one ------")
	for i := 0; i < 10; i++ {
		aLog.Println(i)
	}
}

func TestLogDefer(t *testing.T) {
	file, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	logger := log.New(file, "logDefer", log.LstdFlags)
	logger.Println("Hello there!")
	logger.Println("Another log entry!")
	logDefer1(logger)
	logDefer2(logger)
}
