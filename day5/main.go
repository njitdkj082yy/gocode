package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a1 int32 = 10
	// var a2 int64 = a1  // 不可以不显示的转换，会报错：cannot use a1 (type int32) as type int64 in assignment
	var a2 int64 = int64(a1)
	fmt.Println(a1, a2)

	var a3 int32 = 11111111
	var a4 int8 = int8(a3) //溢出
	fmt.Println(a4)

	var n1 int = 20
	var f1 float64 = 20.0008
	var b1 bool = false
	var c1 byte = '1'
	fmt.Printf("%T\n", fmt.Sprintf("%d", n1))
	fmt.Printf("%T\n", fmt.Sprintf("%f", f1))
	fmt.Printf("%T\n", fmt.Sprintf("%t", b1)) //用%t表示布尔型
	fmt.Printf("%T\n", fmt.Sprintf("%c", c1))

	var s1 string = strconv.FormatInt(int64(a1), 10) //使用strconv.FormatInt对int64进行字符串转换，其它类似
	fmt.Printf("%T %s\n", s1, s1)

	var s2 string = "true"
	var b2 bool
	b2, _ = strconv.ParseBool(s2) //函书会返回两个值，只要第一个就好
	fmt.Printf("%T %t\n", b2, b2)
}
