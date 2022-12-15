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

type Sensor struct {
	position Coordinate
	radius   int
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)
	max_coord, _ := strconv.Atoi(os.Args[2])

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sensors := make([]Sensor, 0)
	minX, maxX, minY, maxY := 0, max_coord, 0, max_coord

	for scanner.Scan() {
		var sensorX, sensorY, beaconX, beaconY int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sensorX, &sensorY, &beaconX, &beaconY)
		distance := calculateDistance(sensorX, sensorY, beaconX, beaconY)
		sensors = append(sensors, Sensor{Coordinate{sensorX, sensorY}, distance})
	}
	file.Close()

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			found := true
			for _, sensor := range sensors {
				dist := calculateDistance(x, y, sensor.position.x, sensor.position.y)
				if dist <= sensor.radius {
					found = false
					x += sensor.radius - int(math.Abs(float64(sensor.position.y-y))) + sensor.position.x - x
					break	
				}
			}
			if found {
				fmt.Println("Tuning frequency is", x*4000000+y)
				os.Exit(0)
			}
		}
	}
	os.Exit(1)
}

func calculateDistance(x1, y1, x2, y2 int) int {
	return int(math.Abs(float64(x2)-float64(x1)) + math.Abs(float64(y2)-float64(y1)))
}
