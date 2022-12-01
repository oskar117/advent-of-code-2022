package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
    file_name := os.Args[1]
    file, err := os.Open(file_name)

    if err != nil {
        fmt.Println("Error opening file!")
    }
    
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanLines)

    index, largest_index := 0, 0
    temp_value, largest_value := 0, 0

    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            if temp_value > largest_value {
                largest_index = index
                largest_value = temp_value
            }
            index++
            temp_value = 0
        } else {
            textAsNumber, _ := strconv.Atoi(text) 
            temp_value += textAsNumber
        }
    }

    fmt.Println("Elf", largest_index + 1, "is carrying the most calories:", largest_value)
    file.Close()
}

