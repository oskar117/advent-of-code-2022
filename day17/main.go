package main

import (
	"fmt"
	"os"
)

type Properties struct {
	x, y, sizeX, sizeY int
}

type Rock interface {
	getShape() [][]bool
	setPosition(int, int)
	getPosition() (int, int)
}

type Shape struct {
	Properties
}

func (shape *Shape) setPosition(x, y int) {
	shape.x = x
	shape.y = y
}

func (shape *Shape) getPosition() (int, int) {
	return shape.x, shape.y
}

type MinusShape struct {
	Shape
}

func (shape *MinusShape) getShape() [][]bool {
	return [][]bool{
		{true, true, true, true},
	}
}

type PlusShape struct {
	Shape
}

func (shape *PlusShape) getShape() [][]bool {
	return [][]bool{
		{false, true, false},
		{true, true, true},
		{false, true, false},
	}
}

type ReversedLShape struct {
	Shape
}

func (shape *ReversedLShape) getShape() [][]bool {
	return [][]bool{
		{false, false, true},
		{false, false, true},
		{true, true, true},
	}
}

type WallShape struct {
	Shape
}

func (shape *WallShape) getShape() [][]bool {
	return [][]bool{
		{true},
		{true},
		{true},
		{true},
	}
}

type SquareShape struct {
	Shape
}

func (shape *SquareShape) getShape() [][]bool {
	return [][]bool{
		{true, true},
		{true, true},
	}
}

func nextShape(rockNumber int) Rock {
	index := rockNumber % 5
	switch index {
	case 0:
		return &MinusShape{Shape{Properties{sizeX: 4, sizeY: 1}}}
	case 1:
		return &PlusShape{Shape{Properties{sizeX: 3, sizeY: 3}}}
	case 2:
		return &ReversedLShape{Shape{Properties{sizeX: 3, sizeY: 3}}}
	case 3:
		return &WallShape{Shape{Properties{sizeX: 1, sizeY: 4}}}
	case 4:
		return &SquareShape{Shape{Properties{sizeX: 2, sizeY: 2}}}
	}
	return nil
}

func moveLeft(rock Rock, grid [][]bool) bool {
	posX, posY := rock.getPosition()
	shape := rock.getShape()
	if posX-1 < 0 {
		return false
	}
	if posY+len(shape) > len(grid) {
		return false
	}
	for ir, row := range shape {
		for ic, col := range row {
			if col && grid[posY+ir][posX-1+ic] {
				return false
			}
		}
	}
	rock.setPosition(posX-1, posY)
	return true
}

func moveRight(rock Rock, grid [][]bool) bool {
	posX, posY := rock.getPosition()
	shape := rock.getShape()
	if posX+1+len(shape[0]) > len(grid[0]) {
		return false
	}
	if posY+len(shape) > len(grid) {
		return false
	}
	for ir, row := range shape {
		for ic, col := range row {
			if col && grid[posY+ir][posX+1+ic] {
				return false
			}
		}
	}
	rock.setPosition(posX+1, posY)
	return true
}

func moveDown(rock Rock, grid [][]bool) bool {
	posX, posY := rock.getPosition()
	shape := rock.getShape()
	if posY+1+len(shape) > len(grid) {
		return false
	}
	for ir, row := range shape {
		for ic, col := range row {
			if col && grid[posY+1+ir][posX+ic] {
				return false
			}
		}
	}
	rock.setPosition(posX, posY+1)
	return true
}

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}
	pattern := string(file)
	fmt.Print("Part 1: ")
	part1length := 2022
	part2length := 1000000000000
	part1Height, _ := simulate(pattern, part1length, func(i int, patternReset bool) bool { return false })
	fmt.Println("Tower height is", part1Height, "after simulating", part1length, "rocks.")
	fmt.Print("Part 2: ")
	startHeight, startIter := simulate(pattern, part2length, func(i int, patternReset bool) bool { return patternReset && i != 0 })
	nextHeight, nextIter := simulate(pattern, part2length, func(i int, patternReset bool) bool { return patternReset && i > startIter })
	rocksPerCycle := nextIter - startIter
	numberOfCycles := part2length / rocksPerCycle
	totalRocks := rocksPerCycle * numberOfCycles + startIter
	heightPerCycle := nextHeight - startHeight
	totalHeight := heightPerCycle * numberOfCycles + startHeight
	overshoot := totalRocks - part2length
	part2Height, _ := simulate(pattern, part2length, func(i int, patternReset bool) bool { return i == startIter - overshoot})
	part2Res := totalHeight - (startHeight - part2Height)
	fmt.Println("Tower height is", part2Res, "after simulating", part2length, "rocks.")
}

func simulate(pattern string, simulationLength int, finishPredicate func(i int, patternReset bool) bool) (int, int) {
	lenPattern := len(pattern) - 1
	patternIndex := 0
	grid := make([][]bool, 3)
	for i := range grid {
		grid[i] = make([]bool, 7)
	}
	for i := 0; i < simulationLength; i++ {
		distToTop, line_index := distanceToTop(grid)
		rock := nextShape(i)
		if requiredHeight := 3 + len(rock.getShape()); distToTop < requiredHeight {
			for x := 0; x < requiredHeight-distToTop; x++ {
				grid = append(make([][]bool, 1), grid...)
				grid[0] = make([]bool, 7)
				line_index++
			}
		}
		rock.setPosition(2, line_index-3-len(rock.getShape()))
		for {
			if finishPredicate(i, patternIndex % lenPattern == 0) {
				dist, _ := distanceToTop(grid)
				return len(grid) - dist, i
			}
			switch pattern[patternIndex%lenPattern] {
			case '<':
				moveLeft(rock, grid)
			case '>':
				moveRight(rock, grid)
			}
			patternIndex++
			success := moveDown(rock, grid)
			if !success {
				posX, posY := rock.getPosition()
				shape := rock.getShape()
				for ir, row := range shape {
					for ic, col := range row {
						if col && !grid[posY+ir][posX+ic] {
							grid[posY+ir][posX+ic] = true
						}
					}
				}
				break
			}
		}
	}
	dist, _ := distanceToTop(grid)
	return len(grid) - dist, simulationLength
}

func distanceToTop(lines [][]bool) (int, int) {
	dist := 0
	for line_index, line := range lines {
		for _, pos := range line {
			if pos {
				return dist, line_index
			}
		}
		dist++
	}
	return dist, dist
}

func printGrid(grid [][]bool) {
	for _, row := range grid {
		for _, col := range row {
			if col {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}
