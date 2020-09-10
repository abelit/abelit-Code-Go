package main

import (
	"fmt"
	"time"
)

type CoffeeMachine struct {
	UID                 string
	Description         string
	NumberOfCoffeeBeans int
}

func NewCoffeeMachine() CoffeeMachine {
	return CoffeeMachine{}
}

func (cm CoffeeMachine) SetUID(n string) CoffeeMachine {
	cm.UID = n
	return cm
}

func (cm CoffeeMachine) SetDescription(n string) CoffeeMachine {
	cm.Description = n
	return cm
}

func (cm CoffeeMachine) SetNumberOfCoffeeBeans(n int) CoffeeMachine {
	cm.NumberOfCoffeeBeans = n
	return cm
}

func main() {
	cm := NewCoffeeMachine()

	start := time.Now()

	for i := 0; i < 10000000; i++ {
		cm = cm.SetUID(fmt.Sprintf("random-generated-uid-%d", i))
		cm = cm.SetDescription(fmt.Sprintf("This is the best coffee machine that is around! This is version %d", i))
		cm = cm.SetNumberOfCoffeeBeans(i)
	}

	elapsed := time.Since(start)

	fmt.Printf("It took %s\n", elapsed)
}
