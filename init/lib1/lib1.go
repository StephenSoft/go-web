package lib1

import "fmt"

func lib1Test() {
	fmt.Println("Lib2 lib2Test()")
}

func init() {
	fmt.Println("lib1")
}
