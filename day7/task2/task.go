package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var cards = map[string]int{
	"J":  0,
	"2":  1,
	"3":  2,
	"4":  3,
	"5":  4,
	"6":  5,
	"7":  6,
	"8":  7,
	"9":  8,
	"10": 9,
	"T":  10,
	"Q":  11,
	"K":  12,
	"A":  13,
}

const (
	highCard     = 0
	onePair      = 1
	twoPair      = 2
	threeOfAKind = 3
	fullHouse    = 4
	fourOfAKind  = 5
	fiveOfAKind  = 6
)

type Hand struct {
	cardString string
	cardValues []int
	handType   int
	bid        int
}

func calculateHandType(hand Hand) int {
	handValues := hand.cardValues
	// probably not the smartest way but it works fine
	pair1 := 0
	pair2 := 0
	pair1Value := -1
	jokerCounter := 0

	sortedHandValues := make([]int, len(handValues))
	copy(sortedHandValues, handValues)

	sort.Slice(sortedHandValues, func(i, j int) bool {
		return sortedHandValues[i] < sortedHandValues[j]
	})

	for i := 0; i < len(sortedHandValues); i++ {
		if sortedHandValues[i] == 0 {
			jokerCounter++
			continue
		}
		if i < len(sortedHandValues)-1 && sortedHandValues[i] == sortedHandValues[i+1] {
			if pair1 > 0 && sortedHandValues[i] != pair1Value {
				pair2++
			} else {
				pair1++
				pair1Value = sortedHandValues[i]
			}
		}
	}

	if jokerCounter == 5 {
		// special case: all jokers
		pair1 = 4
	} else if pair1 > pair2 || pair1 == pair2 {
		pair1 += jokerCounter
	} else {
		pair2 += jokerCounter
	}

	returnValue := 0
	if pair1 == 4 {
		returnValue = fiveOfAKind
	} else if pair1 == 3 {
		returnValue = fourOfAKind
	} else if (pair1 == 2 && pair2 == 1) || (pair1 == 1 && pair2 == 2) {
		returnValue = fullHouse
	} else if pair1 == 2 || pair2 == 2 {
		returnValue = threeOfAKind
	} else if pair1 == 1 {
		if pair2 == 1 {
			returnValue = twoPair
		} else {
			returnValue = onePair
		}
	} else {
		returnValue = highCard
	}

	return returnValue
}

func transformStringToHand(handString string) Hand {
	hand := Hand{
		cardValues: []int{},
		bid:        0,
	}

	handParts := strings.Split(handString, " ")
	hand.cardString = handParts[0]

	// assign the values
	for i := 0; i < len(handParts[0]); i++ {
		character := string(handParts[0][i])
		hand.cardValues = append(hand.cardValues, cards[character])
	}

	// assign the bid
	bid, err := strconv.Atoi(handParts[1])
	check(err)
	hand.bid = bid

	// calculate the type of hand
	hand.handType = calculateHandType(hand)

	return hand
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	hands := []Hand{}
	for _, line := range lines {
		hands = append(hands, transformStringToHand(line))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			// not equal hand types so sort by that
			return hands[i].handType < hands[j].handType
		}
		// equal hand types so sort by first card difference
		for k := 0; k < len(hands[i].cardValues); k++ {
			if hands[i].cardValues[k] != hands[j].cardValues[k] {
				return hands[i].cardValues[k] < hands[j].cardValues[k]
			}
		}
		// fallback in case none of the cases happen
		return true
	})

	bidSum := 0
	for i, hand := range hands {
		bidSum += hand.bid * (i + 1)
	}

	fmt.Print(bidSum)
}
