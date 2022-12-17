package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Valve struct {
	name        string
	flowRate    int
	connections []string
}

type Route struct {
	name     string
	distance int
}

type Path struct {
	valve       string
	pressure    int
	sumPressure int
	minute      int
	open        [16]bool
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	valves := make(map[string]*Valve)

	for scanner.Scan() {
		var name string
		var flowRate int
		input_parts := strings.Split(scanner.Text(), "; ")
		fmt.Sscanf(input_parts[0], "Valve %s has flow rate=%d", &name, &flowRate)
		connections := strings.Split(strings.TrimSpace(input_parts[1][22:]), ", ")
		valves[name] = &Valve{name, flowRate, connections}
	}
	file.Close()

	dists := make(map[string]map[string]int)
	nonEmptyValves := make([]string, 0)

	for name, valve := range valves {
		if name != "AA" && valve.flowRate == 0 {
			continue
		}

		if name != "AA" {
			nonEmptyValves = append(nonEmptyValves, name)
		}

		dists[name] = map[string]int{"AA": 0}
		visited := make(map[string]bool)
		visited[name] = true

		queue := make([]Route, 0)
		queue = append(queue, Route{name, 0})

		for len(queue) > 0 {
			item := queue[0]
			queue = queue[1:]
			for _, neighbor := range valves[item.name].connections {
				if visited[neighbor] {
					continue
				}
				visited[neighbor] = true
				if valves[neighbor].flowRate != 0 {
					dists[name][neighbor] = item.distance + 1
				}
				queue = append(queue, Route{neighbor, item.distance + 1})
			}
		}
		delete(dists[name], name)
		if name != "AA" {
			delete(dists[name], "AA")
		}
	}

	indicies := make(map[string]int)

	for index, element := range nonEmptyValves {
		indicies[element] = index
	}

	maxPressure := getMaxPressure(30, "AA", 0, valves, dists, indicies)
	fmt.Println("Max pressure is", maxPressure)
}

func getMaxPressure(time int, name string, bitmask int,
	valves map[string]*Valve, dists map[string]map[string]int, indicies map[string]int) int {
	max := 0
	for neighbor, distance := range dists[name] {
		bit := 1 << indicies[neighbor]
		if bitmask&bit != 0 {
			continue
		}
		remeaningTime := time - distance - 1
		if remeaningTime <= 0 {
			continue
		}
		max = int(math.Max(float64(max),
			float64(getMaxPressure(remeaningTime, neighbor, bitmask|bit, valves, dists, indicies)+valves[neighbor].flowRate*remeaningTime)))
	}
	return max
}
