package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	containers, instructions := func() ([]string, []string) {
		parts := strings.Split(string(file), "\n\n")
		return trimStringToArray(parts[0]), trimStringToArray(parts[1])
	}()

	stack_num := (len(containers[0]) + 1) >> 2
	stacks := make([][]string, stack_num)
	for i := range stacks {
		stacks[i] = make([]string, 0)
	}

	for i := len(containers) - 1; i >= 0; i-- {
		containers[i] += " "
		line := func() []string {
			var res []string
			for c := 0; c <= len(containers[i])-1; c += 4 {
				res = append(res, containers[i][c:c+4])
			}
			return res
		}()
		for line_index, container := range line {
			if strings.TrimSpace(container) != "" {
				stacks[line_index] = append(stacks[line_index], container)
			}
		}
	}

	for _, instruction := range instructions {
		elements := strings.Fields(instruction)
		source_id, _ := strconv.Atoi(elements[3])
		target_id, _ := strconv.Atoi(elements[5])
		source_id -= 1
		target_id -= 1
		element_num, _ := strconv.Atoi(elements[1])
		for i := 0; i < element_num; i++ {
			pop_element := stacks[source_id][len(stacks[source_id])-1]
			stacks[source_id] = stacks[source_id][:len(stacks[source_id])-1]
			stacks[target_id] = append(stacks[target_id], pop_element)
		}
	}

	res := ""
	for _, stack := range stacks {
		res += strings.Split(stack[len(stack) - 1], "")[1]
	}
	fmt.Println("Message is", res)
}

func trimStringToArray(data string) []string {
	lines := strings.Split(data, "\n")
	return lines[:len(lines)-1]
}
