package days

import (
	_ "embed"
	"fmt"
	"strconv"
)

func (c *coord) GetHashKey() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
}

//go:embed inputs/eight
var input8 string

func Eight() {
	fmt.Println("-Day 8-")

	inputParsed := parseInput2dStrSlice(input8)

	fmt.Printf("Unique Antinodes: %v\n", uniqueAntinodes(inputParsed))
	fmt.Printf("Part 2: %v\n", uniqueAntinodesResHarmonics(inputParsed))
}

func uniqueAntinodes(input [][]string) int {
	antiNodes := make(map[string]bool)
	antennas := make(map[string][]coord)

	for y, line := range input {
		for x, char := range line {
			if char == "." {
				continue
			}

			// if letter, check for pair + mark antinodes
			checkForAntinodes(input, antiNodes, antennas, char, coord{x, y})
		}
	}

	return len(antiNodes)
}

func uniqueAntinodesResHarmonics(input [][]string) int {
	antiNodes := make(map[string]bool)
	antennas := make(map[string][]coord)

	for y, line := range input {
		for x, char := range line {
			if char == "." || char == "#" {
				continue
			}

			// if letter, check for pair + mark antinodes
			checkForResHarmAntinodes(input, antiNodes, antennas, char, coord{x, y})
		}
	}

	return len(antiNodes)
}

func checkForAntinodes(layout [][]string, antiNodes map[string]bool, antennas map[string][]coord, char string, pos coord) {

	matches, ok := antennas[char]

	antennas[char] = append(antennas[char], pos)

	if !ok {
		// no other of char, return
		return
	}
	// foreach pair, calculate antinodes
	maxX := len(layout[0])
	maxY := len(layout)
	for _, pair := range matches {
		setAntinode(pair, pos, antiNodes, maxX, maxY)
		setAntinode(pos, pair, antiNodes, maxX, maxY)
	}
}

func setAntinode(a coord, b coord, antiNodes map[string]bool, maxX int, maxY int) {
	diff := coord{a.x - b.x, a.y - b.y}
	antinode := coord{a.x + diff.x, a.y + diff.y}
	if inBounds(antinode, maxX, maxY) {
		antiNodes[antinode.GetHashKey()] = true
	}
}

func checkForResHarmAntinodes(layout [][]string, antiNodes map[string]bool, antennas map[string][]coord, char string, pos coord) {
	matches, ok := antennas[char]

	antennas[char] = append(antennas[char], pos)

	if !ok {
		// no other of char, return
		return
	}
	// foreach pair, calculate antinodes
	maxX := len(layout[0])
	maxY := len(layout)
	antiNodes[pos.GetHashKey()] = true
	for _, pair := range matches {
		antiNodes[pair.GetHashKey()] = true
		setResHarmAntinode(pair, pos, antiNodes, maxX, maxY)
		setResHarmAntinode(pos, pair, antiNodes, maxX, maxY)
	}
}

func setResHarmAntinode(a coord, b coord, antiNodes map[string]bool, maxX int, maxY int) {
	// loop recursively until not in bounds
	diff := coord{a.x - b.x, a.y - b.y}
	antinode := coord{a.x + diff.x, a.y + diff.y}
	if inBounds(antinode, maxX, maxY) {
		antiNodes[antinode.GetHashKey()] = true
		setResHarmAntinode(antinode, a, antiNodes, maxX, maxY)
	}
}

func inBounds(c coord, maxX int, maxY int) bool {
	return c.y >= 0 && c.y < maxY && c.x >= 0 && c.x < maxX
}
