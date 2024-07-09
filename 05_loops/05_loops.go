package main

import "fmt"

func main() {

	fmt.Println("for - loop")
	for _, num := range []int{1, 2, 3, 4} {
		fmt.Println(num)
	}

}
