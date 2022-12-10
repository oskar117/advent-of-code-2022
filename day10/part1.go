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

	X, cycle_counter, signal_strength := 1, 0, 0

	for scanner.Scan() {
		output := strings.Fields(scanner.Text())
		instruction := output[0]
		switch instruction {
		case "noop":
			perform(1, &cycle_counter, &signal_strength, X)
		case "addx":
			perform(2, &cycle_counter, &signal_strength, X)
			value, _ := strconv.Atoi(output[1])
			X += value
		}
	}

	fmt.Println("Final signal strength is", signal_strength)
	file.Close()
}

func perform(cycles int, counter *int, signal_strength *int, registry int) {
	for i := 0; i < cycles; i++ {
		*counter++
		if (*counter - 20) % 40 == 0 && *counter <= 220 {
			*signal_strength += *counter * registry
		}
	}
}
	
