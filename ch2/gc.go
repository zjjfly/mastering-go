package ch2

import (
	"fmt"
	"runtime"
)

func PrintStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("Memory Alloc:",mem.Alloc)
	fmt.Println("Memory HeapAlloc:",mem.HeapAlloc)
	fmt.Println("Memory TotalAlloc:",mem.TotalAlloc)
	fmt.Println("Memory NumGC:",mem.NumGC)
	fmt.Println("-----")
}