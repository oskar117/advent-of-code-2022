package main

import (
	"fmt"
	"math"
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
	var playerX, playerY, bestSignalX, bestSignalY int

	for y, line := range strings.Split(strings.TrimSpace(string(file)), "\n") {
		heighmap = append(heighmap, make([]rune, len(line)))
		for x, char := range line {
			if char == 'S' {
				heighmap[y][x] = 'a'
				playerX = x
				playerY = y
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
	queue = append(queue, data{0, playerX, playerY})
	was_here[playerY][playerX] = true

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
			if heighmap[neighbour.nextY][neighbour.nextX]-heighmap[item.y][item.x] > 1 {
				continue
			}
			if neighbour.nextX == bestSignalX && neighbour.nextY == bestSignalY {
				fmt.Println("Shortest path is", item.distance + 1)
				os.Exit(0)
			}
			was_here[neighbour.nextY][neighbour.nextX] = true
			queue = append(queue, data{item.distance + 1, neighbour.nextX, neighbour.nextY})
		}
	}

	os.Exit(2)
	shortest_path := findShortestPath(heighmap, playerX, playerY, bestSignalX, bestSignalY, 0, &was_here, math.MaxInt)

	fmt.Println("Shortest path is", shortest_path)
}

func findShortestPath(heighmap [][]rune, x, y, finishX, finishY, currentLength int, was_here *[][]bool, current_min int) int {
	if x == finishX && y == finishY {
		current_min = int(math.Min(float64(current_min), float64(currentLength)))
		return current_min
	}
	(*was_here)[y][x] = true
	if x != len(heighmap[0])-1 && !(*was_here)[y][x+1] {
		if canWalkThere(heighmap, x, y, x+1, y) {
			current_min = findShortestPath(heighmap, x+1, y, finishX, finishY, currentLength+1, was_here, current_min)
		}
	}
	if y != len(heighmap)-1 && !(*was_here)[y+1][x] {
		if canWalkThere(heighmap, x, y, x, y+1) {
			current_min = findShortestPath(heighmap, x, y+1, finishX, finishY, currentLength+1, was_here, current_min)
		}
	}
	if x != 0 && !(*was_here)[y][x-1] {
		if canWalkThere(heighmap, x, y, x-1, y) {
			current_min = findShortestPath(heighmap, x-1, y, finishX, finishY, currentLength+1, was_here, current_min)
		}
	}
	if y != 0 && !(*was_here)[y-1][x] {
		if canWalkThere(heighmap, x, y, x, y-1) {
			current_min = findShortestPath(heighmap, x, y-1, finishX, finishY, currentLength+1, was_here, current_min)
		}
	}
	// (*was_here)[y][x] = false
	return current_min
}

func canWalkThere(heighmap [][]rune, x, y, nextX, nextY int) bool {
	current := heighmap[y][x]
	next := heighmap[nextY][nextX]
	diff := next - current
	return diff == 1 || diff == 0
}
