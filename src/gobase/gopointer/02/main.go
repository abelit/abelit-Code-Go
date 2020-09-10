package main

import (
	"fmt"
)

// CoffeeMachine is a struct
type CoffeeMachine struct {
	// NumberOfCoffeeBeans is variable
	NumberOfCoffeeBeans int
}

// NewCoffeeMachine is a function
func NewCoffeeMachine() CoffeeMachine {
	return CoffeeMachine{}
}

// SetNumberOfCoffeeBeans is a function
func (cm CoffeeMachine) SetNumberOfCoffeeBeans(n int) CoffeeMachine {
	cm.NumberOfCoffeeBeans = n
	return cm
}

func main() {
	// Without pointer
	cm := NewCoffeeMachine()
	cm = cm.SetNumberOfCoffeeBeans(300)

	fmt.Printf("The coffee machine has %d beans\n", cm.NumberOfCoffeeBeans)
}
