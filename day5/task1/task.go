package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var numberRegex = regexp.MustCompile(`\d+`)

type MappingRule struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func stringsToInts(strings []string) []int {
	ints := []int{}
	for _, string := range strings {
		int, err := strconv.Atoi(string)
		check(err)
		ints = append(ints, int)
	}
	return ints
}

func processRuleForInput(rule MappingRule, input int) int {
	if input >= rule.sourceRangeStart && input < rule.sourceRangeStart+rule.rangeLength {
		return rule.destinationRangeStart + (input - rule.sourceRangeStart)
	}
	return -1
}

func processRulesetForInput(ruleset []MappingRule, input int) int {
	for _, rule := range ruleset {
		ruleResult := processRuleForInput(rule, input)
		if ruleResult != -1 {
			return ruleResult
		}
	}
	return input
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	seeds := []int{}
	rulesets := [][]MappingRule{}

	rulesetCounter := 0
	for i := 0; i < len(lines); i++ {
		line := lines[i]

		// go through entire input to assemble all the rulesets
		if i == 0 {
			// starting seeds
			seedsString := strings.Split(line, ": ")[1]
			seeds = stringsToInts(strings.Split(seedsString, " "))
		} else if strings.Contains(line, ":") {
			// start of mapping section
			rulesetCounter = len(rulesets) // set the counter for the following mappings
			rulesets = append(rulesets, []MappingRule{})
		} else if numberRegex.Match([]byte(line)) {
			// entry in mapping section
			entries := stringsToInts(strings.Split(line, " "))

			newRule := MappingRule{
				destinationRangeStart: entries[0],
				sourceRangeStart:      entries[1],
				rangeLength:           entries[2],
			}

			rulesets[rulesetCounter] = append(rulesets[rulesetCounter], newRule)
		}
	}

	// go through the seeds and process all rulesets until we end up with a final number for each
	seedResults := []int{}
	for i, seed := range seeds {
		seedResults = append(seedResults, seed)
		for _, ruleset := range rulesets {
			rulesetResult := processRulesetForInput(ruleset, seedResults[i])
			if rulesetResult != -1 {
				seedResults[i] = rulesetResult
			}
		}
	}

	sort.Slice(seedResults, func(i, j int) bool {
		return seedResults[i] < seedResults[j]
	})

	fmt.Print(seedResults[0])
}
