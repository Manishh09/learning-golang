package main

import (
	"fmt"
)

type Tech struct {
	Name     string
	isLatest bool
}

// NESTED STRUCTS
type user struct {
	Name    string
	Email   string
	Age     int
	Address address
}

type address struct {
	State   string
	Country string
	Zip     int
}

// EMBEDDED STRUCTS
type users struct {
	address
	Name  string
	Email string
	Age   int
}

// STRUCT METHODS
// will be used in receiver functions
type rect struct {
	width  int
	height int
}

// area has a receiver of (r rect) : receiver here is a special type of function parameter
func (r rect) area() int {
	return r.width * r.height
}

func main() {
	tech := Tech{
		Name:     "Backend Development",
		isLatest: true,
	}
	fmt.Printf("\n***STRUCTS EXAMPLE***\n")
	fmt.Printf("Understanding Structs in %s is very important for learning %s \n", tech.Name, tech.isLatest)

	fmt.Printf("\n***NESTED STRUCTS***\n\n")
	// Nested struct - assigning values to struct fields and accessing them
	tempUser := user{
		Name:  "manish",
		Email: "manish.abc@nop.com",
		Age:   27,
		Address: address{
			State:   "TS",
			Country: "INDIA",
			Zip:     504251,
		},
	}

	fmt.Printf("User name is %s and email id is %s", tempUser.Name, tempUser.Email)
	fmt.Printf("\nAddress\nState is: %s \nCountry is: %s ", tempUser.Address.State, tempUser.Address.Country)
	fmt.Printf("\n")
	// fmt.Printf("\n***ANONYMOUS STRUCTS***\n\n")

	// tempUsers := struct {
	// 	Name  string
	// 	Email string
	// }{
	// 	Name:  "manish",
	// 	Email: "manish.abc@def.com",
	// }
	// fmt.Printf("User name is %s and email id is %s", tempUsers.Name, tempUsers.Email)
	// fmt.Printf("\n\n***EMBEDDED STRUCTS***\n\n")

	// embeddedStruct := users{
	// 	address: address{
	// 		State:   "TS",
	// 		Country: "IND",
	// 		Zip:     504251,
	// 	},
	// 	Name:  "manish",
	// 	Email: "manish.abc@def.com",
	// }
	// fmt.Printf("User name is %s and email id is %s", embeddedStruct.Name, embeddedStruct.Email)
	// fmt.Printf("\nAddress\nState is: %s \nCountry is: %s ", embeddedStruct.address.State, embeddedStruct.Country)
	// fmt.Printf("\n\n***ACCESSING EMBEDDED STRUCTS AT TOP LEVEL OF OTHER STRUCT***\n")
	// fmt.Printf("\nAddress\nState is: %s \nCountry is: %s ", embeddedStruct.State, embeddedStruct.Country)

	// fmt.Printf("\n\n****STRUCT METHODS****\n")

	// rectangle := rect{
	// 	width:  5,
	// 	height: 10,
	// }

	// fmt.Printf("Area of Rectangle is: %v cm^2", rectangle.area())
}
