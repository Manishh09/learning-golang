package main

import "fmt"

func main() {
	fmt.Println("Interfaces - Demo")

	// machine system

	cmachine := colormachine{
		color: "Yellow",
	}

	bwmachine := bwmachine{
		color: "Violet",
	}

	print(cmachine)

	print(bwmachine)

	// EMPTY interface

	acceptAnyType("Hello!")

	acceptAnyType(true)

	acceptAnyType(bwmachine)

	// notification system

	sms := sms{
		message: "Hi from SMS system",
	}

	email := email{
		message: "Hello from EMAIL system",
	}

	printNotification(sms)

	printNotification(email)

	// geometry system

	rect := rectangle{
		radius: 2.5,
	}

	measure(rect)

}

func acceptAnyType(inp interface{}) {

	fmt.Println((inp))
}

// OBSERVATIONS:

// interfaces are a contract to manage the types
// every type will satisfy the EMPTY interface
// interfaces are implicit
