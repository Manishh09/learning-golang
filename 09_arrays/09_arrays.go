package main

import "fmt"

func main() {
	arrayDS()
}

func arrayDS() {
	// declaration
	var intArr [4]int32
	tempArr := [5]int32{1, 2, 3, 4}

	fmt.Println(intArr) // default values
	fmt.Println(tempArr)
	fmt.Println(tempArr[1:2])
	fmt.Println(tempArr[:2])
	fmt.Println(tempArr[1:])
	fmt.Println(tempArr[:])
	fmt.Println(tempArr[:0])
	fmt.Println(tempArr[0:0])

	// len and cap of an array in go

	fmt.Println(len(tempArr))
	fmt.Println(cap(tempArr))

}

// arrays - are
// fixed length
// same time
// indexed
// contiguous memory
