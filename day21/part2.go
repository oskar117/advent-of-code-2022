package main

import (
	"bufio"
	"fmt"
	"math"
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

	monkeyValues := make(map[string]float64)
	monkeyQueue := make([]Monkey, 0)
	var rootEquation string

	for scanner.Scan() {
		text := strings.Split(scanner.Text(), ": ")
		valueNum, err := strconv.Atoi(text[1])
		if err != nil {
			if text[0] == "root" {
				rootEquation = text[1]
			} else {
				monkeyQueue = append(monkeyQueue, Monkey{text[0], text[1]})
			}
		} else {
			monkeyValues[text[0]] = float64(valueNum)
		}
	}
	file.Close()

	
	low, high := 0, math.MaxInt
	for {
		mid := (low + high) / 2
		tempMonkeyValues := copyMap(monkeyValues)
		tempMonkeyQueue := monkeyQueue
		tempMonkeyValues["humn"] = float64(mid)
		for len(tempMonkeyQueue) != 0 {
			item := tempMonkeyQueue[0]
			tempMonkeyQueue = tempMonkeyQueue[1:]
			equation := strings.Split(item.equation, " ")
			firstValue, firstOk := tempMonkeyValues[equation[0]]
			secondValue, secondOk := tempMonkeyValues[equation[2]]
			if firstOk && secondOk {
				switch equation[1] {
				case "+":
					tempMonkeyValues[item.name] = firstValue + secondValue
				case "-":
					tempMonkeyValues[item.name] = firstValue - secondValue
				case "*":
					tempMonkeyValues[item.name] = firstValue * secondValue
				case "/":
					tempMonkeyValues[item.name] = firstValue / secondValue
				}
			} else {
				tempMonkeyQueue = append(tempMonkeyQueue, item)
			}
		}
		equation := strings.Split(rootEquation, " ")
		firstValue := tempMonkeyValues[equation[0]]
		secondValue := tempMonkeyValues[equation[2]]
		if firstValue == secondValue {
			fmt.Println("You should yell", mid)
			break
		} 
		if firstValue > secondValue {
			low = mid + 1
		} else {
			high = mid
		}
	}
}
func copyMap(m map[string]float64) map[string]float64 {
	mapCopy := make(map[string]float64)
	for k, v := range m {
		mapCopy[k] = v
	}
	return mapCopy
}
