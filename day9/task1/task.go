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

var numberRegex = regexp.MustCompile(`-*\d+`) // optional - for negative values

func byteArraysIntoInts(input [][]byte) []int {
	result := []int{}
	for _, item := range input {
		int, err := strconv.Atoi(string(item))
		check(err)
		result = append(result, int)
	}
	return result
}

// returns an array containing the original array and all derived difference arrays until it's all 0
func assembleDifferenceArrays(input []int) [][]int {
	result := [][]int{}

	allZeroes := true
	differences := []int{}
	for i, value := range input {
		// ignore the last one since it has no next value to compare with
		if i != len(input)-1 {
			difference := input[i+1] - value
			differences = append(differences, difference)

			if allZeroes && difference != 0 {
				// track if this set of differences is the final one (all zeroes)
				allZeroes = false
			}
		}
	}

	result = append(result, differences)

	if !allZeroes {
		// if this isn't the final set of differences, recursively assemble the next one
		nextLevelDifferences := assembleDifferenceArrays(differences)
		result = append(result, nextLevelDifferences...)
	}
	return result
}

// recursively calculates the next value for the given index of the differenceArrays
func calculateNextValueForDifferenceArray(differenceArrays [][]int, indexToBeCalculated int) int {
	nextValueDifference := 0
	if len(differenceArrays[indexToBeCalculated+1]) != len(differenceArrays[indexToBeCalculated]) {
		// if the underlying difference array's next value isn't calculated yet, do that first
		nextValueDifference = calculateNextValueForDifferenceArray(differenceArrays, indexToBeCalculated+1)
	}

	currentDifferenceArray := differenceArrays[indexToBeCalculated]
	previousValue := currentDifferenceArray[len(currentDifferenceArray)-1]

	return previousValue + nextValueDifference
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	sum := 0
	for _, line := range lines {
		numbers := byteArraysIntoInts(numberRegex.FindAll([]byte(line), -1))

		differenceArrays := append([][]int{numbers}, assembleDifferenceArrays(numbers)...)

		// add trailing 0 to last difference array to start off the process
		differenceArrays[len(differenceArrays)-1] = append(differenceArrays[len(differenceArrays)-1], 0)

		// go through difference arrays backwards and calculate the next value for each array
		nextValue := calculateNextValueForDifferenceArray(differenceArrays, 0)
		sum += nextValue
	}

	fmt.Print(sum)
}
