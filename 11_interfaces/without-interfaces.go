package main

type colormachine struct {
	color string
}

type bwmachine struct {
	color string
}

func (bw bwmachine) getColor() string {
	return bw.color
}

func (c colormachine) getColor() string {
	return c.color
}

// methods with same name are not allowed

// func printColor(bw bwmachine) {
// 	fmt.Println(bw.getColor())
// }

// func printColor(c colormachine) {
// 	fmt.Println(c.getColor())
// }

// machine that prints different colors

// printing is common

// we should have a common machine system
