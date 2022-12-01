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
	value, _ := findMaxSum()

	fmt.Printf("result is %d\n", value)
}

//
// read values from input and the max sum found in input
//
func findMaxSum() (maxSum int, err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	result := 0
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
			if currentVal > result {
				result = currentVal
			}
			currentVal = 0
		}

		if err != nil {
			break
		}
	}
	if err != io.EOF {
		fmt.Printf(" > Failed with error: %v\n", err)
		return 0, err
	}
	return result, nil
}
