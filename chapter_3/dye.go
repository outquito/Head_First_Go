package main

import (
	"fmt"
	"log"
)

func paintNeeded(width, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	liters := area / 10.0
	fmt.Printf("%.2f liters needed\n", liters)
	return liters, nil
}

func main() {
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed\n", amount)
	/*
		fmt.Println(err)
		fmt.Printf("%0.2f liters needed\n", amount)
		/*
		   total := 0.0
		   total += paintNeeded(4.3, 3.0)
		   total += paintNeeded(5.2, 3.5)
		   total += paintNeeded(5.0, 3.3)

		   fmt.Println("Total liters:", total)
	*/
}
