package main

import "fmt"

/**
 * This problem stumped me for more than a couple hours.
 * The root cause was that I thought I was modifying the
 * original slice when iterating over it.
 */

func main() {
	type Car struct {
		Make  string
		Model string
	}

	cars := []Car{
		Car{
			Make:  "Porsche",
			Model: "Taycan",
		},
		Car{
			Make:  "Porsche",
			Model: "718",
		},
		Car{
			Make:  "Porsche",
			Model: "Macan",
		},
	}

	// This will not work the way I'd expect it to because range creates a new copy of the slice and structs
	for index, car := range cars {
		car.Model = "Civic"
		// Let's print the memory addresses
		fmt.Printf("original array: %p range array: %p\n", &cars[index], &car)
	}

	// Prove that we haven't changed the original slice
	fmt.Println(cars)

	// How can we change the value of every element in the slice
	for index, _ := range cars {
		cars[index].Model = "Civic"
	}

	// Result is achieved
	fmt.Println(cars)
}
