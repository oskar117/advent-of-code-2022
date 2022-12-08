package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := make([][]int, 0)

	index := 0
	for scanner.Scan() {
		temp_line := make([]int, 0)
		for _, height := range scanner.Text() {
			n, _ := strconv.Atoi(string(height))
			temp_line = append(temp_line, n)
		}
		grid = append(grid, temp_line)
		index++
	}
	file.Close()

	len_x, len_y := len(grid[0]), len(grid)
	max_score := 0

	for y := 0; y < len_y; y++ {
		for x := 0; x < len_x; x++ {
			scenic_score := 1
			tree_height := grid[x][y]
			for x1 := x - 1; x1 >= 0; x1-- {
				if grid[x1][y] >= tree_height || x1 == 0 {
					scenic_score *= x - x1
					break
				}
			}
			for x2 := x + 1; x2 < len_x; x2++ {
				if grid[x2][y] >= tree_height || x2 == len_x-1 {
					scenic_score *= x2 - x
					break
				}
			}
			for y1 := y - 1; y1 >= 0; y1-- {
				if grid[x][y1] >= tree_height || y1 == 0 {
					scenic_score *= y - y1
					break
				}
			}
			for y2 := y + 1; y2 < len_y; y2++ {
				if grid[x][y2] >= tree_height || y2 == len_y-1 {
					scenic_score *= y2 - y
					break
				}
			}
			if scenic_score > max_score {
				max_score = scenic_score
			}
		}
	}

	fmt.Println("Highest scenic score is", max_score)
}

func check_neighbours() bool {
	return false
}
