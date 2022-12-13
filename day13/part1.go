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

	pairStrings := strings.Split(string(file), "\n\n")
	indexSum := 0

	for index, pairString := range pairStrings {
		lines := strings.Split(pairString, "\n")
		var upperLine, lowerLine []any
		json.Unmarshal([]byte(lines[0]), &upperLine)
		json.Unmarshal([]byte(lines[1]), &lowerLine)
		rightOrder, _ := isRightOrder(upperLine, lowerLine)
		if rightOrder {
			indexSum += index + 1
		}
	}

	fmt.Println("Sum of the indices is", indexSum)

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
