// simulate rock paper scissors strategy input.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//
// read in data and search and calculate result
//
func main() {
	score, _ := calculateStrategy()

	fmt.Printf("score of rock paper scissors strategy is %d\n", score)
}

//
// read values from input and calculate the result
//
func calculateStrategy() (result int, err error) {
	fmt.Println("reading input.txt")

	file, err := os.Open("./input.txt")
	if err != nil {
		fmt.Printf(" > Failed opening file with error: %v\n", err)
		return
	}
	defer file.Close()

	result = 0

	// Start reading from the file with a reader.
	// A for Rock, B for Paper, and C for Scissors
	// X for Rock, Y for Paper, and Z for Scissors
	// Score:
	// 1 for Rock, 2 for Paper, and 3 for Scissors)
	// plus the score for the outcome of the round (0 if you lost, 3 if the round was a draw, and 6 if you won
	// Rock defeats Scissors, Scissors defeats Paper, and Paper defeats Rock.
	// X means you need to lose, Y means you need to end the round in a draw, and Z means you need to win.
	scanner := bufio.NewScanner(file)
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		switch line {
		case "A X" :
			result += 3 + 0
		case "B X" :
			result += 1 + 0
		case "C X" :
			result += 2 + 0
		case "A Y" :
			result += 1 + 3
		case "B Y" :
			result += 2 + 3
		case "C Y" :
			result += 3 + 3
		case "A Z" :
			result += 2 + 6
		case "B Z" :
			result += 3 + 6
		case "C Z" :
			result += 1 + 6
		}

	}

	return result, nil
}
