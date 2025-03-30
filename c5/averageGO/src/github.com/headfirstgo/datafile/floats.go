package datafile //для читання даних з файлу

import (
	"bufio"
	"os"
	"strconv"
)

//читає значення float64 з кожного рядка файлу
func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64 //оголошення повертаного масиву
	file, err := os.Open(fileName) //відкриваємо файл
	if err != nil {
		return numbers, err
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		numbers = append(numbers, number) //перехід до наступного індексу масиву, нове значення до сегменту
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