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

func checkDrawValidity(cubeDraws []string) bool {
	for _, draw := range cubeDraws {
		draw := strings.TrimSpace(draw)
		drawValues := strings.Split(draw, " ")

		amount, err := strconv.Atoi(drawValues[0])
		check(err)
		color := drawValues[1]

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

		gameIDRegex := regexp.MustCompile(`^Game ([0-9]+)`)
		gameID := gameIDRegex.FindSubmatch([]byte(line))[1]

		cubeString := strings.Split(line, ":")[1]
		cubeDrawCollections := strings.Split(cubeString, ";")
		var cubeDraws []string
		for _, collection := range cubeDrawCollections {
			cubeDraws = append(cubeDraws, strings.Split(collection, ",")...)
		}

		if checkDrawValidity(cubeDraws) {
			gameIDNum, err := strconv.Atoi(string(gameID))
			check(err)
			sumOfValidDraws += gameIDNum
		}
	}

	fmt.Print(sumOfValidDraws)
}
