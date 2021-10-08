package ch2

//在go源码中嵌入c代码,方式是放在包声明之后的注释中,并导入C这个包
//这个包实际是一个虚拟的包,只是为了告诉go build命令在编译这个go文件之前使用cgo对它进行预处理
//这种方式适用于调用少量的简单的C代码

//#include <stdio.h>
//void callC() {
//	printf("Calling C code!\n");
//}
import "C"
import "fmt"

func callEmbedC() {
	fmt.Println("A Go statement!")
	C.callC()
	fmt.Println("Another Go statement!")
}


