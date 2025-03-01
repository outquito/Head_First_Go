package datafile //для читання даних з файлу

import (
	"bufio"
	"os"
	"strconv"
)

//читає значення float64 з кожного рядка файлу
func GetFloats(fileName string) ([3]float64, error) {
	var numbers [3]float64 //оголошення повертаного масиву
	file, err := os.Open(fileName) //відкриваємо файл
	if err != nil {
		return numbers, err
	}
	i := 0 //змінна для збереження індексу, по якому має виконуватися зміни
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++ //перехід до наступного індексу масиву
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}
	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil //якщо до цього дійшло, значить помилок не було
}