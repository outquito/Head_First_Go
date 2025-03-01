package main // сканування файлу

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt") //відкриття файлу для читання
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file) //створення нового значення для Scanner для файлу
	for scanner.Scan() { //читання рядку з файлу
		fmt.Println(scanner.Text()) //вивід прочитаного рядка
	}
	err = file.Close() //ЗАКРИТТЯ ФАЙЛУ для звільнення ресурсів
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}