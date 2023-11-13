package main

import "fmt"

func main() {
	res := increment(2)
	fmt.Println(res)
	sum(1, 2, 3)

	arr := []int{1, 2, 3, 4}
	sum(arr...)
}

// FUNCTIONS
// PASS BY VALUE
func increment(x int) int {
	x++
	return x
}

// VARIADIC FUNCTIONS

// takes n number of int inputs
// E.g: fmt.Println

func sum(inp ...int) int { // inp parameter in this function signature is equivalent to []int
	sum := 0
	for _, num := range inp {
		sum += num
	}
	fmt.Println(sum)
	return sum
}
