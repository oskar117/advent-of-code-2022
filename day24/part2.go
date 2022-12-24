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

	blizzardMap := make(map[int][]Blizzard)
	blizzardMap[0] = blizzards

	toExit := measureTime(&blizzardMap, player, finish, walls, len(stringLines[0]), len(stringLines), 0)
	fromExit := measureTime(&blizzardMap, finish, player, walls, len(stringLines[0]), len(stringLines), toExit + 1)
	againToExit := measureTime(&blizzardMap, player, finish, walls, len(stringLines[0]), len(stringLines), fromExit + 1)
	fmt.Println("You can go back and forth in", againToExit, "minutes")
}

func measureTime(blizzardMap *map[int][]Blizzard, player, finish Position, walls map[Position]bool, maxX, maxY, initMinute int) int {
	was_here := make(map[Data]bool)
	queue := make([]Data, 0)
	queue = append(queue, Data{initMinute, player})

	for len(queue) > 0 {
		currPos := queue[0]
		queue = queue[1:]
		currBlizzards, exists := (*blizzardMap)[currPos.minute]
		if !exists {
			currBlizzards = make([]Blizzard, len((*blizzardMap)[0]))
			for i, blizzard := range (*blizzardMap)[currPos.minute-1] {
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
						nextPos.x = maxX - 2
					}
				case 'v':
					nextPos = Position{blizzard.position.x, blizzard.position.y + 1}
					if walls[nextPos] {
						nextPos.y = 1
					}
				case '^':
					nextPos = Position{blizzard.position.x, blizzard.position.y - 1}
					if walls[nextPos] {
						nextPos.y = maxY - 2
					}
				}
				currBlizzards[i] = Blizzard{blizzard.direction, nextPos}
			}
			(*blizzardMap)[currPos.minute] = currBlizzards
		}
		for _, neighbour := range get_neighbours(currPos.position) {
			nextData := Data{currPos.minute + 1, neighbour}
			if neighbour.y < 0 || neighbour.x < 0 || neighbour.y >= maxY || neighbour.x >= maxX {
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
				return currPos.minute
			}
			was_here[nextData] = true
			queue = append(queue, nextData)
		}
	}
	return -1
}

func isNextPlaceBlizzard(blizzards []Blizzard, pos Position) bool {
	for _, blizzard := range blizzards {
		if pos == blizzard.position {
			return true
		}
	}
	return false
}
