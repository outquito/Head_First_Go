package main

import "fmt"

func main() {
	a := "a"
	b := "b"
	if true {
		c := "c"
		if true {
			d := "d"
			fmt.Println("1.", a)
			fmt.Println("1.", b)
			fmt.Println("1.", c)
			fmt.Println("1.", d)
		}
		fmt.Println("2.", a)
		fmt.Println("2.", b)
		fmt.Println("2.", c)
	}
	fmt.Println("3.", a)
	fmt.Println("3.", b)
}
