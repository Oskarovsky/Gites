package main

import "fmt"

func main() {

	x := []int{1, 10, 100, 1000}
	y := x[:2]
	z := x[1:]

	x[1] = 11
	y[0] = 111
	z[1] = 0
	fmt.Println("x=", x)
	fmt.Println("y=", y)
	fmt.Println("z=", z)
	fmt.Println("z=", z)
	fmt.Println("z=", z)
	fmt.Println("z=1", z)
	fmt.Println("z=2", z)
	fmt.Println("z=3", z)
	fmt.Println("z=4", z)
}
