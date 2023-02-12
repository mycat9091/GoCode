package main

import "fmt"

const (
	pi = 3.1415
	e  = 2.7182
)

func main() {

	const (
		n1 = iota
		n2 = 20
		n3 = iota
		n4
	)
	const n5 = iota

	fmt.Println("n1,n2,n3,n4,n5\n", n1, n2, n3, n4, n5)

	fmt.Println("pi", pi)
}
