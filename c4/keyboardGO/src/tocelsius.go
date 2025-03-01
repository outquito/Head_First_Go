// tocelsius перетворює температуру в градусах по Фаренгейту
package main

import (
	"Head_First_Go/chapter_4/keyboardGO/src/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a temperature in Fahrenheit: ")
	fahrenheit , err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}