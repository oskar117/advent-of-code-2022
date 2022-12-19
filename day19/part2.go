package main

import (
	"bufio"
	"fmt"
	"os"
)

type Container struct {
	ore, clay, obsidian, geode int
}

func (c1 *Container) add(c2 Container) {
	c1.ore += c2.ore
	c1.clay += c2.clay
	c1.obsidian += c2.obsidian
	c1.geode += c2.geode
}

type Blueprint struct {
	number                                                                 int
	oreRobotCost, clayRobotCost, obsidianRobotCost, geodeRobotCost, limits Container
}

type Robot interface {
	mine() Container
	canBuild(Blueprint, Container, Container) bool
	consumeMaterials(Blueprint, *Container, *Container)
}

type OreRobot struct{}
type ClayRobot struct{}
type ObsidianRobot struct{}
type GeodeRobot struct{}

func (robot *OreRobot) mine() Container {
	return Container{ore: 1}
}
func (robot *ClayRobot) mine() Container {
	return Container{clay: 1}
}
func (robot *ObsidianRobot) mine() Container {
	return Container{obsidian: 1}
}
func (robot *GeodeRobot) mine() Container {
	return Container{geode: 1}
}

func (robot *OreRobot) consumeMaterials(b Blueprint, c *Container, r *Container) {
	r.ore++
	c.ore -= b.oreRobotCost.ore
}
func (robot *ClayRobot) consumeMaterials(b Blueprint, c *Container, r *Container) {
	r.clay++
	c.ore -= b.clayRobotCost.ore
}
func (robot *ObsidianRobot) consumeMaterials(b Blueprint, c *Container, r *Container) {
	r.obsidian++
	c.ore -= b.obsidianRobotCost.ore
	c.clay -= b.obsidianRobotCost.clay
}
func (robot *GeodeRobot) consumeMaterials(b Blueprint, c *Container, r *Container) {
	r.geode++
	c.ore -= b.geodeRobotCost.ore
	c.obsidian -= b.geodeRobotCost.obsidian
}

func (robot *OreRobot) canBuild(b Blueprint, c Container, sum Container) bool {
	return c.ore >= b.oreRobotCost.ore && b.limits.ore > sum.ore
}
func (robot *ClayRobot) canBuild(b Blueprint, c Container, sum Container) bool {
	return c.ore >= b.clayRobotCost.ore && b.limits.clay > sum.clay
}
func (robot *ObsidianRobot) canBuild(b Blueprint, c Container, sum Container) bool {
	return c.ore >= b.obsidianRobotCost.ore && c.clay >= b.obsidianRobotCost.clay && b.limits.obsidian > sum.obsidian
}
func (robot *GeodeRobot) canBuild(b Blueprint, c Container, sum Container) bool {
	return c.ore >= b.geodeRobotCost.ore && c.obsidian >= b.geodeRobotCost.obsidian
}

type State struct {
	container      Container
	robotContainer Container
	minute         int
}

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	blueprints := make([]Blueprint, 0)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	counter := 0
	for scanner.Scan() && counter < 3 {
		text := scanner.Text()
		limits := Container{}
		var number, oreRobotOre, clayRobotOre, obsitianRobotOre, obsidianRobotClay, geodeRobotOre, geodeRobotObsidian int
		fmt.Sscanf(text, "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian.",
			&number, &oreRobotOre, &clayRobotOre, &obsitianRobotOre, &obsidianRobotClay, &geodeRobotOre, &geodeRobotObsidian)
		limits.ore = oreRobotOre
		limits.clay = obsidianRobotClay
		limits.obsidian = geodeRobotObsidian
		if clayRobotOre > limits.ore {
			limits.ore = clayRobotOre
		}
		if obsitianRobotOre > limits.ore {
			limits.ore = obsitianRobotOre
		}
		blueprints = append(blueprints, Blueprint{
			number,
			Container{ore: oreRobotOre},
			Container{ore: clayRobotOre},
			Container{ore: obsitianRobotOre, clay: obsidianRobotClay},
			Container{ore: geodeRobotOre, obsidian: geodeRobotObsidian},
			limits,
		})
		counter++
	}
	file.Close()

	qualityLevel := 1

	for _, blueprint := range blueprints {
		robotContainer := Container{1, 0, 0, 0}
		container := Container{}
		queue := []State{{container, robotContainer, 0}}
		visited := map[State]bool{}
		currentMax := 0
		for len(queue) > 0 {
			state := queue[0]
			queue = queue[1:]
			if visited[state] {
				continue
			}
			visited[state] = true
			if state.container.geode > currentMax {
				currentMax = state.container.geode
			}
			if state.minute >= 32 {
				continue
			}
			if state.container.geode < currentMax {
				continue
			}
			nextState := state.container
			nextState.ore += state.robotContainer.ore
			nextState.clay += state.robotContainer.clay
			nextState.obsidian += state.robotContainer.obsidian
			nextState.geode += state.robotContainer.geode
			geodeRobot := &GeodeRobot{}
			if geodeRobot.canBuild(blueprint, state.container, state.robotContainer) {
				geodeRobot.consumeMaterials(blueprint, &nextState, &state.robotContainer)
				queue = append(queue, State{nextState, state.robotContainer, state.minute + 1})
			} else {
				for _, robot := range []Robot{&OreRobot{}, &ClayRobot{}, &ObsidianRobot{}} {
					robotNextState := nextState
					robotNextContainer := state.robotContainer
					if robot.canBuild(blueprint, state.container, robotNextContainer) {
						robot.consumeMaterials(blueprint, &robotNextState, &robotNextContainer)
						queue = append(queue, State{robotNextState, robotNextContainer, state.minute + 1})
					}
				}
				queue = append(queue, State{nextState, state.robotContainer, state.minute + 1})
			}
		}
		qualityLevel *= currentMax
	}
	fmt.Println("Product of mined geodes number in first 3 blueprints is", qualityLevel)
}
