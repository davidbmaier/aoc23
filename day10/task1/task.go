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

var startRegex = regexp.MustCompile(`S`)
var lines = []string{}

type Tile struct {
	x               int
	y               int
	directions      [2]int
	originDirection int
	value           string
}

const (
	up    = 0
	right = 1
	down  = 2
	left  = 3
)

func identifyTile(startingTile Tile, direction int) Tile {
	tile := Tile{}

	switch direction {
	case up:
		tile.x = startingTile.x
		tile.y = startingTile.y - 1
	case right:
		tile.x = startingTile.x + 1
		tile.y = startingTile.y
	case down:
		tile.x = startingTile.x
		tile.y = startingTile.y + 1
	case left:
		tile.x = startingTile.x - 1
		tile.y = startingTile.y
	}

	if tile.y >= len(lines) || tile.y < 0 || tile.x >= len(lines[0]) || tile.x < 0 {
		// tile position outside the grid
		return Tile{x: -1, y: -1}
	}

	tileValue := rune(lines[tile.y][tile.x])

	switch tileValue {
	case 'F':
		tile.directions = [2]int{right, down}
	case 'J':
		tile.directions = [2]int{left, up}
	case 'L':
		tile.directions = [2]int{up, right}
	case '7':
		tile.directions = [2]int{left, down}
	case '-':
		tile.directions = [2]int{left, right}
	case '|':
		tile.directions = [2]int{up, down}
	}

	// for debugging purposes
	tile.value = string(tileValue)

	return tile
}

func findNextTile(startingTile Tile, previousDirection int, nextDirection int) Tile {
	if nextDirection == -1 {
		// check all directions
		for i := 0; i < 4; i++ {
			if i == previousDirection {
				continue
			}

			neighborTile := identifyTile(startingTile, i)
			if neighborTile.x == -1 && neighborTile.y == -1 {
				continue
			}

			counterDirection := i - 2
			if counterDirection < 0 {
				counterDirection += 4
			}

			if neighborTile.directions[0] == counterDirection || neighborTile.directions[1] == counterDirection {
				neighborTile.originDirection = i
				return neighborTile
			}
		}
	}

	nextTile := identifyTile(startingTile, nextDirection)
	nextTile.originDirection = nextDirection
	return nextTile
}

func findNextDirection(tile Tile) int {
	// find the opposite direction for the origin (because we need the other one for the next direction)
	counterDirection := tile.originDirection - 2
	if counterDirection < 0 {
		counterDirection += 4
	}

	if tile.directions[0] == counterDirection {
		return tile.directions[1]
	} else {
		return tile.directions[0]
	}
}

func main() {
	data, err := os.ReadFile("./data.txt")
	check(err)
	inputs := string(data)
	lines = strings.Split(inputs, "\n")

	// find starting point
	start := Tile{}
	for i, line := range lines {
		startIndex := startRegex.FindIndex([]byte(line))
		if startIndex != nil {
			start = Tile{
				x: startIndex[0],
				y: i,
			}
		}
	}

	tile1 := findNextTile(start, -1, -1)
	fmt.Printf("Tile 1: %s (%d/%d)\n", tile1.value, tile1.x, tile1.y)
	tile2 := findNextTile(start, tile1.originDirection, -1)
	fmt.Printf("Tile 2: %s (%d/%d)\n", tile2.value, tile2.x, tile2.y)

	stepCounter := 1 // start with 1 since the first step has already been taken
	// repeat until the two paths find the same tile again
	for tile1.x != tile2.x || tile1.y != tile2.y {

		tile1NextDirection := findNextDirection(tile1)
		tile1 = findNextTile(tile1, tile1.originDirection, tile1NextDirection)
		fmt.Printf("Tile 1: %s (%d/%d)\n", tile1.value, tile1.x, tile1.y)

		tile2NextDirection := findNextDirection(tile2)
		tile2 = findNextTile(tile2, tile2.originDirection, tile2NextDirection)
		fmt.Printf("Tile 2: %s (%d/%d)\n", tile2.value, tile2.x, tile2.y)

		stepCounter++
	}

	fmt.Print(stepCounter)
}
