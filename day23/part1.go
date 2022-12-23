package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type Direction int

const (
	North Direction = iota
	South
	West
	East
)

type Position struct {
	x, y int
}

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	gridLines := strings.Split(string(file), "\n")
	elves := make(map[Position]bool)
	minX, minY, maxX, maxY := math.MaxInt, math.MaxInt, math.MinInt, math.MinInt

	for lineIndex, line := range gridLines {
		for charIndex, char := range line {
			if char == '#' {
				elves[Position{charIndex, lineIndex}] = true
			}
		}
	}

	for round := 0; round < 10; round++ {
		proposedPositions := map[Position]Position{}
		onPosCounter := 0
		for elf := range elves {
			n, is_n := elves[Position{elf.x, elf.y - 1}]
			s, is_s := elves[Position{elf.x, elf.y + 1}]
			w, is_w := elves[Position{elf.x - 1, elf.y}]
			e, is_e := elves[Position{elf.x + 1, elf.y}]
			ne, is_ne := elves[Position{elf.x + 1, elf.y - 1}]
			se, is_se := elves[Position{elf.x + 1, elf.y + 1}]
			nw, is_nw := elves[Position{elf.x - 1, elf.y - 1}]
			sw, is_sw := elves[Position{elf.x - 1, elf.y + 1}]
			if !is_n && !is_ne && !is_nw && !is_s && !is_se && !is_sw && !is_w && !is_e {
				onPosCounter++
				continue
			}
		elfLoop:
			for i := 0; i < 4; i++ {
				switch Direction((round + i) % 4) {
				case North:
					if !n && !ne && !nw {
						proposedPositions[elf] = Position{elf.x, elf.y - 1}
						break elfLoop
					}
					break
				case South:
					if !s && !se && !sw {
						proposedPositions[elf] = Position{elf.x, elf.y + 1}
						break elfLoop
					}
					break
				case West:
					if !w && !nw && !sw {
						proposedPositions[elf] = Position{elf.x - 1, elf.y}
						break elfLoop
					}
					break
				case East:
					if !e && !ne && !se {
						proposedPositions[elf] = Position{elf.x + 1, elf.y}
						break elfLoop
					}
					break
				}
			}
		}
		if len(elves) == onPosCounter {
			break
		}
		legalPositions := map[Position]bool{}
		for _, proposedPosition := range proposedPositions {
			if _, has := legalPositions[proposedPosition]; has {
				delete(legalPositions, proposedPosition)
			} else {
				legalPositions[proposedPosition] = true
			}
		}
		for elf, proposedPosition := range proposedPositions {
			if _, has := legalPositions[proposedPosition]; has {
				delete(elves, elf)
				elves[proposedPosition] = true
			}
		}
		for elf := range elves {
			if elf.x > maxX {
				maxX = elf.x
			}
			if elf.y > maxY {
				maxY = elf.y
			}
			if elf.y < minY {
				minY = elf.y
			}
			if elf.x < minX {
				minX = elf.x
			}
		}
	}
	groundTiles := 0
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			if _, is := elves[Position{x, y}]; is {
				// fmt.Print("#")
			} else {
				// fmt.Print(".")
				groundTiles++
			}
		}
		// fmt.Print("\n")
	}

	fmt.Println("There are", groundTiles, "empty tiles in the rectangle.")
}
