package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}

func calculateNeighbours(c Cube) []Cube {
	return []Cube{
		{c.x + 1, c.y, c.z},
		{c.x, c.y + 1, c.z},
		{c.x, c.y, c.z + 1},
		{c.x - 1, c.y, c.z},
		{c.x, c.y - 1, c.z},
		{c.x, c.y, c.z - 1},
	}
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := 0

	grid := make(map[Cube]bool)

	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), ",")
		x, y, z := getCoord(splitText[0]), getCoord(splitText[1]), getCoord(splitText[2])
		cube := Cube{x, y, z}
		grid[cube] = true
		result += 6
		for _, neighbour := range calculateNeighbours(cube) {
			if grid[neighbour] {
				result -= 2
			}
		}
	}

	fmt.Println("Total surface is", result)
	file.Close()
}

func getCoord(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}
