package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"reflect"
	"strings"
)

func main() {
	file_name := os.Args[1]
	file, err := os.ReadFile(file_name)

	if err != nil {
		fmt.Println("Error opening file!")
	}

	first_divider := "[[2]]"
	second_divider := "[[6]]"

	lines := strings.Split(strings.ReplaceAll(string(file), "\n\n", "\n"), "\n")
	lines = append(lines, first_divider)
	lines = append(lines, second_divider)

	var anyLines [][]any

	for _, line := range lines {
		var converted []any
		json.Unmarshal([]byte(line), &converted)
		anyLines = append(anyLines, converted)
	}

	for i := 0; i < len(anyLines); i++ {
		currMin := i
		for j := 0 + i; j < len(anyLines); j++ {
			isSorted, _ := isRightOrder(anyLines[j], anyLines[currMin])
			if isSorted {
				currMin = j
			}
		}
		if currMin != i {
			temp := anyLines[i]
			anyLines[i] = anyLines[currMin]
			anyLines[currMin] = temp
		}
	}

	decoderKey := 1

	for index, line := range anyLines {
		lineString := fmt.Sprintf("%v", line)
		if lineString == first_divider || lineString == second_divider {
			decoderKey *= index
		}
	}

	fmt.Println("Decoder key for the distress signal is", decoderKey)

}

func isRightOrder(upperLine, lowerLine []any) (bool, bool) {
	length := int(math.Max(float64(len(upperLine)), float64(len(lowerLine))))
	for i := 0; i < length; i++ {
		if i > len(upperLine)-1 {
			return true, false
		}
		if i > len(lowerLine)-1 {
			return false, false
		}
		if reflect.TypeOf(upperLine[i]).String() == "float64" && reflect.TypeOf(lowerLine[i]).String() == "float64" {
			left := upperLine[i].(float64)
			right := lowerLine[i].(float64)
			if left > right {
				return false, false
			}
			if right > left {
				return true, false
			}
		} else if reflect.TypeOf(upperLine[i]).String() == "[]interface {}" && reflect.TypeOf(lowerLine[i]).String() == "[]interface {}" {
			res, cont := isRightOrder(upperLine[i].([]any), lowerLine[i].([]any))
			if !cont {
				return res, false
			}
		} else if reflect.TypeOf(upperLine[i]).String() == "float64" {
			res, cont := isRightOrder([]any{upperLine[i]}, lowerLine[i].([]any))
			if !cont {
				return res, false
			}
		} else if reflect.TypeOf(lowerLine[i]).String() == "float64" {
			res, cont := isRightOrder(upperLine[i].([]any), []any{lowerLine[i]})
			if !cont {
				return res, false
			}
		}
	}
	return false, true
}
