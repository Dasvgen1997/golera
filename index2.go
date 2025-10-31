package main

import "fmt"

func main() {
	added := add(3, 8)
	fmt.Println(added, " added")

	subtred := subst(10, 8)
	fmt.Println(subtred, " substred")

	multipled := multip(3, 6)

	fmt.Println(multipled, " multipled")

	divised := divis(10, 3)

	fmt.Println(divised, " divised")
}

func add(first int, second int) int {
	return first + second
}

func subst(first int, second int) int {
	result := first - second
	return result
}

func multip(first int, second int) int {
	var result int
	result = first * second
	return result
}

func divis(first int, second int) float64 {
	return float64(first) / float64(second)
}
