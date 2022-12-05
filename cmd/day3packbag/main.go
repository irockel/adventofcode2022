// identify and calculate priorities of double entries
package main

import (
	"bufio"
	"fmt"
	"os"
)

// read in data and search and calculate result
func main() {
	priorities, _ := processBags()

	fmt.Printf("priority sum of all backpacks is %d\n", priorities)
}

// read values from input, identify the double entries and calculate the priorities
func processBags() (result int, err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	result = 0

	// Start reading from the file with a scanner.
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		bagPriorities, err := scanBag(line)

		if err != nil {
			return 0, nil
		}
		result += bagPriorities
	}

	return result, nil
}

// scan the bag content, find the double entries and calculate the priorities
func scanBag(line string) (result int, err error) {
	result = 0
	halfBucket := make(map[int]bool)
	doubleEntries := make(map[int]bool)

	for i, charAt := range line {
		if i < (len(line) / 2) {
			halfBucket[int(charAt)] = true
		} else {
			if halfBucket[int(charAt)] && !doubleEntries[int(charAt)] {
				// found double entries, calculate priority
				doubleEntries[int(charAt)] = true
				if int(charAt) > 96 {
					// small letters
					result += int(charAt) - 96
				} else {
					// capital letters
					result += int(charAt) - 38
				}
			}
		}
	}

	return result, nil
}
