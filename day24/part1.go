package main

import (
	"fmt"
	"os"
	"strings"
)

type Data struct {
	minute   int
	position Position
}

type Position struct {
	x, y int
}

type Blizzard struct {
	direction rune
	position  Position
}

func get_neighbours(p Position) []Position {
	return []Position{
		{p.x, p.y + 1},
		{p.x, p.y - 1},
		{p.x + 1, p.y},
		{p.x - 1, p.y},
		{p.x, p.y},
	}
}

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}
	stringLines := strings.Split(strings.TrimSpace(string(file)), "\n")
	blizzards := make([]Blizzard, 0)
	walls := make(map[Position]bool)
	player := Position{1, 0}
	finish := Position{len(stringLines[0]) - 2, len(stringLines) - 1}

	for y, line := range stringLines {
		for x, char := range line {
			switch char {
			case '#':
				walls[Position{x, y}] = true
			case '<', '>', 'v', '^':
				blizzards = append(blizzards, Blizzard{char, Position{x, y}})
			}
		}
	}

	was_here := make(map[Data]bool)
	queue := make([]Data, 0)
	queue = append(queue, Data{1, player})
	blizzardMap := make(map[int][]Blizzard)
	blizzardMap[0] = blizzards

	for len(queue) > 0 {
		currPos := queue[0]
		queue = queue[1:]
		currBlizzards, exists := blizzardMap[currPos.minute]
		if !exists {
			currBlizzards = make([]Blizzard, len(blizzards))
			for i, blizzard := range blizzardMap[currPos.minute-1] {
				var nextPos Position
				switch blizzard.direction {
				case '>':
					nextPos = Position{blizzard.position.x + 1, blizzard.position.y}
					if walls[nextPos] {
						nextPos.x = 1
					}
				case '<':
					nextPos = Position{blizzard.position.x - 1, blizzard.position.y}
					if walls[nextPos] {
						nextPos.x = len(stringLines[0]) - 2
					}
				case 'v':
					nextPos = Position{blizzard.position.x, blizzard.position.y + 1}
					if walls[nextPos] {
						nextPos.y = 1
					}
				case '^':
					nextPos = Position{blizzard.position.x, blizzard.position.y - 1}
					if walls[nextPos] {
						nextPos.y = len(stringLines) - 2
					}
				}
				currBlizzards[i] = Blizzard{blizzard.direction, nextPos}
			}
			blizzardMap[currPos.minute] = currBlizzards
		}
		for _, neighbour := range get_neighbours(currPos.position) {
			nextData := Data{currPos.minute + 1, neighbour}
			if neighbour.y < 0 || neighbour.x < 0 || neighbour.y >= len(stringLines) || neighbour.x >= len(stringLines[0]) {
				continue
			}
			if walls[neighbour] {
				continue
			}
			if was_here[nextData] {
				continue
			}
			if isNextPlaceBlizzard(currBlizzards, neighbour) {
				continue
			}
			if neighbour == finish {
				fmt.Println("Shortest path takes", currPos.minute, "minutes")
				os.Exit(0)
			}
			was_here[nextData] = true
			queue = append(queue, nextData)
		}
	}

	os.Exit(2)

}

func isNextPlaceBlizzard(blizzards []Blizzard, pos Position) bool {
	for _, blizzard := range blizzards {
		if pos == blizzard.position {
			return true
		}
	}
	return false
}
