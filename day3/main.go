package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a1 int8 = 127
	var a2 uint8 = 127
	fmt.Println(a1+1, a2+1)
	var a3 int = 10000
	fmt.Printf("%T %T %T\n", a1, a2, a3)
	fmt.Println(unsafe.Sizeof(a1), unsafe.Sizeof(a2), unsafe.Sizeof(a3))
	var f1 float32 = 256.00001 //打印发现丢失精度
	var f2 float64 = 256.00001
	fmt.Println(f1, f2)
	var c1 byte = 'a'
	fmt.Println(c1)
	fmt.Printf("%c\n", c1)  //格式化输出字符
	fmt.Println("aaa\nbbb") //有了转义字符\n，则bbb会换行输出
	fmt.Println("aaa\rbb")  //\r表示回到本行开头 ，所以本来是输出aaa,再回头开头继续输出bb，则输出bba
}
