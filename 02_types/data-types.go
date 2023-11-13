package main

import "fmt"

func printType() {
	var temp bool

	var name string

	var age int

	var height float32

	var runRate float64

	var complexVal complex64

	var complexBig complex128

	var byteVal byte

	var runVal rune

	fmt.Printf("%T", temp)
	fmt.Printf("%T", name)
	fmt.Printf("%T", age)
	fmt.Printf("%T", height)

	fmt.Printf("%T", runRate)
	fmt.Printf("%T", complexVal)
	fmt.Printf("%T", complexBig)

	fmt.Printf("%T", byteVal)

	fmt.Printf("%T", runVal)
}
