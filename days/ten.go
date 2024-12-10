package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/ten
var input10 string

func Ten() {
	fmt.Println("-Day 10-")

	topMap := parseInput2dIntSlice(input10)

	fmt.Printf("Trailhead sum: %v\n", countTrailheads(topMap))
}

func parseInput2dIntSlice(input string) [][]int {
	lines := strings.Split(input, "\r\n")

	var output [][]int

	for _, line := range lines {
		if line == "" {
			continue
		}
		chars := strings.Split(line, "")
		row := make([]int, len(chars))
		for i, char := range chars {
			n, err := strconv.Atoi(char)
			if err != nil {
				n = -1
			}
			row[i] = n
		}

		output = append(output, row)
	}

	return output
}

func countTrailheads(m [][]int) int {
	trailHeadScoreSum := 0

	for y := range m {
		for x := range m[y] {
			if m[y][x] == 0 {
				//check for trails
				peaksSeen := make(map[string]bool)
				trailScore := checkTrailHead(m, x, y, 0, peaksSeen)
				trailHeadScoreSum += trailScore
			}
		}
	}

	return trailHeadScoreSum
}

func checkTrailHead(m [][]int, x int, y int, level int, peaksSeen map[string]bool) int {
	if level == 9 {
		// end of trail
		c := coord{x, y}
		if peaksSeen[c.GetHashKey()] {
			return 0
		} else {
			peaksSeen[c.GetHashKey()] = true
			return 1
		}
	}

	next := level + 1
	count := 0
	// 4 search directions
	dirs := []coord{coord{0, -1}, coord{1, 0}, coord{0, 1}, coord{-1, 0}}

	for _, dir := range dirs {
		ny := y + dir.y
		nx := x + dir.x

		if (ny < 0) || (nx < 0) || (ny >= len(m)) || (nx >= len(m[0])) {
			continue
		}

		if m[ny][nx] == next {
			count += checkTrailHead(m, nx, ny, next, peaksSeen)
		}
	}

	return count
}
