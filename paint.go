package main

import (
	"fmt"
)

func paint() {
	var total, amount float64
	var err error
	amount, err = paintNeeded(5.7, -3.5)
	if err != nil {
		fmt.Println(err)
	} else {
		total += amount
	}
	amount, err = paintNeeded(6.2, 2.7)
	if err != nil {
		fmt.Println(err)
	} else {
		total += amount
	}
	fmt.Printf("Total: %.2f\n", total)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width <= 0 {
		return 0, fmt.Errorf("width of %.2f is invalid", width)
	}
	if height <= 0 {
		return 0, fmt.Errorf("height of %.2f is invalid", height)
	}
	area := width * height
	result := area / 10
	fmt.Printf("You need %.2f liters of paint\n", result)
	return result, nil
}
