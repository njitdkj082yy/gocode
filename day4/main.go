package main

import "fmt"

func main() {
	var flag1 bool = true
	var flag2 bool = 3 > 9 //false

	fmt.Println("flage1:", flag1)
	fmt.Println("flage2:", flag2)

	var s1 string = "Hello 乔四"
	fmt.Println(s1)
	s1 = "乔四运维！！！Go!!!"
	fmt.Println(s1)
	//s1[1] = '1' //  不可以改值
	var s2 string = `hello "" \t go!` //反引号赋值，反引号在键盘1，2，3，4，5的左边那个
	fmt.Println(s2)
	var s3 string = s1 + //换行写的时候，要将“+”号留在本行
		s2
	fmt.Println(s3)
}
