package main

import "fmt"

func main() {
	var myInt int
	var myIntPointer *int
	myInt = 42
	myIntPointer = &myInt
	fmt.Println("1. myIntPointer", myIntPointer)
	fmt.Println("2. myIntPointer", *myIntPointer)
	fmt.Println("3. myIntPointer", &myIntPointer)
	fmt.Println("4. myInt", myInt)
	//fmt.Println("5. myInt", *myInt)
	fmt.Println("6. myInt", &myInt)

}