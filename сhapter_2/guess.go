package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("Я вибрав випадкове число від 1 до 100. \nЧи можете ви вгадати це число?")

	reader := bufio.NewReader(os.Stdin)
	success := false

	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("У вас залишилося", 10-guesses, "спроб.")

		fmt.Print("Введіть своє число: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal("Помилка!!! (", err, ") Повторіть запуск програми.")
		}

		if guess < target {
			fmt.Println("Oops. Ваше число меньше.")
		} else if guess > target {
			fmt.Println("Oops. Ваше число більше.")
		} else {
			success = true
			fmt.Println("Ви вгадали ЧИСЛО!!!")
			break
		}
	}
	if !success {
		fmt.Println("Нажаль ви не вгадали число:", target)
	}
}
