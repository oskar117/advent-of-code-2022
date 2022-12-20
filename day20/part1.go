package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type NumberWithIndex struct {
	index, value int
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	initial_state := make([]int, 0)
	numbers := make([]NumberWithIndex, 0)

	i := 0
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		initial_state = append(initial_state, value)
		numbers = append(numbers, NumberWithIndex{i, value})	
		i++
	}
	file.Close()


	for i, number := range initial_state {
		index := indexOf(numbers, i, number)
		nextIndex := (len(numbers) - 1 + index + number) % (len(numbers) - 1)
		if nextIndex < 0 {
			nextIndex += len(numbers) - 1
		}
		numbers = append(numbers[:index], numbers[index+1:]...)
		numbers = append(numbers[:nextIndex], append([]NumberWithIndex{{-1, number}}, numbers[nextIndex:]...)...)
	}

	result := 0
	zeroIndex := indexOf(numbers, -1, 0)
	for i := 1000; i <= 3000; i += 1000 {
		x := numbers[(zeroIndex+i)%(len(numbers))]
		result += x.value
	}

	fmt.Println("Grove coordinates are", result)
}

func indexOf(slice []NumberWithIndex, index int, item int) int {
	for i, val := range slice {
		if val.value == item && index == val.index {
			return i
		}
	}
	return -1
}
