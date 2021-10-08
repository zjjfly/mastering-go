package ch2

import (
	"fmt"
	"testing"
	"unsafe"
)

//unsafe包的工具可以让你绕过Go的类型安全和内存安全机制,所以除非必要,不要使用!

func TestPointer(t *testing.T) {
	var value int64 = 1
	var p1 = &value
	//使用Pointer方法可以将任意的Go指针转换成Pointer对象,Pointer对象又可以转成任意的Go指针
	//这样就实现了不同类型的指针的转换
	var p2 = (*int32)(unsafe.Pointer(p1))
	fmt.Println("*p1: ", *p1) //1
	fmt.Println("*p2: ", *p2) //1
	//使用*号可以访问或修改指针指向的内容
	*p1 = 5434123412312431212
	fmt.Println("value:", value) //5434123412312431212
	fmt.Println("*p2: ", *p2)    //-930866580,因为32位整数的指针无法放下64位的整数
	*p2 = 3
	fmt.Println("value:", value) //3
	fmt.Println("*p1: ", *p1)    //3
}

//使用unsafe可以遍历数组
func TestArrayAccess(t *testing.T) {
	array := [...]int{0, 1, -2, 3, 4}
	pointer := &array[0]
	fmt.Print(*pointer, " ")
	step := unsafe.Sizeof(array[0])
	memoryAddress := uintptr(unsafe.Pointer(pointer)) + step
	for i := 0; i < len(array)-1; i++ {
		pointer = (*int)(unsafe.Pointer(memoryAddress))
		fmt.Print(*pointer, " ")
		memoryAddress = uintptr(unsafe.Pointer(pointer)) + step
	}
	fmt.Println()
	//获取数组最后一个元素之后的一个int的指针,go编译器不会报错,所以使用unsafe是很危险的
	pointer = (*int)(unsafe.Pointer(memoryAddress))
	fmt.Print("One more: ", *pointer, " ")
	memoryAddress = uintptr(unsafe.Pointer(pointer)) +
		unsafe.Sizeof(array[0])
	fmt.Println()
}
