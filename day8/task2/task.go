package main

import (
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Node struct {
	left  string
	right string
}

var instructionRegex = regexp.MustCompile(`[LR]+`)
var lineRegex = regexp.MustCompile(`([A-Z0-9]+) = \(([A-Z0-9]+), ([A-Z0-9]+)\)`)

// Euclid's algorithm
func greatestCommonDivisor(a, b int) int {
	tempValue := 0
	for b != 0 {
		tempValue = a % b
		a = b
		b = tempValue
	}
	return a
}

// lcm implementation using the gcd
func leastCommonMultiple(integers []int) int {
	// start with the first two
	result := integers[0] * integers[1] / greatestCommonDivisor(integers[0], integers[1])

	// if necessary, go through more values one by one recursively
	for i := 2; i < len(integers); i++ {
		result = leastCommonMultiple([]int{result, integers[i]})
	}

	return result
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	instructions := []rune(string(instructionRegex.Find([]byte(lines[0]))))
	nodesMap := map[string]Node{}

	instructionCounter := 0
	currentNodeNames := []string{}
	finalCounters := []int{}
	endNodeRune := 'Z'

	// assemble map of all nodes
	for i := 2; i < len(lines); i++ {
		line := lines[i]

		matches := lineRegex.FindSubmatch([]byte(line))
		node := Node{
			left:  string(matches[2]),
			right: string(matches[3]),
		}

		nodeName := string(matches[1])

		nodesMap[nodeName] = node
		// identify starting nodes
		if rune(nodeName[2]) == 'A' {
			currentNodeNames = append(currentNodeNames, nodeName)
		}
	}

	// follow instructions until no nodes have to be resolved anymore
	for len(currentNodeNames) > 0 {
		mappedInstructionCounter := instructionCounter % len(instructions)
		instruction := instructions[mappedInstructionCounter]

		instructionCounter++
		// temporary array of finished nodes
		// (because playing around with the current nodes during the loop would be a bad idea)
		finishedNodeNames := []string{}

		for j, currentNodeName := range currentNodeNames {
			if instruction == 'L' {
				currentNodeNames[j] = nodesMap[currentNodeName].left
			} else {
				currentNodeNames[j] = nodesMap[currentNodeName].right
			}

			// check if we've reached an end node - if so, remember that
			if rune(currentNodeNames[j][2]) == endNodeRune {
				finishedNodeNames = append(finishedNodeNames, currentNodeNames[j])
			}
		}

		// go through known finished nodes
		for _, finishedNodeName := range finishedNodeNames {
			// remove node from currentNodeNames
			nodeIndex := slices.Index(currentNodeNames, finishedNodeName)
			currentNodeNames = append(currentNodeNames[:nodeIndex], currentNodeNames[nodeIndex+1:]...)
			// add current counter to final counters
			finalCounters = append(finalCounters, instructionCounter)
		}
	}

	// get the least common multiple for all final counters
	lcm := leastCommonMultiple(finalCounters)

	fmt.Print(lcm)
}
