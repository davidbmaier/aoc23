package main

import (
	"fmt"
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

func updateDuplicateMap(winCount int, cardIndex int, duplicateMap map[int]int) map[int]int {
	// iterate from 1 up to (including) winCount to identify the cards we need to duplicate later
	for j := 1; j <= winCount; j++ {
		duplicateMap[cardIndex+j]++
	}
	return duplicateMap
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	// map to track how many additional times a card has to be processed
	duplicateMap := map[int]int{}

	numberOfProcessedCards := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// split the data into smaller strings to work with further
		cardString := strings.Split(line, ":")[1]
		winnersString := strings.Split(cardString, "|")[0]
		picksString := strings.Split(cardString, "|")[1]

		// parse numbers from each section of the card
		winners := numberRegex.FindAll([]byte(winnersString), -1)
		picks := numberRegex.FindAll([]byte(picksString), -1)

		// figure out how many picks are winners
		winCount := 0
		for _, winner := range winners {
			for _, pick := range picks {
				if string(winner) == string(pick) {
					winCount++
				}
			}
		}

		// update the duplicateMap for the original processing of this card
		duplicateMap = updateDuplicateMap(winCount, i, duplicateMap)
		numberOfProcessedCards++

		// update the duplicateMap for the duplicates of this card
		if duplicateAmount, ok := duplicateMap[i]; ok {
			for j := 1; j <= duplicateAmount; j++ {
				duplicateMap = updateDuplicateMap(winCount, i, duplicateMap)
				numberOfProcessedCards++
			}
		}

		fmt.Print("Done with card ", i+1, ", processed cards: ", numberOfProcessedCards, "\n")
	}

	fmt.Print(numberOfProcessedCards)
}
