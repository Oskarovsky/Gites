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

	fmt.Println("a=", z)
	fmt.Println("b=", z)
	fmt.Println("c=", z)
	fmt.Println("d=", z)

	fmt.Println("001")
	fmt.Println("002")
	fmt.Println("003")
	fmt.Println("004")

	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
}
