package ch2

//存放指针的方式会影响到GC的效率,尤其是对大的集合

import (
	"runtime"
	"testing"
)

type data struct {
	i, j int
}

func BenchmarkSliceGC(b *testing.B) {
	b.N = 40000000
	var structure []data
	for i := 0; i < b.N; i++ {
		value := i
		structure = append(structure, data{value, value})
	}
	runtime.GC()
	//加下面这一行是为了防止GC过早的把structure回收,因为structure只在上面的循环中被引用了
	_ = structure[0]
}

func BenchmarkMapStarGC(b *testing.B) {
	b.N = 40000000
	myMap := make(map[int]*int)
	for i := 0; i < b.N; i++ {
		value := i
		myMap[value] = &value
	}
	runtime.GC()
	_ = myMap[0]
}

func BenchmarkMapNoStarGC(b *testing.B) {
	b.N = 40000000
	myMap := make(map[int]int)
	for i := 0; i < b.N; i++ {
		value := i
		myMap[value] = value
	}
	runtime.GC()
	_ = myMap[0]
}

//把上面的大map分拆为很多的小map,并使用hash算法放到不同的bucket中
func BenchmarkMapSplitGC(b *testing.B) {
	b.N = 40000000
	split := make([]map[int]int, 200)
	for i := range split {
		split[i] = make(map[int]int)
	}
	for i := 0; i < b.N; i++ {
		value := i
		split[value%200][value] = value
	}
	runtime.GC()
	_ = split[0][0]
}

//测试结果:
//BenchmarkSliceGC-8       	40000000	        51.0 ns/op
//BenchmarkMapStarGC-8     	40000000	       403 ns/op
//BenchmarkMapNoStarGC-8   	40000000	       278 ns/op
//BenchmarkMapSplitGC-8    	40000000	       261 ns/op
//显然,Slice的性能是最好的,而map的性能就不是很理想了,这和GC的停顿时间有关系,所以要避免在map中存太多的数据
