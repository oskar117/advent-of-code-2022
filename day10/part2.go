package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	X, cycle_counter := 1, 0

	screen := make([][]string, 6)

	for scanner.Scan() {
		output := strings.Fields(scanner.Text())
		instruction := output[0]
		switch instruction {
		case "noop":
			screen = perform(1, &cycle_counter, X, screen)
		case "addx":
			screen = perform(2, &cycle_counter, X, screen)
			value, _ := strconv.Atoi(output[1])
			X += value
		}
	}
	file.Close()

	for _, row := range screen {
		for _, char := range row {
			fmt.Print(char)
		}
		fmt.Print("\n")
	}
}

func perform(cycles int, counter *int, registry int, screen [][]string) [][]string {
	for i := 0; i < cycles; i++ {
		*counter++
		row := (*counter - 1) / 40
		pixel := func() string {
			row_len := len(screen[row])
			if row_len >= registry-1 && row_len <= registry+1 {
				return "#"
			} else {
				return "."
			}
		}()
		screen[row] = append(screen[row], pixel)
	}
	return screen
}
