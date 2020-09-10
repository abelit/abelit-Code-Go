package main

import "fmt"

// CoffeeMachine is a struct
type CoffeeMachine struct {
	NumberOfCoffeeBeans int
}

// NewCoffeeMachine is a function
func NewCoffeeMachine() *CoffeeMachine {
	return &CoffeeMachine{}
}

// SetNumberOfCoffeeBeans is a function
func (cm *CoffeeMachine) SetNumberOfCoffeeBeans(n int) {
	cm.NumberOfCoffeeBeans = n
}

func main() {
	// With Pointer
	cm := NewCoffeeMachine()
	cm.SetNumberOfCoffeeBeans(200)

	fmt.Printf("The coffee machine has %d beans\n", cm.NumberOfCoffeeBeans)
}
