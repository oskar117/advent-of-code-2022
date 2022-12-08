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

	visible_trees := 0
	len_x, len_y := len(grid[0]), len(grid)

	for y := 0; y < len_y; y++ {
		for x := 0; x < len_x; x++ {
			tree_height := grid[x][y]
			is_visible := func() bool {
				invisible_counter := 0
				for x1 := 0; x1 < x; x1++ {
					if grid[x1][y] >= tree_height {
						invisible_counter++
						break
					}
				}
				for x2 := x + 1; x2 < len_x; x2++ {
					if grid[x2][y] >= tree_height {
						invisible_counter++
						break
					}
				}
				for y1 := 0; y1 < y; y1++ {
					if grid[x][y1] >= tree_height {
						invisible_counter++
						break
					}
				}
				for y2 := y + 1; y2 < len_y; y2++ {
					if grid[x][y2] >= tree_height {
						invisible_counter++
						break
					}
				}
				return invisible_counter < 4
			}()
			if is_visible {
				visible_trees++
			}
		}
	}

	fmt.Println("There are", visible_trees, "trees visible")
}

func check_neighbours() bool {
	return false
}
