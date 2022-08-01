package main

import "fmt"

func swap(a, b *int) {
	var tmp int
	tmp = *a
	*a = *b
	*b = tmp
}

func unswap(a, b int) {
	var tmp int
	tmp = a
	a = b
	b = tmp
}

func main() {
	// var a int = 10
	// var b int = 20
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a:", a, "b:", b)
	a = 10
	b = 20
	unswap(a, b)
	fmt.Println("a:", a, "b:", b)
}
