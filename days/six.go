package days

import (
	_ "embed"
	"errors"
	"fmt"
	"strings"
)

//go:embed inputs/six
var input6 string

type coord struct {
	x int
	y int
}

func Six() {
	fmt.Println("-Day Six-")

	fmt.Printf("Part 1: %v\n", countGuardPath(input6))
}

func parseInput6(input string) [][]string {
	lines := strings.Split(input, "\r\n")

	var output [][]string

	for _, line := range lines {
		if line == "" {
			continue
		}
		chars := strings.Split(line, "")

		output = append(output, chars)
	}

	return output
}

func printOutput(input [][]string) {
	for _, line := range input {
		for _, char := range line {
			fmt.Printf(char)
		}
		fmt.Println()
	}
}

func countGuardPath(input string) int {
	parsed := parseInput6(input6)
	printOutput(parsed)

	guardPos := getGuardPosition(parsed)

	fmt.Printf("Guard Position: %v\n", guardPos)

	guardMoveVector := coord{0, -1}

	layout, err := moveGuard(parsed, &guardPos, &guardMoveVector)

	for err == nil {
		//printOutput(layout)
		layout, err = moveGuard(layout, &guardPos, &guardMoveVector)
	}

	return countXs(layout)
}

func getGuardPosition(input [][]string) coord {
	for y, line := range input {
		for x, char := range line {
			if char == "^" {
				return coord{x, y}
			}
		}
	}

	panic("No guard found.")
}

func moveGuard(input [][]string, guard *coord, vec *coord) ([][]string, error) {
	width := len(input[0])

	input[guard.y][guard.x] = "X"

	next := coord{guard.x + vec.x, guard.y + vec.y}

	if next.x < 0 || next.x >= width || next.y < 0 || next.y >= len(input) {
		return input, errors.New("out of bounds")
	}

	if input[next.y][next.x] == "#" {
		*vec = rotateGuard(*vec)
		return moveGuard(input, guard, vec)
	}

	*guard = next

	input[guard.y][guard.x] = "^"
	return input, nil
}

func move(guard coord, vec coord) coord {
	return coord{guard.x + vec.x, guard.y + vec.y}
}

func rotateGuard(vec coord) coord {
	// (x, y) rotated 90 degrees around (0, 0) is (-y, x).
	return coord{-vec.y, vec.x}
}

func countXs(input [][]string) int {
	count := 0
	for _, line := range input {
		for _, char := range line {
			if char == "X" {
				count++
			}
		}
	}
	return count
}
