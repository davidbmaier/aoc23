package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	// map for all strings that we'll be looking for (along with their associated int value)
	stringMap := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	var lineValues []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// initialize lowest and highest index values (with their minimum possible values)
		lowestIndex := len(line) - 1
		highestIndex := 0
		var lowestValue, highestValue int // declare lowest and highest values for separate tracking

		// go through all strings we're looking for
		for stringKey, stringValue := range stringMap {
			firstIndex := strings.Index(line, stringKey)    // index of the first occurrence
			lastIndex := strings.LastIndex(line, stringKey) // index of the last occurrence

			// bail out if the string isn't contained at all
			if firstIndex == -1 {
				continue
			}

			// update lowest and highest index/value variables if necessary
			if firstIndex <= lowestIndex {
				lowestIndex = firstIndex
				lowestValue = stringValue
			}
			if lastIndex >= highestIndex {
				highestIndex = lastIndex
				highestValue = stringValue
			}
		}

		lineValue := fmt.Sprint(lowestValue) + fmt.Sprint(highestValue)
		fmt.Print(line, ": ", lineValue, "\n")
		lineValues = append(lineValues, lineValue)
	}

	// add up all the lines' values
	sumOfLines := 0
	for i := 0; i < len(lineValues); i++ {
		intValue, err := strconv.Atoi(lineValues[i])
		check(err)
		sumOfLines += intValue
	}
	fmt.Print(sumOfLines)
}
