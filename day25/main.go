package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var symbolMap = map[rune]int{
	'2': 2,
	'1': 1,
	'0': 0,
	'-': -1,
	'=': -2,
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := 0

	for scanner.Scan() {
		text := scanner.Text()
		value := decToSnafu(text)
		sum += value
	}
	file.Close()

	fmt.Println("Sum of the fuel requirements is", snafuToDec(sum))
}

func decToSnafu(text string) int {
	textLen := len(text)
	value := 0
	for i, char := range text {
		power := int(math.Pow(5, float64(textLen-i-1)))
		value += power * symbolMap[char]
	}
	return value
}

func snafuToDec(input int) string {
	value := ""
	for input != 0 {
		value = string("=-012"[(input+2)%5]) + value
		input = (input + 2) / 5
	}
	return value
}
