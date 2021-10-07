package main

import (
	"fmt"
)

var s1 string = "string1"
var z1 int = 10
var f1 float64 = 10.5

var (
	s2 string
	z2 int
	f2 float64
)

const c1 string = "const1"
const (
	c2 = iota
	c3 = 100
	c4
	c5 = iota
)

func main() {
	fmt.Println("hello go->day2")
	fmt.Println(s1, z1, f1)
	fmt.Println(s2, z2, f2)
	fmt.Println(c1)
	fmt.Println(c2, c3, c4, c5)
	a1 := 1
	a2 := 2
	a3 := 10
	const c2 int = 100
	fmt.Println(a1, a2, a3, c2)
}
