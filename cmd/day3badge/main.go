// identify and calculate security badges of group of three backpacks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// read in data and search and calculate result
func main() {
	priorities, _ := processBags()

	fmt.Printf("priority sum of all backpacks security badges is %d\n", priorities)
}

// read values from input, identify the badges and calculate the priorities
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

	for scanner.Scan() {
		// always scan three lines, requires that the count of lines
		// can be divided by three, which the input data fulfills
		lines := [3]string{scanner.Text(), "", ""}
		scanner.Scan()
		lines[1] = scanner.Text()
		scanner.Scan()
		lines[2] = scanner.Text()

		bagPriorities, err := scanBags(lines)

		if err != nil {
			return 0, nil
		}
		result += bagPriorities
	}

	return result, nil
}

func scanBags(lines [3]string) (result int, err error) {
	result = 0

	for _, charAt := range lines[0] {
		if strings.Contains(lines[1], string(charAt)) &&
			strings.Contains(lines[2], string(charAt)) {
			if int(charAt) > 96 {
				// small letters
				result = int(charAt) - 96
			} else {
				// capital letters
				result = int(charAt) - 38
			}

			// it is defined, there's only exact one entry shared in the three bags
			// so we can stop here
			break
		}
	}

	return result, nil
}
