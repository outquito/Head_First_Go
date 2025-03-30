package main

import "fmt"

func main() {
	numbers := make([]float64, 3)
	numbers[0] = 19.7
	numbers[2] = 25.2
	for i, numbers := range numbers {
		fmt.Println(i, numbers)
	}
	var letters = []string{"a", "b", "c"}
	for i, letters := range letters {
		fmt.Println(i, letters)
	}
}