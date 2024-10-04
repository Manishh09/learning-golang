package main

import (
	"fmt"
)

func conditionals() {
	res, rem, err := Div(10, 0)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else if rem == 0 {
		fmt.Printf("result of div is %v", res)
	} else {
		fmt.Printf("res is %v, rem is %v", res, rem)
	}

	// Switch - 1 ; with case expression
	switch {
	case err != nil:
		fmt.Println("error:", err.Error())
	case rem == 0:
		fmt.Printf("resul of div is %v", res)
	default:
		fmt.Printf("res is %v, rem is %v", res, rem)
	}

	// Switch -2;  with expression
	switch rem {
	case 0:
		fmt.Printf("exact division")
	case 1, 2:
		fmt.Printf("div operation was closer")
	default:
		fmt.Printf("was not closer")
	}

	// switch - 3 with an initializer

	switch a := div(0); a {

	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	}
}
