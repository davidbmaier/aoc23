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

var numberRegex = regexp.MustCompile(`[0-9]{1}`)

func getNumberFromPositionInLine(line string, index int) int {
	numberString := string(line[index])
	// go backwards in the line until the start of the number is reached
	for i := index; i >= 0; i-- {
		if !numberRegex.Match([]byte{line[i]}) {
			break
		}
		if i != index {
			numberString = string(line[i]) + numberString
		}
	}
	// go forwards in the line until the end of the number is reached
	for i := index; i < len(line); i++ {
		if !numberRegex.Match([]byte{line[i]}) {
			break
		}
		if i != index {
			numberString = numberString + string(line[i])
		}
	}

	number, err := strconv.Atoi(numberString)
	check(err)
	return number
}

func findNumbersInPartialLine(line string, index int) []int {
	oneNumberRegex := regexp.MustCompile(`^[^0-9\s]*[0-9]+[^0-9\s]*$`)
	twoNumbersRegex := regexp.MustCompile(`^[0-9]{1}[^0-9]{1}[0-9]{1}$`)

	startIndex := index - 1
	endIndex := index + 1

	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(line) {
		endIndex = len(line)
	}

	partialLine := line[startIndex : endIndex+1]

	if oneNumberRegex.Match([]byte(partialLine)) {
		// find the index of a number character to start with
		numberIndex := numberRegex.FindIndex([]byte(partialLine))[0] + startIndex
		return []int{
			getNumberFromPositionInLine(line, numberIndex),
		}
	}
	if twoNumbersRegex.Match([]byte(partialLine)) {
		return []int{
			getNumberFromPositionInLine(line, startIndex),
			getNumberFromPositionInLine(line, endIndex),
		}
	}
	return []int{}
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	sumOfGearRatios := 0
	for lineIndex, line := range lines {
		// find all * characters
		starRegex := regexp.MustCompile(`\*`)
		// use FindAllIndex to be able to identify the exact match (since finding it by its value again later can lead to mistakes)
		starPositions := starRegex.FindAllIndex([]byte(line), -1)

		for _, starPosition := range starPositions {
			starIndex := starPosition[0]
			check(err)

			adjacentNumbers := []int{}

			// go through adjacent characters and check how many parts there are
			if lineIndex > 0 {
				// check the line above
				lineAbove := lines[lineIndex-1]
				adjacentNumbers = append(adjacentNumbers, findNumbersInPartialLine(lineAbove, starIndex)...)
			}
			// check the current line
			adjacentNumbers = append(adjacentNumbers, findNumbersInPartialLine(line, starIndex)...)
			if lineIndex < len(lines)-1 {
				// check the next line
				lineBelow := lines[lineIndex+1]
				adjacentNumbers = append(adjacentNumbers, findNumbersInPartialLine(lineBelow, starIndex)...)
			}

			if len(adjacentNumbers) == 2 {
				sumOfGearRatios += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}
	fmt.Print(sumOfGearRatios)
}
