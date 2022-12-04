package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file_name := os.Args[1]
	file, err := os.Open(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	overlapping_pairs := 0

	for scanner.Scan() {
		a, b := split_pair(scanner.Text(), ",", func(s string) string { return s })
		al, ah := split_pair(a, "-", better_atoi)
		bl, bh := split_pair(b, "-", better_atoi)
		if math.Max(al, bl) <= math.Min(ah, bh) {
			overlapping_pairs++
		}
	}

	fmt.Println(overlapping_pairs, "pairs overlap")
	file.Close()
}

func split_pair[T any](text string, separator string, convert func(string) T) (T, T) {
	splitted := strings.Split(text, separator)
	return convert(splitted[0]), convert(splitted[1])
}

func better_atoi(val string) float64 {
	num, _ := strconv.Atoi(val)
	return float64(num)
}
