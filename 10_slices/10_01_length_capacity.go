package main

import "fmt"

func main() {
	evaluateLengthCapOfSlice()
}

func evaluateLengthCapOfSlice() {
	movies := []string{"avengers", "eagle"}

	fmt.Println("initial len", len(movies))
	fmt.Println("initial cap", cap(movies))

	movies = append(movies, "star-wars")

	movies = append(movies, "star-wars - 2")

	movies = append(movies, "star-wars - 3")

	fmt.Println("len - current", len(movies))
	fmt.Println("cap - current", cap(movies))

	// Note: capacity of a slice gets doubled if the current length of slice exceeds the previous slice's length

	// e.g: a slice current length = 3
	// 				capacity	   = 3

	// 	lets say we add one item to the current slice
	//  length  now is    4
	//  capacity will be  6

	// Now lets say we added 2 more elements
	// length will be    6
	// capacity will be  6

	// Note: capacity remains same till the length matches with current capacity

	// Now add one more item to the slice
	// length will be 7
	// capacity becomes 12
}
