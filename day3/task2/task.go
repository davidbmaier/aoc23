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

var numberRegex = regexp.MustCompile(`[0-9]{1}`) // simple regex to match individual digits

// helper function to get the full number from any known index inside a string
func getNumberFromPositionInLine(line string, index int) int {
	// initial value of the string is the provided index's value - we know this is part of the number
	numberString := string(line[index])
	// go backwards in the line until the start of the number is reached
	for i := index; i >= 0; i-- {
		if !numberRegex.Match([]byte{line[i]}) {
			// not a digit, so we've reached the end
			break
		}
		// ignore the originally provided index, that was already added in the beginning
		if i != index {
			numberString = string(line[i]) + numberString
		}
	}
	// go forwards in the line until the end of the number is reached
	for i := index; i < len(line); i++ {
		if !numberRegex.Match([]byte{line[i]}) {
			// not a digit, so we've reached the end
			break
		}
		// ignore the originally provided index, that was already added in the beginning
		if i != index {
			numberString = numberString + string(line[i])
		}
	}

	number, err := strconv.Atoi(numberString)
	check(err)
	return number
}

// helper function to find all numbers in the line that are adjacent to the given index
func findNumbersInPartialLine(line string, index int) []int {
	// regex for identifying that there's exactly one continuous number in a three-character string
	oneNumberRegex := regexp.MustCompile(`^[^0-9\s]*[0-9]+[^0-9\s]*$`)
	// regex for identifying that there's exactly two continuous numbers in a three-character string
	twoNumbersRegex := regexp.MustCompile(`^[0-9]{1}[^0-9]{1}[0-9]{1}$`)

	startIndex := index - 1
	endIndex := index + 1

	// make sure start and end don't exceed the line
	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(line) {
		endIndex = len(line)
	}

	// get a three-character substring around the starting index to check if there are any numbers in those three positions
	partialLine := line[startIndex : endIndex+1]

	if oneNumberRegex.Match([]byte(partialLine)) {
		// exactly one number found, find the index of a number character to start with (since it could be in any position)
		// also add the startIndex since the found index is relative to the substring, not the full line
		numberIndex := numberRegex.FindIndex([]byte(partialLine))[0] + startIndex
		return []int{
			getNumberFromPositionInLine(line, numberIndex),
		}
	}
	if twoNumbersRegex.Match([]byte(partialLine)) {
		// exactly two numbers found, they have to be in startIndex and endIndex (because the middle one is empty)
		return []int{
			getNumberFromPositionInLine(line, startIndex),
			getNumberFromPositionInLine(line, endIndex),
		}
	}
	// no numbers found, return an empty array
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

			// collect full adjacent numbers in an array for later
			adjacentNumbers := []int{}

			// go through adjacent characters and find numbers
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

			// if there's exactly two numbers, add their product to the final sum
			if len(adjacentNumbers) == 2 {
				sumOfGearRatios += adjacentNumbers[0] * adjacentNumbers[1]
			}
		}
	}
	fmt.Print(sumOfGearRatios)
}
