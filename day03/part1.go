package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

	for scanner.Scan() {
		text := scanner.Text()
		length := len(text)
		pivot := length / 2
		left := text[:pivot]
		right := text[pivot:length]

		var uniqueChar rune

		for _, char := range left {
			if strings.ContainsRune(right, char) {
				uniqueChar = char
				break
			}
		}

		if uniqueChar == 0 {
			for _, char := range left {
				if strings.ContainsRune(right, char) {
					uniqueChar = char
					break
				}
			}
		}

		priority := func() int {
			if unicode.IsLower(uniqueChar) {
				return int(uniqueChar) - 96
			} else {
				return int(uniqueChar) - 38
			}
		}()

		result += priority
	}

	fmt.Println("Sum of the priorities is", result)
	file.Close()
}
