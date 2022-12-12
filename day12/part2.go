package main

import (
	"fmt"
	"os"
	"strings"
)

type data struct {
	distance int
	x        int
	y        int
}

type next struct {
	nextX int
	nextY int
}

func get_neighbours(x, y int) []next {
	return []next{
		{x, y + 1},
		{x, y - 1},
		{x + 1, y},
		{x - 1, y},
	}
}

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	heighmap := make([][]rune, 0)
	var bestSignalX, bestSignalY int

	for y, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		heighmap = append(heighmap, make([]rune, len(line)))
		for x, char := range line {
			if char == 'S' {
				heighmap[y][x] = 'a'
			} else if char == 'E' {
				heighmap[y][x] = 'z'
				bestSignalX = x
				bestSignalY = y
			} else {
				heighmap[y][x] = char
			}
		}
	}

	was_here := make([][]bool, len(heighmap))
	for i := 0; i < len(heighmap); i++ {
		was_here[i] = make([]bool, len(heighmap[i]))
	}

	queue := make([]data, 0)
	queue = append(queue, data{0, bestSignalX, bestSignalY})
	was_here[bestSignalY][bestSignalX] = true

	for len(queue) != 0 {
		item := queue[0]
		queue = queue[1:]
		for _, neighbour := range get_neighbours(item.x, item.y) {
			if neighbour.nextY < 0 || neighbour.nextX < 0 || neighbour.nextY >= len(heighmap) || neighbour.nextX >= len(heighmap[0]) {
				continue
			}
			if was_here[neighbour.nextY][neighbour.nextX] {
				continue
			}
			if heighmap[neighbour.nextY][neighbour.nextX]-heighmap[item.y][item.x] < -1 {
				continue
			}
			if heighmap[neighbour.nextY][neighbour.nextX] == 'a' {
				fmt.Println("Shortest path is", item.distance + 1)
				os.Exit(0)
			}
			was_here[neighbour.nextY][neighbour.nextX] = true
			queue = append(queue, data{item.distance + 1, neighbour.nextX, neighbour.nextY})
		}
	}

	os.Exit(2)
}

