package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cube struct {
	x, y, z int
}
type Coord struct {
	a, b int
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

	grid := make(map[Cube]bool)
	result := 0
	minX, minY, minZ := math.MaxInt, math.MaxInt, math.MaxInt
	maxX, maxY, maxZ := math.MinInt, math.MinInt, math.MinInt

	for scanner.Scan() {
		splitText := strings.Split(scanner.Text(), ",")
		x, y, z := getCoord(splitText[0]), getCoord(splitText[1]), getCoord(splitText[2])
		cube := Cube{x, y, z}
		grid[cube] = true
		minX = min(minX, x)
		minY = min(minY, y)
		minZ = min(minZ, z)
		maxX = max(maxX, x)
		maxY = max(maxY, y)
		maxZ = max(maxZ, z)
		result += 6
		for _, neighbour := range calculateNeighbours(cube) {
			if grid[neighbour] {
				result -= 2
			}
		}
	}

	minX--
	minY--
	minZ--
	maxX++
	maxY++
	maxZ++

	airSpace := make(map[Cube]bool)
	for z := minZ; z <= maxZ; z++ {
		for x := minX; x <= maxX; x++ {
			for y := minY; y <= maxY; y++ {
				pos := Cube{x, y, z}
				if !grid[pos] {
					airSpace[pos] = true
				}
			}
		}
	}
	removeOpenAirSpace(&airSpace, 0, 0, 0, minX, minY, minZ, maxX, maxY, maxZ)

	calculatedAirSpace := make(map[Cube]bool)
	for cube := range airSpace {
		calculatedAirSpace[cube] = true
		result -= 6
		for _, neighbour := range calculateNeighbours(cube) {
			if calculatedAirSpace[neighbour] {
				result += 2
			}
		}
	}

	fmt.Println("Total surface is", result)
	file.Close()
}

func removeOpenAirSpace(grid *map[Cube]bool, x, y, z, minX, minY, minZ, maxX, maxY, maxZ int) {
	if x > maxX || x < minX || y > maxY || y < minY || z > maxZ || z < minZ {
		return
	}
	val, ok := (*grid)[Cube{x, y, z}]
	if ok && val {
		delete(*grid, Cube{x, y, z})
	}
	if ok {
		removeOpenAirSpace(grid, x+1, y, z, minX, minY, minZ, maxX, maxY, maxZ)
		removeOpenAirSpace(grid, x-1, y, z, minX, minY, minZ, maxX, maxY, maxZ)
		removeOpenAirSpace(grid, x, y+1, z, minX, minY, minZ, maxX, maxY, maxZ)
		removeOpenAirSpace(grid, x, y-1, z, minX, minY, minZ, maxX, maxY, maxZ)
		removeOpenAirSpace(grid, x, y, z+1, minX, minY, minZ, maxX, maxY, maxZ)
		removeOpenAirSpace(grid, x, y, z-1, minX, minY, minZ, maxX, maxY, maxZ)
	}
}

func getCoord(str string) int {
	val, _ := strconv.Atoi(str)
	return val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
