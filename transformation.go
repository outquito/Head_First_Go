package main

import (
	"fmt"
	"reflect"
)

var length float64 = 1.2
var width int = 2

func main() {
	//	fmt.Println("Area is", length*float64(width))
	//	fmt.Println("length > width", length > float64(width))
	length = float64(width)
	fmt.Println(length)
	fmt.Println(reflect.TypeOf(length))
}
