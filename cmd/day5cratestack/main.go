// move crates according to given plan and print out the top crates of every pile after
// it is finished.
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// global pile structure, stores the current state of the crate
// pile
var cratePile [][]string

// read in data, parse and do the simulation
func main() {
	_ = parseAndSimulate()

	fmt.Println("final crate pile:")
	printCratePile()
}

// read values from input, parse the crate pile data structure and simulate moving the crates
// return top crates in the end.
func parseAndSimulate() (err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	// Start reading from the file with a scanner.
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "[") {
			crateLine, _ := parseCrateLine(line)
			cratePile = append(cratePile, crateLine[:])
		} else if strings.Contains(line, "move") {
			// move a crate, parsing of crate pile is finished if first
			// move line appears
			crateAmount, start, target := parseMoves(line)
			moveCrates(crateAmount, start, target)
		} else {
			//fmt.Printf("skipping line: %s\n", line)
		}
	}

	return nil
}

// move the amount of crates from start to target, just one after the another.
func moveCrates(amount int, start int, target int) {
	for i := 0; i < amount; i++ {
		firstCrateIndexStart := findFirstCrateIndex(start - 1)
		crate := cratePile[firstCrateIndexStart][start-1]
		cratePile[firstCrateIndexStart][start-1] = ""
		firstCrateIndexTarget := findFirstCrateIndex(target - 1)
		if firstCrateIndexTarget == 0 {
			// we need to make room for a new line
			extendCratePile()
			cratePile[0][target-1] = crate
		} else {
			cratePile[firstCrateIndexTarget-1][target-1] = crate
		}
	}
}

// find first index of a crate
func findFirstCrateIndex(startPile int) (index int) {
	for i := 0; i < len(cratePile); i++ {
		if !isEmpty(cratePile[i][startPile]) {
			return i
		}
	}
	return len(cratePile)
}

// extend the crate pile by adding one line before the first line.
func extendCratePile() {
	newLine := make([][]string, 1)
	// initialize with the proper slice
	newLine[0] = make([]string, 10)
	newCratePile := append(newLine, cratePile...)
	cratePile = newCratePile
}

// parse and return crate line
func parseCrateLine(line string) (result [10]string, err error) {
	for i := 0; i < len(line); i += 4 {
		token := line[i : i+3]
		if !isEmpty(token) {
			result[i/4] = token
		}
	}
	return result, nil
}

// check if string is empty (either nil or just contains blanks).
func isEmpty(stringElem string) (result bool) {
	return len(strings.TrimSpace(stringElem)) == 0
}

// check if the whole slice just contains empty strings.
func isSliceEmpty(slice []string) (result bool) {
	for i := 0; i < len(slice); i++ {
		if !isEmpty(slice[i]) {
			return false
		}
	}
	return true
}

func parseMoves(line string) (crateAmount int, start int, target int) {
	r := regexp.MustCompile("move ([0-9]+) from ([0-9]+) to ([0-9]+)")
	matches := r.FindAllStringSubmatch(line, -1)
	// there is only one match
	resultTokens := matches[0]
	crateAmount, _ = strconv.Atoi(resultTokens[1])
	start, _ = strconv.Atoi(resultTokens[2])
	target, _ = strconv.Atoi(resultTokens[3])

	return crateAmount, start, target
}

// beauty print the crate pile
func printCratePile() {
	for i := 0; i < len(cratePile); i++ {
		if !isSliceEmpty(cratePile[i]) {
			for j := 0; j < len(cratePile[i]); j++ {
				if isEmpty(cratePile[i][j]) {
					fmt.Printf("    ")
				} else {
					fmt.Printf("%s ", cratePile[i][j])
				}
			}
			fmt.Println("")
		}
	}
}
