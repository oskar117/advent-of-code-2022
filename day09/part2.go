package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	Right Direction = "R"
	Left            = "L"
	Up              = "U"
	Down            = "D"
)

type Position struct {
	x int
	y int
}

func (position *Position) move(direction Direction, steps int) {
    switch direction {
    case Right:
        position.x += steps
    case Left:
        position.x -= steps
    case Up:
        position.y += steps
    case Down:
        position.y -= steps
    }
}

func (position *Position) moveCloserTo(targetPosition Position) {
    if position.x == targetPosition.x {
        position.y += (targetPosition.y - position.y) / 2
    } else if position.y == targetPosition.y {
        position.x += (targetPosition.x - position.x) / 2
    } else {
        if position.x > targetPosition.x {
            position.x--
        } else {
            position.x++
        }
        if position.y > targetPosition.y {
            position.y--
        } else {
            position.y++
        }
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

    knots := make([]*Position, 10)
    for i := 0; i < 10; i++ {
        knots[i] = &Position{0, 0}
    }

    tail_visited_positions := make(map[Position]bool)
    tail_visited_positions[*knots[len(knots) - 1]] = true

	for scanner.Scan() {
		output := strings.Fields(scanner.Text())
		direction := output[0]
		steps, _ := strconv.Atoi(output[1])
        for i := 0; i < steps; i++ {
            knots[0].move(Direction(direction), 1)
            for knot_index, knot := range knots[1:] {
                earlier_knot := knots[knot_index]
                if canMoveCloser(*earlier_knot, *knot) {
                    knot.moveCloserTo(*earlier_knot)
                    if knot_index == len(knots) - 2 {
                        tail_visited_positions[*knot] = true
                    }
                }
            }
        }
	}

	fmt.Println("Tail visited", len(tail_visited_positions), "positions")
	file.Close()
}

func canMoveCloser(p1, p2 Position) bool {
    if math.Abs(float64(p1.x - p2.x)) <= 1 && math.Abs(float64(p1.y - p2.y)) <= 1 {
        return false
    }
    return true
}
