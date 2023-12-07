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
	"2":  0,
	"3":  1,
	"4":  2,
	"5":  3,
	"6":  4,
	"7":  5,
	"8":  6,
	"9":  7,
	"10": 8,
	"T":  9,
	"J":  10,
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

func calculateHandType(handValues []int) int {
	// probably not the smartest way but it works fine
	pair1 := 0
	pair2 := 0
	pair1Value := -1

	sortedHandValues := make([]int, len(handValues))
	copy(sortedHandValues, handValues)

	sort.Slice(sortedHandValues, func(i, j int) bool {
		return sortedHandValues[i] < sortedHandValues[j]
	})

	for i := 0; i < len(sortedHandValues); i++ {
		if i < len(sortedHandValues)-1 && sortedHandValues[i] == sortedHandValues[i+1] {
			if pair1 > 0 && sortedHandValues[i] != pair1Value {
				pair2++
			} else {
				pair1++
				pair1Value = sortedHandValues[i]
			}
		}
	}

	if pair1 == 4 {
		return fiveOfAKind
	}
	if pair1 == 3 {
		return fourOfAKind
	}
	if (pair1 == 2 && pair2 == 1) || (pair1 == 1 && pair2 == 2) {
		return fullHouse
	}
	if pair1 == 2 || pair2 == 2 {
		return threeOfAKind
	}
	if pair1 == 1 {
		if pair2 == 1 {
			return twoPair
		}
		return onePair
	}
	return highCard
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
	hand.handType = calculateHandType(hand.cardValues)

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
