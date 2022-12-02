package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
		a, b := func() (string, string) {
			text := strings.Fields(scanner.Text())
			return text[0], text[1]
		}()
		switch b {
		case "X":
			result += 1
			switch a {
			case "A":
				result += 3
			case "B":
				result += 0
			case "C":
				result += 6
			}
		case "Y":
			result += 2
			switch a {
			case "A":
				result += 6
			case "B":
				result += 3
			case "C":
				result += 0
			}
		case "Z":
			result += 3
			switch a {
			case "A":
				result += 0
			case "B":
				result += 6
			case "C":
				result += 3
			}
		}
	}
	fmt.Println("You'd get", result, "points.")
	file.Close()
}
