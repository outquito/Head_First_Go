package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	fmt.Println("Я вибрав випадкове число від 1 до 100.\nЧи можете ви вгадати це число?")

	reader := bufio.NewReader(os.Stdin)

	for guesses := 10; guesses > 0; guesses-- {
		fmt.Printf("У вас залишилося %d спроб.\n", guesses)
		fmt.Print("Введіть своє число: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Будь ласка, введіть коректне число.")
			continue
		}

		if guess < target {
			fmt.Println("Ваше число менше.")
		} else if guess > target {
			fmt.Println("Ваше число більше.")
		} else {
			fmt.Println("Вітаємо! Ви вгадали число!")
			return
		}
	}

	fmt.Printf("Нажаль, ви не вгадали число. Воно було: %d\n", target)
}
