package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Direction byte

const (
	Right Direction = iota
	Down
	Left
	Up
)

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	splitString := strings.Split(string(file), "\n\n")
	longestLine := maxLineLength(splitString[0])
	splitGrid := strings.Split(splitString[0], "\n")
	grid := make([][]byte, len(splitGrid))
	steps, rotation := make([]int, 0), make([]rune, 0)

	for lineIndex, line := range splitGrid {
		grid[lineIndex] = make([]byte, longestLine)
		for charIndex, char := range line {
			switch char {
			case '.':
				grid[lineIndex][charIndex] = 1
			case '#':
				grid[lineIndex][charIndex] = 2
			}
		}
	}
	tempNumberString := ""
	for _, char := range strings.TrimSpace(splitString[1]) {
		if unicode.IsLetter(char) {
			rotation = append(rotation, char)
			if tempNumberString != "" {
				num, _ := strconv.Atoi(tempNumberString)
				steps = append(steps, num)
				tempNumberString = ""
			}
		} else {
			tempNumberString += string(char)
		}
	}
	if tempNumberString != "" {
		num, _ := strconv.Atoi(tempNumberString)
		steps = append(steps, num)
	}

	locX, locY := 0, 0
	locDir := Right
	for len(steps) != 0 || len(rotation) != 0 {
		if len(steps) != 0 {
			step := steps[0]
			steps = steps[1:]
			for i := 0; i < step; i++ {
				ok, x, y := walk(grid, locX, locY, locDir)
				if !ok {
					break
				}
				locX, locY = x, y
			}
		}
		if len(rotation) != 0 {
			rot := rotation[0]
			rotation = rotation[1:]
			switch rot {
			case 'R':
				locDir = (locDir + 1) % 4
			case 'L':
				locDir = (locDir - 1) % 4
			}
		}
	}

	result := 1000*(locY+1) + 4*(locX+1) + int(locDir)
	fmt.Println("Final password is", result)
}

func walk(grid [][]byte, x, y int, dir Direction) (bool, int, int) {
	switch dir {
	case Right:
		if x+1 >= len(grid[y]) {
			x -= len(grid[y])
		}
		if grid[y][x+1] == 2 {
			return false, x, y
		} else if grid[y][x+1] == 0 {
			return walk(grid, x+1, y, dir)
		}
		return true, x + 1, y
	case Left:
		if x-1 < 0 {
			x += len(grid[y])
		}
		if grid[y][x-1] == 2 {
			return false, x, y
		} else if grid[y][x-1] == 0 {
			return walk(grid, x-1, y, dir)
		}
		return true, x - 1, y
	case Down:
		if y+1 >= len(grid) {
			y -= len(grid)
		}
		if grid[y+1][x] == 2 {
			return false, x, y
		} else if grid[y+1][x] == 0 {
			return walk(grid, x, y+1, dir)
		}
		return true, x, y + 1
	case Up:
		if y-1 < 0 {
			y += len(grid)
		}
		if grid[y-1][x] == 2 {
			return false, x, y
		} else if grid[y-1][x] == 0 {
			return walk(grid, x, y-1, dir)
		}
		return true, x, y - 1
	}
	return false, x, y
}

func maxLineLength(input string) (size int) {
	lines := strings.Split(input, "\n")
	for _, v := range lines {
		if len(v) >= size {
			size = len(v)
		}
	}
	return
}
