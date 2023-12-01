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

		lowestIndex := len(line) - 1
		highestIndex := 0
		var lowestValue, highestValue int

		for stringKey, stringValue := range stringMap {
			firstIndex := strings.Index(line, stringKey)
			lastIndex := strings.LastIndex(line, stringKey)

			if firstIndex == -1 {
				continue
			}

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

	sumOfLines := 0
	for i := 0; i < len(lineValues); i++ {
		intValue, err := strconv.Atoi(lineValues[i])
		check(err)
		sumOfLines += intValue
	}
	fmt.Print(sumOfLines)
}
