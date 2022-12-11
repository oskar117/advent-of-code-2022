package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type monkey struct {
	items             []int
	operation         func(old int) int
	divisor           int
	trueMonkeyIndex   int
	falseMonkeyIndex  int
	inspectionCounter int
}

func (monkey *monkey) processNextItem() (int, int) {
	monkey.inspectionCounter++
	item := monkey.items[0]
	monkey.items = monkey.items[1:]
	item = monkey.operation(item)
	item /= 3
	if item%monkey.divisor == 0 {
		return item, monkey.trueMonkeyIndex
	} else {
		return item, monkey.falseMonkeyIndex
	}
}

func (monkey *monkey) hasItems() bool {
	return len(monkey.items) > 0
}

func (monkey *monkey) catchItem(item int) {
	monkey.items = append(monkey.items, item)
}

func Monkey(input string) *monkey {
	lines := strings.Split(input, "\n")
	items := make([]int, 0)
	var left, operator, right string
	var divisor, trueMonkeyIndex, falseMonkeyIndex int
	itemsStringArray := strings.Split(strings.TrimSpace(lines[1])[16:], ", ")
	for _, item := range itemsStringArray {
		value, _ := strconv.Atoi(item)
		items = append(items, value)
	}
	fmt.Sscanf(strings.TrimSpace(lines[2]), "Operation: new = %s %s %s", &left, &operator, &right)
	fmt.Sscanf(strings.TrimSpace(lines[3]), "Test: divisible by %d", &divisor)
	fmt.Sscanf(strings.TrimSpace(lines[4]), "If true: throw to monkey %d", &trueMonkeyIndex)
	fmt.Sscanf(strings.TrimSpace(lines[5]), "If false: throw to monkey %d", &falseMonkeyIndex)
	return &monkey{items, getOperationFunction(left, operator, right), divisor, trueMonkeyIndex, falseMonkeyIndex, 0}
}

func getOperationFunction(left, operator, right string) func(old int) int {
	return func(old int) int {
		var num1, num2 int
		if left == "old" {
			num1 = old
		} else {
			num, _ := strconv.Atoi(left)
			num1 = num
		}
		if right == "old" {
			num2 = old
		} else {
			num, _ := strconv.Atoi(right)
			num2 = num
		}
		switch operator {
		case "*":
			return num1 * num2
		case "/":
			return num1 / num2
		case "+":
			return num1 + num2
		case "-":
			return num1 - num2
		default:
			return num1
		}
	}
}

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	monkey_input := strings.Split(string(file), "\n\n")
	monkeys := make([]*monkey, len(monkey_input))

	for i := 0; i < len(monkey_input); i++ {
		monkeys[i] = Monkey(monkey_input[i])
	}

	turns := 20

	for turn := 0; turn < turns; turn++ {
		for _, monkey := range monkeys {
			for monkey.hasItems() {
				item, nextMonkeyIndex := monkey.processNextItem()
				monkeys[nextMonkeyIndex].catchItem(item)
			}
		}
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].inspectionCounter > monkeys[j].inspectionCounter
	})

	fmt.Println("Monkey business of top2 monkeys is", monkeys[0].inspectionCounter*monkeys[1].inspectionCounter)
}
