// passFail.go буде повідомляти чи користувач сдав екзамен чи ні
package main

import (
	"Head_First_Go/chapter_4/keyboardGO/go/src/keyboard"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}

	var status string
	if grade >= 60 {
		status = "passing"
	} else {
		status = "failing"
	}

	fmt.Println("A grade of", grade, "is", status)
}
