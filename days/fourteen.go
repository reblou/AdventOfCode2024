package days

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type guard struct {
	p coord
	v coord
}

//go:embed inputs/fourteen
var input14 string

func Fourteen() {
	fmt.Println("-Day 14-")

	guards := parseInput14(input14)

	fmt.Printf("Score: %v\n", minSafetyScore(guards, 101, 103, 100))
	// fmt.Printf("Score: %v\n", minSafetyScore(guards, 11, 7, 100))
}

func parseInput14(input string) []guard {

	lines := strings.Split(input, "\r\n")

	output := make([]guard, len(lines))
	r := regexp.MustCompile(`-?\d+`)
	for l, line := range lines {
		matches := r.FindAllString(line, -1)

		nums := make([]int, len(matches))
		for i := range matches {
			n, _ := strconv.Atoi(matches[i])
			nums[i] = n
		}

		output[l] = guard{coord{nums[0], nums[1]}, coord{nums[2], nums[3]}}
	}

	return output
}

func minSafetyScore(guards []guard, w int, l int, s int) int {
	//101 tiles wide and 103 tiles tall
	// apply v 100 times, mod x by 101 and y by 103
	quads := make([][]int, 2)
	quads[0] = make([]int, 2)
	quads[1] = make([]int, 2)

	for _, guard := range guards {
		newp := coord{guard.p.x + (s * guard.v.x), guard.p.y + (s * guard.v.y)}

		newp.x = mod(newp.x, w)
		newp.y = mod(newp.y, l)

		// calculate quadrant
		y, x := getQuad(newp, w, l)
		if y != -1 && x != -1 {
			quads[y][x]++
		}
	}

	safety := 1
	for i := range quads {
		for n := range quads[i] {
			safety *= quads[i][n]
		}
	}

	return safety
}

func getQuad(g coord, w int, l int) (int, int) {
	var y int
	var x int
	if g.y < (l)/2 {
		y = 0
	} else if g.y == (l)/2 {
		y = -1
	} else {
		y = 1
	}

	if g.x == (w)/2 {
		x = -1
	} else if g.x < (w)/2 {
		x = 0
	} else {
		x = 1
	}

	return y, x
}

func mod(a, b int) int {
	return (a%b + b) % b
}
