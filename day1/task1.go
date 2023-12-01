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
	data, err := os.ReadFile("./data1.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	var finalDigits []string
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		digits := ""
		// find first digit
		for j := 0; j < len(line); j++ {
			lineChar := line[j]
			intValue, err := strconv.Atoi(string(lineChar))
			if err == nil {
				digits += fmt.Sprint(intValue)
				break
			}
		}

		// find last digit
		for j := len(line) - 1; j > -1; j-- {
			lineChar := line[j]
			intValue, err := strconv.Atoi(string(lineChar))
			if err == nil {
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
