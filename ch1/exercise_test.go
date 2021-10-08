package ch1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestExercise3(t *testing.T) {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if _, err := strconv.ParseInt(text,10,64); err != nil {
			if text == "STOP" {
				return
			} else {
				fmt.Println("Please input a valid integer!")
			}
		}
	}

}