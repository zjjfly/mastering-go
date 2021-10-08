package ch1

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"testing"
)

func TestStdOut(t *testing.T) {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}
	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")
}

func TestCommandLineArgs(t *testing.T) {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please give me one argument!")
		os.Exit(1)
	}
	min, _ := strconv.ParseFloat(arguments[1], 64)
	max, _ := strconv.ParseFloat(arguments[1], 64)
	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
