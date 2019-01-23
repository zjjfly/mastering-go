package ch1

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestPrint(t *testing.T) {
	fmt.Println("a", "b", "c")
	//等价于
	fmt.Print("a", " ", "b", " ", "c", "\n")
	//S开头的函数返回给定格式的字符串
	assert.Equal(t, fmt.Sprint("abc 123"), fmt.Sprintf("%s %d", "abc", 123))
	//F开头的函数使用io.Writer把字符串写入文件
	file, e := os.OpenFile("test.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if e != nil {
		t.Error("open file error!")
	}
	defer file.Close()
	_, err := fmt.Fprintln(file, "test")
	if err != nil {
		t.Error(err)
	}
}
