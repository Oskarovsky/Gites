package main

import "fmt"

func main() {
	fmt.Println("Oskar init git")
	var x = []int{10, 3: 20, 30}

	fmt.Printf("Test slices: %v\n", x)

	var turbo []int
	turbo = append(turbo, 2)
	turbo = append(turbo, 10)
	fmt.Printf("Turbo: %v", turbo)

	fmt.Printf("\n=====\n")
	fmt.Println(turbo, len(turbo), cap(turbo))
	turbo = append(turbo, 10)

	fmt.Println(turbo, len(turbo), cap(turbo))
	turbo = append(turbo, 20)

	fmt.Println(turbo, len(turbo), cap(turbo))
	turbo = append(turbo, 40)

	fmt.Println(turbo, len(turbo), cap(turbo))
	turbo = append(turbo, 50)

	fmt.Println(turbo, len(turbo), cap(turbo))
	turbo = append(turbo, 60)

	fmt.Printf("\n")
}
