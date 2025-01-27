package main

import "fmt"

func main() {
	if true {
		fmt.Println("1. true")
	}
	if false {
		fmt.Println("2. false")
	}
	if !false {
		fmt.Println("3. !false")
	}
	if true {
		fmt.Println("4. if true")
	} else {
		fmt.Println("5. else")
	}
	if false {
		fmt.Println("6. if false")
	} else if true {
		fmt.Println("7. else if true")
	}
	if 12 == 12 {
		fmt.Println("8. 12 == 12")
	}
	if 12 != 12 {
		fmt.Println("9. 12 != 12")
	}
	if 12 > 12 {
		fmt.Println("10. 12 > 12")
	}
	if 12 >= 12 {
		fmt.Println("11. 12 >= 12")
	}
	if 12 == 12 && 5.9 == 5.9 {
		fmt.Println("12. 12 == 12 && 5.9 == 5.9")
	}
	if 12 == 12 && 5.9 == 6.4 {
		fmt.Println("13. 12 == 12 && 5.9 == 6.4")
	}
	if 12 == 12 || 5.9 == 6.4 {
		fmt.Println("14. 12 == 12 || 5.9 == 6.4")
	}
}
