package main

import "fmt"

func main() {
	fmt.Println("VARIABLES \t TYPES")

	var fName string = "Go" // explicit  type inference

	lName := "GoPher" // implicit type inference

	var age int = 20
	profession := "Go Developer"
	fmt.Println((profession[0]))

	isAuthenticatedUser := false

	heighOfPerson := 70.255

	fmt.Printf("types of fName variable is %T \n", fName)

	fmt.Println(lName)

	fmt.Printf("types of age variable is %T \n", age)

	fmt.Printf("types of isAuthenticatedUser variable is %T \n", isAuthenticatedUser)
	fmt.Printf("types of heighOfPerson variable is %T and height is %f \n", heighOfPerson, heighOfPerson)

	fmt.Printf("types of profession variable is %T \n", profession)

	var myVar1, myVar2 int = 21, 22
	fmt.Print(myVar1, myVar2)

	// printZeroValues()
}
