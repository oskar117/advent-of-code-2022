package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

    var elves []int
    index := 0
    temp_value := 0

    for scanner.Scan() {
        text := scanner.Text()
        if text == "" {
            elves = append(elves, temp_value)
            index++
            temp_value = 0
        } else {
            textAsNumber, _ := strconv.Atoi(text) 
            temp_value += textAsNumber
        }
    }

    if temp_value != 0 {
        elves = append(elves, temp_value)
    }

    sort.Sort(sort.Reverse(sort.IntSlice(elves)))
    top_elves := elves[:3] 

    fmt.Println("Top 3 elves are carrying", sum(top_elves), "calories")
    file.Close()
}

func sum(numbers []int) int {
    result := 0
    for _, number := range numbers {
        result += number
    }
    return result
}

