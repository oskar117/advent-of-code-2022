package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Monkey struct {
	name, equation string
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	monkeyValues := make(map[string]int)
	monkeyQueue := make([]Monkey, 0)

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ": ")
		valueNum, err := strconv.Atoi(text[1])
		if err != nil {
			monkeyQueue = append(monkeyQueue, Monkey{text[0], text[1]})
		} else {
			monkeyValues[text[0]] = valueNum
		}
	}
	file.Close()

	for len(monkeyQueue) != 0 {
		item := monkeyQueue[0]
		monkeyQueue = monkeyQueue[1:]
		equation := strings.Split(item.equation, " ")
		firstValue, firstOk := monkeyValues[equation[0]]
		secondValue, secondOk := monkeyValues[equation[2]]
		if firstOk && secondOk {
			switch equation[1] {
			case "+":
				monkeyValues[item.name] = firstValue + secondValue
			case "-":
				monkeyValues[item.name] = firstValue - secondValue
			case "*":
				monkeyValues[item.name] = firstValue * secondValue
			case "/":
				monkeyValues[item.name] = firstValue / secondValue
			}
		} else {
			monkeyQueue = append(monkeyQueue, item)
		}
	}

	fmt.Println("Root monkey yells", monkeyValues["root"])
}
