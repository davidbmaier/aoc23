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

// helper function to collect the maximum draws for each color per game
// maximum draws done per color = minimum amount of cubes required per color for the game to be valid
func findMaxDrawsPerColor(cubeDraws []string) map[string]int {
	// use a map to store the color/amount mappings
	maxDraws := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, draw := range cubeDraws {
		// same cleanup/splitting as in task 1
		draw := strings.TrimSpace(draw)
		drawValues := strings.Split(draw, " ")

		amount, err := strconv.Atoi(drawValues[0])
		check(err)
		color := drawValues[1]

		// update maxDraws if necessary
		if amount > maxDraws[color] {
			maxDraws[color] = amount
		}
	}

	return maxDraws
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	sumOfPowers := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// same splitting logic as in task 1
		cubeString := strings.Split(line, ":")[1]
		cubeDrawCollections := strings.Split(cubeString, ";")
		var cubeDraws []string
		for _, collection := range cubeDrawCollections {
			cubeDraws = append(cubeDraws, strings.Split(collection, ",")...)
		}

		maxDraws := findMaxDrawsPerColor(cubeDraws)
		// calculate the power of each game's draws and add it to the final sum
		powerOfLine := maxDraws["blue"] * maxDraws["red"] * maxDraws["green"]
		sumOfPowers += powerOfLine
	}

	fmt.Print(sumOfPowers)
}
