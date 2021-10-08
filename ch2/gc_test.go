package ch2

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGc(t *testing.T) {
	var mem runtime.MemStats
	PrintStats(mem)
	//分配大块内存,打印内存和GC情况
	for i := 0; i < 10; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed!")
		}
	}
	PrintStats(mem)
	//分配大块内存,之前的大块内存绝大部分会被回收
	for i := 0; i < 10; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed!")
			time.Sleep(5 * time.Second)
		}
	}
	PrintStats(mem)
}
