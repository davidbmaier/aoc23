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

var specialSymbolRegex = regexp.MustCompile(`[^0-9\.\s]{1}`)

func checkPartialLineForSpecialSymbol(line string, offset int, length int) bool {
	// limit offset and length so we don't exceed the line
	if offset < 0 {
		offset = 0
	}
	if offset+length > len(line) {
		length = len(line) - offset
	}

	partialLine := line[offset:(offset + length)]

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

		for _, numberPosition := range numberPositions {
			numberIndex := numberPosition[0]
			numberString := line[numberPosition[0]:numberPosition[1]]
			check(err)

			numberLength := len(numberString)

			partNumberConfirmed := false

			// go through adjacent characters and check if they are special symbols
			if lineIndex > 0 {
				// check the line above
				lineAbove := lines[lineIndex-1]
				partNumberConfirmed = checkPartialLineForSpecialSymbol(lineAbove, numberIndex-1, numberLength+2)
			}
			if !partNumberConfirmed {
				// check the current line
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
