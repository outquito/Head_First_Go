// average вирахування середнього значення
package main

import (
	"fmt"
	"log"
	"netmonitor/Head_First_Go/c5/averageGO/src/github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}