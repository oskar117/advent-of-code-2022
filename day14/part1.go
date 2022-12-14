package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Field struct {
	x int
	y int
}

func NewField(str string) *Field {
	params := strings.Split(str, ",")
	x, _ := strconv.Atoi(params[0])
	y, _ := strconv.Atoi(params[1])
	return &Field{x, y}
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	grid := make(map[Field]bool)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	max_y := 0

	for scanner.Scan() {
		text := scanner.Text()
		instructions := strings.Split(text, " -> ")
		for i := 1; i < len(instructions); i++ {
			first := NewField(instructions[i-1])
			second := NewField(instructions[i])
			startX, endX := func() (int, int) {
				if first.x > second.x {
					return second.x, first.x
				} else {
					return first.x, second.x
				}
			}()
			startY, endY := func() (int, int) {
				if first.y > second.y {
					return second.y, first.y
				} else {
					return first.y, second.y
				}
			}()
			if startX == endX {
				for i := startY; i <= endY; i++ {
					grid[Field{startX, i}] = true
				}
			} else if startY == endY {
				for i := startX; i <= endX; i++ {
					grid[Field{i, startY}] = true
				}
			} else {
				fmt.Println("ERROR, invalid input")
			}
			if max_y < startY {
				max_y = startY
			} else if max_y < endY {
				max_y = endY
			}
		}
	}
	file.Close()

	sand_placed := 0
	for placeSand(500, 0, &grid, max_y) {
		sand_placed++
	}

	fmt.Println("There were", sand_placed, "sands placed.")
}

func placeSand(x, y int, grid *map[Field]bool, minY int) bool {
	if minY < y {
		return false
	}
	isElementAtNextPos := (*grid)[Field{x, y+1}]
	if isElementAtNextPos {
		isElementAtLeft := (*grid)[Field{x-1, y+1}]
		if !isElementAtLeft {
			return placeSand(x-1, y, grid, minY)
		}
		isElementAtRight := (*grid)[Field{x+1, y+1}]
		if !isElementAtRight {
			return placeSand(x+1, y, grid, minY)
		}
		(*grid)[Field{x, y}] = true
		return true
	} else {
		return placeSand(x, y+1, grid, minY)
	}
}
