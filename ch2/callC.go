package ch2

//调用独立的c代码,这种情况适用于调用复杂的C代码
//下面的代码前两行告诉了go程序在哪里找callC.h和callC.a

//#cgo CFLAGS: -I${SRCDIR}/clib
//#cgo LDFLAGS: ${SRCDIR}/clib/callC.a
//#include <stdlib.h>
//#include <callC.h>
import "C"
import (
	"fmt"
	"unsafe"
)

func callSeparateCCode() {
	fmt.Println("Going to call a C function!")
	C.cHello()
	fmt.Println("Going to call another C function!")
	//获取一个C的string对象
	myMessage := C.CString("This is Mihalis!")
	//是否这个字符串所占用的内存
	defer C.free(unsafe.Pointer(myMessage))
	C.printMessage(myMessage)
	fmt.Println("All perfectly done!")
}
