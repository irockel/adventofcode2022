// reads input file input.txt and find the elf carrying most
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

//
// read in data and search and calculate result
//
func main() {
	values, _ := findMaxSum()

	fmt.Printf("result is %d\n", values[0] + values[1] + values[2])
}

//
// read values from input and the max sum found in input
//
func findMaxSum() (maxSums [3]int, err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	result := [3]int {0, 0, 0}
	currentVal := 0

	// Start reading from the file with a reader.
	reader := bufio.NewReader(file)
	var line string
	for {
		line, err = reader.ReadString('\n')
		if err != nil && err != io.EOF {
			break
		}
		if strings.TrimSpace(line) != "" {
			val, err := strconv.Atoi(strings.TrimSpace(line))
			if err== nil {
				currentVal += val
			}
		} else {
			slot := -1
			currentDiff := 0
			for i := 0; i < 3; i++ {
				if currentDiff < (currentVal - result[i]) {
					slot = i
					currentDiff = currentVal - result[i]
				}
			}
			if slot >= 0 {
				result[slot] = currentVal
			}
			currentVal = 0
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return [3]int {0, 0, 0}, err
	}
	return result, nil
}
