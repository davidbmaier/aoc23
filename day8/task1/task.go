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

type Node struct {
	left  string
	right string
}

var instructionRegex = regexp.MustCompile(`[LR]+`)

// three capturing groups for the three relevant node names
var lineRegex = regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines := strings.Split(inputs, "\n")

	// turn the instructions into a rune array for nicer conditions later on
	instructions := []rune(string(instructionRegex.Find([]byte(lines[0]))))

	// assemble map of all nodes
	nodesMap := map[string]Node{}
	for i := 2; i < len(lines); i++ {
		line := lines[i]

		matches := lineRegex.FindSubmatch([]byte(line))
		node := Node{
			left:  string(matches[2]),
			right: string(matches[3]),
		}

		nodesMap[string(matches[1])] = node
	}

	endNodeFound := false
	instructionCounter := 0
	currentNodeName := "AAA"
	endNodeName := "ZZZ"
	// follow instructions until end node was found
	for !endNodeFound {
		// use modulo to keep the instructionCounter intact (because it's also the final result)
		mappedInstructionCounter := instructionCounter % len(instructions)
		instruction := instructions[mappedInstructionCounter]
		// follow instruction
		if instruction == 'L' {
			currentNodeName = nodesMap[currentNodeName].left
		} else {
			currentNodeName = nodesMap[currentNodeName].right
		}

		// finish condition
		if currentNodeName == endNodeName {
			endNodeFound = true
		} else {
			instructionCounter++
		}
	}

	// plus 1 because the counter is 0-indexed
	fmt.Print(instructionCounter + 1)
}
