// try to find the start message marker from the given input string (only one line is read)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// read in data and search for start packet
func main() {
	startPacket, index, _ := findStartPacket()

	fmt.Printf("the start packet is %s and %d characters were processed\n", startPacket, index)
}

// read string line from input and parse character by character
// BEWARE: only operates properly on ASCII strings, UTF-8 needs a
// different iterating approach
func findStartPacket() (result string, index int, err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	startPacket := ""

	// Start reading from the file with a scanner.
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		dataLine := strings.TrimSpace(scanner.Text())
		for i := 14; i < len(dataLine); i++ {
			currentBlock := dataLine[i-14 : i]
			if hasNoDouble(currentBlock) {
				return currentBlock, i, nil
			}
		}
	}

	return startPacket, 0, nil
}

func hasNoDouble(currentBlock string) (result bool) {
	for i := 0; i < len(currentBlock); i++ {
		for j := 0; j < len(currentBlock); j++ {
			if (i != j) && (currentBlock[i] == currentBlock[j]) {
				return false
			}
		}
	}

	return true
}
