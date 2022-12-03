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

	var lines []string
	result := 0

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for i := 0; i < len(lines); i += 3 {
		first_elf := lines[i]
		second_elf := lines[i+1]
		third_elf := lines[i+2]

		var uniqueChar rune

		for _, char := range first_elf {
			if strings.ContainsRune(second_elf, char) && strings.ContainsRune(third_elf, char) {
				uniqueChar = char
				break
			}
		}

		result += func() int {
			if unicode.IsLower(uniqueChar) {
				return int(uniqueChar) - 96
			} else {
				return int(uniqueChar) - 38
			}
		}()
	}

	fmt.Println("Sum of the priorities is", result)
	file.Close()
}
