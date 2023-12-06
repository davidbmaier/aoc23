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

var numberRegex = regexp.MustCompile(`\d+`)

func byteArraysToInts(inputs [][]byte) []int {
	ints := []int{}
	for _, input := range inputs {
		int, err := strconv.Atoi(string(input))
		check(err)
		ints = append(ints, int)
	}
	return ints
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	// get int arrays for times and distances
	times := byteArraysToInts(numberRegex.FindAll([]byte(lines[0]), -1))
	distances := byteArraysToInts(numberRegex.FindAll([]byte(lines[1]), -1))

	winProduct := 1
	for i, time := range times {
		distance := distances[i]

		// find the first time value that wins the race
		firstWin := -1
		for j := 0; j < time; j++ {
			if (time-j)*j > distance {
				firstWin = j
				break
			}
		}

		// using that first win, subtract the earlier losses (and the later losses at the other end of the times)
		wins := time + 1 - 2*(firstWin)
		winProduct *= wins
	}

	fmt.Print(winProduct)
}
