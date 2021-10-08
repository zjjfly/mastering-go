package ch1

import (
	"bufio"
	"os"
	"testing"
)

func TestStdIn(t *testing.T) {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		t.Log(">",scanner.Text())
	}
}
