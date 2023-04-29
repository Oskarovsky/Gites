package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]

	fmt.Printf("x: %v\n", x)
	fmt.Printf("y: %v\n", y)
	fmt.Printf("z: %v\n", z)
	fmt.Printf("d: %v\n", d)
	fmt.Printf("e: %v\n", e)

	fmt.Printf("\n")
}
