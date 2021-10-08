package ch1

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func returnError(a, b int) error {
	if a == b {
		//生成自定义错误
		err := errors.New("Error in returnError() function!")
		return err
	} else {
		return nil
	}
}

func TestErrorCreate(t *testing.T) {
	//Error方法返回错误的信息
	assert.Equal(t, "Error in returnError() function!", returnError(1, 1).Error())
	//go中最常见的错误处理代码
	e := returnError(1, 1)
	if e != nil {
		t.Log(e)
		//如果这个错误很严重,需要停止程序运行,则使用panic方法
		//panic(e)
	} else {
		//...
	}
}

func TestErrorHandle(t *testing.T) {
	if len(os.Args) == 1 {
		fmt.Println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	var min, max = 0.0, 0.0
	floatArgExist := false
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err == nil {
			if !floatArgExist {
				min = n
				max = n
				floatArgExist = true
			}
			if n < min {
				min = n
			}
			if n > max {
				max = n
			}
		}
	}
	if !floatArgExist {
		fmt.Println("None of the arguments is a float!")
		return
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
