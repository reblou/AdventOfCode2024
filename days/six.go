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

type moveInfo struct {
	location coord
	vector   coord
}

func Six() {
	fmt.Println("-Day Six-")

	guardPath := makeGuardPath(input6)
	fmt.Printf("Part 1: %v\n", countXs(guardPath))

	fmt.Printf("Part 2: %v\n", countLoopObsticles(parseInput2dStrSlice(input6), guardPath))
}

func parseInput2dStrSlice(input string) [][]string {
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

func makeGuardPath(input string) [][]string {
	parsed := parseInput2dStrSlice(input6)
	guardPos := getGuardPosition(parsed)

	guardMoveVector := coord{0, -1}

	layout, err := moveGuard(parsed, &guardPos, &guardMoveVector)

	for err == nil {
		//printOutput(layout)
		layout, err = moveGuard(layout, &guardPos, &guardMoveVector)
	}

	return layout
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

func countLoopObsticles(input [][]string, guardPath [][]string) int {
	guardStart := getGuardPosition(input)
	//foreach X, check if making it an # creates a loop
	xs := getXs(guardPath)

	var count int
	for _, x := range xs {
		if makesLoop(guardPath, x, guardStart) {
			count++
		}
	}

	return count
}

func getXs(layout [][]string) []coord {
	var coords []coord
	for y, line := range layout {
		for x, char := range line {
			if char == "X" {
				coords = append(coords, coord{x, y})
			}
		}
	}
	return coords
}

func makesLoop(layout [][]string, obsticle coord, guardStart coord) bool {
	//run steps, check if loop, record seen #s
	guardPos := guardStart
	guardMoveVector := coord{0, -1}

	pivots := make(map[int]map[int][]moveInfo)
	for i := range len(layout) {
		pivots[i] = make(map[int][]moveInfo)
	}
	newlayout := sliceSliceDeepCopy(layout)

	newlayout[obsticle.y][obsticle.x] = "#"

	guardPos, err := nextGuardPos(newlayout, &guardPos, &guardMoveVector, pivots)

	for err == nil {
		//printOutput(newlayout)
		guardPos, err = nextGuardPos(newlayout, &guardPos, &guardMoveVector, pivots)
	}

	return errors.Is(err, ErrLoop)
}

var ErrOOB = fmt.Errorf("out of bounds")
var ErrLoop = fmt.Errorf("visited before")

func nextGuardPos(input [][]string, guard *coord, vec *coord, pivots map[int]map[int][]moveInfo) (coord, error) {
	width := len(input[0])

	next := coord{guard.x + vec.x, guard.y + vec.y}

	if next.x < 0 || next.x >= width || next.y < 0 || next.y >= len(input) {
		return next, ErrOOB
	}

	if input[next.y][next.x] == "#" {
		*vec = rotateGuard(*vec)
		prevs := pivots[guard.x][guard.y]
		if sliceContainsMove(prevs, *guard, *vec) {
			return next, ErrLoop
		}

		pivots[guard.x][guard.y] = append(pivots[guard.x][guard.y], moveInfo{*guard, *vec})
		return nextGuardPos(input, guard, vec, pivots)
	}

	return next, nil
}

func sliceContainsMove(s []moveInfo, n coord, v coord) bool {
	for _, m := range s {
		if m.location.x == n.x && m.location.y == n.y && m.vector.x == v.x && m.vector.y == v.y {
			return true
		}
	}

	return false
}

func sliceSliceDeepCopy(s [][]string) [][]string {
	var new [][]string

	for _, l := range s {
		new = append(new, append([]string{}, l...))
	}

	return new
}
