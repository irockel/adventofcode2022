// identify and calculate security badges of group of three backpacks
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// read in data and search and calculate result
func main() {
	includingRangeCount, _ := findIncludingRanges()

	fmt.Printf("amount of including ranges is %d\n", includingRangeCount)
}

// read values from input, parse the ranges and identify the segments which are fully contained in the others
func findIncludingRanges() (result int, err error) {
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
		line := strings.TrimSpace(scanner.Text())
		if len(line) > 0 {
			rangeTokens := strings.Split(line, ",")
			// as defined, these are always two tokens.
			firstPairToken := strings.Split(rangeTokens[0], "-")
			secondPairToken := strings.Split(rangeTokens[1], "-")

			pairs := [4]int{}

			pairs[0], _ = strconv.Atoi(firstPairToken[0])
			pairs[1], _ = strconv.Atoi(firstPairToken[1])
			pairs[2], _ = strconv.Atoi(secondPairToken[0])
			pairs[3], _ = strconv.Atoi(secondPairToken[1])

			if (pairs[0] >= pairs[2] && pairs[1] <= pairs[3]) ||
				(pairs[2] >= pairs[0] && pairs[3] <= pairs[1]) {
				result++
			}

		}
		if err != nil {
			return 0, nil
		}
	}

	return result, nil
}
