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

// map of card strings to values (lowest to highest card)
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

// constants to make the order of hand types easier to understand
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

// helper function to calculate the type of hand (see constants above)
func calculateHandType(handValues []int) int {
	// probably not the smartest way but it works fine
	// five values, so there are only two possible sets of duplicates (three would require at least six values)
	duplicateAmount1 := 0
	duplicateAmount2 := 0
	duplicateAmount1Value := -1 // initialize the first duplicate's value with something it can never be

	// deep-copy the hand values so sorting them doesn't affect the original values for later
	// (since we still need the original order during the final sorting of hands)
	sortedHandValues := make([]int, len(handValues))
	copy(sortedHandValues, handValues)

	// sort the hand's values to make it easy to compare neighbors (when sorted, duplicates will always be neigbors)
	sort.Slice(sortedHandValues, func(i, j int) bool {
		return sortedHandValues[i] < sortedHandValues[j]
	})

	// go through all hand values/cards
	for i := 0; i < len(sortedHandValues); i++ {
		// ignore the last card and cards that aren't equal to the next one
		if i < len(sortedHandValues)-1 && sortedHandValues[i] == sortedHandValues[i+1] {
			// this isn't the last card + this card is a duplicate of the next one
			// if there's already a first set of duplicates, and this is not the same underlying value, start counting the second set
			if duplicateAmount1 > 0 && sortedHandValues[i] != duplicateAmount1Value {
				duplicateAmount2++
			} else {
				// by default, increase the first set of duplicates
				duplicateAmount1++
				// make sure we remember what underlying value the first set of duplicates is
				duplicateAmount1Value = sortedHandValues[i]
			}
		}
	}

	// four duplicates = five equal cards
	if duplicateAmount1 == 4 {
		return fiveOfAKind
	}
	// three duplicates = four equal cards (and one spare)
	if duplicateAmount1 == 3 {
		return fourOfAKind
	}
	// one set of two duplicates, one set of one duplicate = 3 + 2 = full house
	if (duplicateAmount1 == 2 && duplicateAmount2 == 1) || (duplicateAmount1 == 1 && duplicateAmount2 == 2) {
		return fullHouse
	}
	// one set of two duplicates = three equal cards (and two spares, otherwise the case above would have matched)
	if duplicateAmount1 == 2 || duplicateAmount2 == 2 {
		return threeOfAKind
	}
	// one duplicate = at least one pair
	if duplicateAmount1 == 1 {
		// another pair = two pairs
		if duplicateAmount2 == 1 {
			return twoPair
		}
		// else just one pair
		return onePair
	}
	// if none of them match, high card is the fallback
	return highCard
}

// helper function to transform a hand string into the custom Hand struct
func transformStringToHand(handString string) Hand {
	hand := Hand{
		cardValues: []int{},
		bid:        0,
	}

	// split the string to get the relevant parts
	handParts := strings.Split(handString, " ")
	// save the original card string for easier debugging
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
		// transform each line into a Hand that can easily be compared during sorting
		hands = append(hands, transformStringToHand(line))
	}

	// sort the hands to know which hand gets which multiplier
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			// not equal hand types, so sort by those
			return hands[i].handType < hands[j].handType
		}
		// equal hand types, so sort by first card difference
		for k := 0; k < len(hands[i].cardValues); k++ {
			if hands[i].cardValues[k] != hands[j].cardValues[k] {
				return hands[i].cardValues[k] < hands[j].cardValues[k]
			}
		}
		// fallback in case none of the cases happen - this will only happen for exactly equal hands
		return true
	})

	// calculate the sum by multiplying each hand's bid with its multiplier
	bidSum := 0
	for i, hand := range hands {
		bidSum += hand.bid * (i + 1)
	}

	fmt.Print(bidSum)
}
