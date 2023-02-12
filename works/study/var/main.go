package main

import "fmt"

func main() {

	var i int = 0
	str := "string"

	fmt.Println("i", i)
	fmt.Println("str", str)

	x, _ := foo()
	_, y := foo()

	fmt.Println("x,y", x, y)
}

func foo() (int, string) {
	return 10, "string"
}
