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

func byteArrayToInt(input []byte) int {
	int, err := strconv.Atoi(string(input))
	check(err)
	return int
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	// remove whitespace and parse time and distance as numbers
	timeString := strings.ReplaceAll(lines[0], " ", "")
	time := byteArrayToInt(numberRegex.Find([]byte(timeString)))

	distanceString := strings.ReplaceAll(lines[1], " ", "")
	distance := byteArrayToInt(numberRegex.Find([]byte(distanceString)))

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

	fmt.Print(wins)
}
