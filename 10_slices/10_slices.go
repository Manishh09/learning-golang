package main

import "fmt"

func main() {
	slicesDS()
}

func slicesDS() {
	var intSlice []int32 = []int32{1, 2, 3}
	fmt.Println(intSlice)

	intSlice = append(intSlice, 4) // [1,2,3,4,*,*] - length is 4 and capacity is 6
	fmt.Println(intSlice)

	fmt.Println(cap(intSlice))

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	var intSlice2 []int32 = []int32{9, 10}
	intSlice2 = append(intSlice, intSlice2...)
	fmt.Println(intSlice2)
	fmt.Println(cap(intSlice2))

	slicesWithMake := make([]int32, 3)

	// slicesWithMake = append(slicesWithMake, []int32{1, 2}...)

	fmt.Println(slicesWithMake, len(slicesWithMake), cap(slicesWithMake))

	sliceTemp := make([]string, 1, 2)
	fmt.Println(sliceTemp)
}
