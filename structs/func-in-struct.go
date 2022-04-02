package main

import (
	"fmt"
	"time"
)

/**
 * What happens if you have a struct that has a field with the type function?
 * Will multiple instances of that function be generated in memory?
 */

func main() {
	porscheTaycan := NewCar("Porsche", "Taycan")
	porsche718 := NewCar("Porsche", "718")
	porscheMacan := NewCarStuttgart("Porsche", "Macan", time.Now)
	porsche991 := NewCarStuttgart("Porsche", "991", GetTime)

	fmt.Printf("func address: %p \n", porscheTaycan.Manufactured)
	fmt.Printf("func address: %p \n", porsche718.Manufactured)
	fmt.Printf("func address: %p \n", porscheMacan.Manufactured)
	fmt.Printf("func address: %p \n", porsche991.Manufactured)
	//func address: 0x107a880
	//func address: 0x107a880
	//func address: 0x107a880
	//func address: 0x108b9e0
}

type Car struct {
	Manufactured func() time.Time
	Make         string
	Model        string
}

func NewCar(carMake string, model string) *Car {
	return &Car{
		Manufactured: time.Now,
		Make:         carMake,
		Model:        model,
	}
}

func NewCarStuttgart(carMake string, model string, manufactured func() time.Time) *Car {
	return &Car{
		Manufactured: manufactured,
		Make:         carMake,
		Model:        model,
	}
}

func GetTime() time.Time {
	return time.Now()
}
