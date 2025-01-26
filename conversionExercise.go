package main

import "fmt"

var price int = 100
var taxRate float64 = 0.08
var tax float64 = float64(price) * taxRate
var total float64 = float64(price) + tax
var availableFunds int = 20

func main() {
	fmt.Println("Price is", price, "$.")
	fmt.Println("Tax is", tax, "$.")
	fmt.Println("Total cost is", total, "$.")
	fmt.Println(availableFunds, "$ available.")
	fmt.Println("within budget?", total <= float64(availableFunds))
}
