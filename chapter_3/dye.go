package main

import "fmt"

func paintNeeded(width, height float64) float64 {
	area := width * height
	liters := area / 10.0
	fmt.Printf("%.2f liters needed\n", liters)
	return liters
}

func main() {
	paintNeeded(4.3, 3.0)
	paintNeeded(5.2, 3.5)
	paintNeeded(5.0, 3.3)
}
