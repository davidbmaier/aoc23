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

func findMaxDrawsPerColor(cubeDraws []string) map[string]int {
	maxDraws := map[string]int{
		"blue":  0,
		"red":   0,
		"green": 0,
	}

	for _, draw := range cubeDraws {
		draw := strings.TrimSpace(draw)
		drawValues := strings.Split(draw, " ")

		amount, err := strconv.Atoi(drawValues[0])
		check(err)
		color := drawValues[1]

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

		cubeString := strings.Split(line, ":")[1]
		cubeDrawCollections := strings.Split(cubeString, ";")
		var cubeDraws []string
		for _, collection := range cubeDrawCollections {
			cubeDraws = append(cubeDraws, strings.Split(collection, ",")...)
		}

		maxDraws := findMaxDrawsPerColor(cubeDraws)
		powerOfLine := maxDraws["blue"] * maxDraws["red"] * maxDraws["green"]
		sumOfPowers += powerOfLine
	}

	fmt.Print(sumOfPowers)
}
