package main

import "fmt"

type employee struct {
	name        string
	designation string
	address     address
}

type address struct {
	zipCode  int
	location string
}

func main() {

	// first approach
	// emp := employee{
	// 	"S",
	// 	"B",
	// }

	// using var
	// var emp2 employee
	// fmt.Println(emp2)

	// emp2.designation = "Go dev"
	// emp2.name = "Manish"

	emp3 := employee{
		name:        "A",
		designation: "Go dev",
		address: address{
			zipCode:  500089,
			location: "Hyd",
		},
	}
	temp := &emp3

	temp.save("Will")

	//emp3.showDetails()

	// slice

	tempslc := []string{"A", "N", "M"}

	res := updateSlc(tempslc)

	fmt.Println("slc", res)

}

func updateSlc(sl []string) []string {
	sl[1] = "Q"
	return sl
}

func (e employee) showDetails() {
	fmt.Println(e)
}

func (e *employee) save(des string) {

	(*e).name = des
}

/******* Before Save

	block0

	block1 --- employee  , name : "A"


	block2 ---

	block3 ---


******/

/******* At Save (emp3.save("Will")) - copy of original struct value will be copied into memory
	block0

	block1 --- employee  , name : "A"


	block2 --- employee, name : "A"

	block3 ---


******/

/******* After Save - copied value will get updated but not the original
	block0

	block1 --- employee  , name : "A"


	block2 --- employee, name : "Will"

	block3 ---


******/
