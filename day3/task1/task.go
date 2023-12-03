package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// simple regex to look for a character that's "special", i.e. not a number, not a "." and not whitespace (e.g. like a line break)
var specialSymbolRegex = regexp.MustCompile(`[^0-9\.\s]{1}`)

func checkPartialLineForSpecialSymbol(line string, offset int, length int) bool {
	// limit offset and length so we don't exceed the line
	if offset < 0 {
		offset = 0
	}
	if offset+length > len(line) {
		length = len(line) - offset
	}

	// get a substring of the full line by using offset and length
	partialLine := line[offset:(offset + length)]
	// check if there's a special symbol in the partial line and return that check's result
	return specialSymbolRegex.Match([]byte(partialLine))
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	sumOfPartNumbers := 0
	for lineIndex, line := range lines {
		// go through all strings of digits
		numberRegex := regexp.MustCompile(`[0-9]+`)
		// use FindAllIndex to be able to identify the exact match (since finding it by its value again later can lead to mistakes)
		numberPositions := numberRegex.FindAllIndex([]byte(line), -1)

		// go through each of the found index arrays
		for _, numberPosition := range numberPositions {
			// numberPosition has two items: the starting index and the ending index of the match
			numberIndex := numberPosition[0]
			numberString := line[numberPosition[0]:numberPosition[1]]
			numberLength := len(numberString)

			partNumberConfirmed := false // helper variable to track if we need to continue checking

			// go through adjacent characters and check if they are special symbols
			// always extend the check to one character to the left and right (starting point - 1, length + 2)
			if lineIndex > 0 {
				// check the line above
				lineAbove := lines[lineIndex-1]
				partNumberConfirmed = checkPartialLineForSpecialSymbol(lineAbove, numberIndex-1, numberLength+2)
			}
			if !partNumberConfirmed {
				// check the current line - this also happens to include the actual number, but that doesn't hurt the check
				partNumberConfirmed = checkPartialLineForSpecialSymbol(line, numberIndex-1, numberLength+2)
			}
			if lineIndex < len(lines)-1 && !partNumberConfirmed {
				// check the next line
				lineBelow := lines[lineIndex+1]
				partNumberConfirmed = checkPartialLineForSpecialSymbol(lineBelow, numberIndex-1, numberLength+2)
			}

			// if yes, add it to the sum of part numbers
			if partNumberConfirmed {
				numberInt, err := strconv.Atoi(numberString)
				check(err)
				sumOfPartNumbers += numberInt
				fmt.Print("✅: ", numberString)
			} else {
				fmt.Print("❌: ", numberString)
			}
			fmt.Print("\t", sumOfPartNumbers, "\n")
		}
	}

	fmt.Print(sumOfPartNumbers)
}
