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

	var finalDigits []string // helper variable to store all first-and-last digit combinations
	// loop through all lines of the input
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		digits := "" // helper variable to store the digits for this line (as a string because they're supposed to be concatenated)
		// go through all characters of the line and stop when we find the first digit
		for j := 0; j < len(line); j++ {
			lineChar := line[j]
			intValue, err := strconv.Atoi(string(lineChar))
			if err == nil {
				// no error when parsing string to integer, so this must be the first digit
				digits += fmt.Sprint(intValue)
				break
			}
		}

		// go through all characters of the line backwards and stop when we find the first (i.e. last) digit
		for j := len(line) - 1; j > -1; j-- {
			lineChar := line[j]
			intValue, err := strconv.Atoi(string(lineChar))
			if err == nil {
				// no error when parsing string to integer, so this must be the last digit
				digits += fmt.Sprint(intValue)
				break
			}
		}

		finalDigits = append(finalDigits, digits)
	}

	// sum up all digits
	var sumOfDigits int
	for i := 0; i < len(finalDigits); i++ {
		fmt.Print(lines[i], ": ", finalDigits[i], "\n")
		intValue, err := strconv.Atoi(finalDigits[i])
		check(err)

		sumOfDigits += intValue
	}

	fmt.Print("Sum: ", sumOfDigits)
}
