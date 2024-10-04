package main

import "fmt"

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

type machine interface {
	getColor() string
}

func print(m machine) {
	fmt.Println(m.getColor())
}
