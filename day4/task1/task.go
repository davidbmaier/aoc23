package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var numberRegex = regexp.MustCompile(`\d+`)

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	sumOfPoints := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// split the data into smaller strings to work with further
		cardString := strings.Split(line, ":")[1]
		winnersString := strings.Split(cardString, "|")[0]
		picksString := strings.Split(cardString, "|")[1]

		// parse numbers from each section of the card
		winners := numberRegex.FindAll([]byte(winnersString), -1)
		picks := numberRegex.FindAll([]byte(picksString), -1)

		winCount := 0
		winPoints := 0

		for _, winner := range winners {
			for _, pick := range picks {
				if string(winner) == string(pick) {
					winCount++
				}
			}
		}

		if winCount > 0 {
			winPoints = int(math.Pow(2, float64(winCount-1)))
		}

		sumOfPoints += winPoints
	}

	fmt.Print(sumOfPoints)
}
