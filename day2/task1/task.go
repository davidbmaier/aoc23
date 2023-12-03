package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var maxCubes = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// helper function to check if a given draw is possible
func checkDrawValidity(cubeDraws []string) bool {
	for _, draw := range cubeDraws {
		draw := strings.TrimSpace(draw)        // remove trailing whitespace (to get rid of extra spaces behind commas/semicolons)
		drawValues := strings.Split(draw, " ") // format is always "number color", so we can get the individual values by splitting the string

		amount, err := strconv.Atoi(drawValues[0])
		check(err)
		color := drawValues[1]

		// if one amount is higher than its color's max, we can bail out and return early
		if amount > maxCubes[color] {
			return false
		}
	}
	return true
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	sumOfValidDraws := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// extract the game number with a simple regex
		// (no lookarounds in Go's default regex library, so use a capturing group to get the number separately)
		gameIDRegex := regexp.MustCompile(`^Game ([0-9]+)`)
		// FindSubmatch is required to return capturing groups - first item is the full match, the following items are the capturing groups
		gameID := gameIDRegex.FindSubmatch([]byte(line))[1]

		cubeString := strings.Split(line, ":")[1]             // format is always "Game number: ...", so we can get the draws by splitting
		cubeDrawCollections := strings.Split(cubeString, ";") // same thing here, just split the individual draws
		var cubeDraws []string
		for _, collection := range cubeDrawCollections {
			// collect all the draw groups in an array - this isn't really needed because you can look at each draw separately
			cubeDraws = append(cubeDraws, strings.Split(collection, ",")...)
		}

		if checkDrawValidity(cubeDraws) {
			// if valid, add the number to the final sum
			gameIDNum, err := strconv.Atoi(string(gameID))
			check(err)
			sumOfValidDraws += gameIDNum
		}
	}

	fmt.Print(sumOfValidDraws)
}
