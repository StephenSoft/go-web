package main

import (
	"fmt"
)

func foo(a string, b string) (name string) {
	name = a + b
	// fmt.Println("姓名为", name)
	return
}

func foo1(a, b int) (c, d int) {
	c = a * 2
	d = b * 3

	return
}

func main() {
	a := []int{1, 2, 3}
	for _, v := range a {

		fmt.Println(v)
	}
	name := foo("王", "小虎")

	fmt.Println("Name:", name)

	retn, retm := foo1(2, 4)

	fmt.Println("n: ", retn, "m:", retm)
}
