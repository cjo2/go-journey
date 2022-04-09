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
	porscheTaycan := NewCar("Porsche", "Taycan", "Atlanta")
	porsche718 := NewCar("Porsche", "718", "Atlanta")
	porscheMacan := NewCarStuttgart("Porsche", "Macan", time.Now)
	porsche991 := NewCarStuttgart("Porsche", "991", GetTime)

	fmt.Printf("func address: %p \n", porscheTaycan.ManufacturedTime)
	fmt.Printf("func address: %p \n", porsche718.ManufacturedTime)
	fmt.Printf("func address: %p \n", porscheMacan.ManufacturedTime)
	fmt.Printf("func address: %p \n", porsche991.ManufacturedTime)
	//func address: 0x107a880
	//func address: 0x107a880
	//func address: 0x107a880
	//func address: 0x108b9e0
}

type Car struct {
	ManufacturedTime     func() time.Time
	ManufacturedLocation string
	Make                 string
	Model                string
}

func NewCar(carMake string, model string, manufacturedLocation string) *Car {
	return &Car{
		ManufacturedTime:     time.Now,
		ManufacturedLocation: manufacturedLocation,
		Make:                 carMake,
		Model:                model,
	}
}

func NewCarStuttgart(carMake string, model string, manufactured func() time.Time) *Car {
	return &Car{
		ManufacturedTime:     manufactured,
		ManufacturedLocation: "Stuttgart",
		Make:                 carMake,
		Model:                model,
	}
}

func GetTime() time.Time {
	return time.Now()
}
