package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	x, y int
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)
	line_number, _ := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	coordinateMap := make(map[Coordinate]bool)

	for scanner.Scan() {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distance := calculateDistance(sensorX, sensorY, beaconX, beaconY)
		for y := sensorY - distance; y <= sensorY+distance; y++ {
			if y == line_number {
				offset := distance - int(math.Abs(float64(sensorY)-float64(y))) 
				for x := sensorX - offset; x <= sensorX+offset; x++ {
					coordinateMap[Coordinate{x, y}] = true
				}
			}
		}
		coordinateMap[Coordinate{sensorX, sensorY}] = false
		coordinateMap[Coordinate{beaconX, beaconY}] = false
	}

	minX, maxX, _, _ := findMinAndMax(coordinateMap)
	takenPositionsCounter := 0
	for x := minX; x <= maxX; x++ {
		if coordinateMap[Coordinate{x, line_number}] {
			takenPositionsCounter++
		}
	}

	fmt.Println("At line", line_number, "there are", takenPositionsCounter, "occupied positions.")
	file.Close()
}

func calculateDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1)))
}

func findMinAndMax(coords map[Coordinate]bool) (int, int, int, int) {
	minX := math.MaxInt
	maxX := math.MinInt
	minY := math.MaxInt
	maxY := math.MinInt
	for k := range coords {
		if k.x < minX {
			minX = k.x
		}
		if k.x > maxX {
			maxX = k.x
		}
		if k.y < minY {
			minY = k.y
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	return minX, maxX, minY, maxY
}
