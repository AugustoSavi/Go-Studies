package main

import (
	"fmt"
)

type gasEngine struct {
	mpg     uint8
	gallons uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type engine interface {
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("yes")
	} else {
		fmt.Println("to fuel")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{2, 10}

	canMakeIt(myEngine, 21)
}
