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

func NewCoffeeMachine() *CoffeeMachine {
	return &CoffeeMachine{}
}

func (cm *CoffeeMachine) SetUID(n string) {
	cm.UID = n
}

func (cm *CoffeeMachine) SetDescription(n string) {
	cm.Description = n
}

func (cm *CoffeeMachine) SetNumberOfCoffeeBeans(n int) {
	cm.NumberOfCoffeeBeans = n
}

func main() {
	cm := NewCoffeeMachine()

	start := time.Now()

	for i := 0; i < 10000000; i++ {
		cm.SetUID(fmt.Sprintf("random-generated-uid-%d", i))
		cm.SetDescription(fmt.Sprintf("This is the best coffee machine that is around! This is version %d", i))
		cm.SetNumberOfCoffeeBeans(i)
	}

	elapsed := time.Since(start)

	fmt.Printf("It took %s\n", elapsed)
}
