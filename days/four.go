package days

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs/four
var input4 string

func Four() {
	fmt.Println("-Day Four-")

	count := countMultDirectionalXmas(input4)
	fmt.Printf("Xmas Total: %v\n", count)

	count2 := countCrossMas(input4)
	fmt.Printf("X-Mas Total: %v\n", count2)
}

func countMultDirectionalXmas(input string) int {
	count := 0

	count += countXmas(input)

	manip := manipulateInput(input)

	count += countXmas(manip[0])

	count += countXmas(manip[1])

	count += countXmas(manip[2])
	return count
}

func manipulateInput(input string) []string {
	// take input string and return up-down, diag- converted to l-r strings
	lines := strings.Split(input, "\r\n")
	lineLen := len(lines[0])
	n := len(lines)

	verts := transformVertical(lines, lineLen, n)
	diagslr := transformDiag(lines, lineLen, n)
	diagsrl := transformDiagRL(lines, n, lineLen)

	output := []string{
		verts,
		diagslr,
		diagsrl,
	}

	return output
}

func transformVertical(lines []string, lineLen int, lineCount int) string {
	var output string

	for x := 0; x < lineLen; x++ {
		for y := 0; y < lineCount; y++ {
			output += string(lines[y][x])
		}
		output += "\r\n"
	}

	return output
}

func transformDiag(lines []string, lineLen int, lineCount int) string {
	var output string

	for y := lineCount - 1; y >= 0; y-- {

		x := 0
		output += string(lines[y][x])
		for ny := y + 1; ny < lineCount; ny++ {
			x++
			output += string(lines[ny][x])
		}
		output += "."
	}
	output += "\r\n"

	for x := lineLen - 1; x > 0; x-- {
		y := 0
		output += string(lines[y][x])
		for nx := x + 1; nx < lineCount; nx++ {
			y++
			output += string(lines[y][nx])
		}
		output += "."
	}

	return output
}

func transformDiagRL(lines []string, lineLen int, lineCount int) string {
	var output string

	for y := lineCount - 1; y >= 0; y-- {

		x := lineLen - 1
		output += string(lines[y][x])
		for ny := y + 1; ny < lineCount; ny++ {
			x--
			output += string(lines[ny][x])
		}
		output += "."
	}
	output += "\r\n"

	for x := 0; x < lineLen-1; x++ {
		y := 0
		output += string(lines[y][x])
		for nx := x - 1; nx >= 0; nx-- {
			y++
			output += string(lines[y][nx])
		}
		output += "."
	}

	return output
}

func countXmas(input string) int {
	l := strings.Count(input, "XMAS")
	r := strings.Count(input, "SAMX")

	return l + r
}

func countCrossMas(input string) int {
	lines := strings.Split(input, "\r\n")

	var count int
	for y, line := range lines {
		n := len(line)
		for x, char := range line {
			if char != 'A' {
				continue
			}

			if masCheck(lines, x, y, n) {
				count += 1
			}
		}
	}
	return count
}

func masCheck(lines []string, x int, y int, lineLen int) bool {
	if x <= 0 || y <= 0 || y >= len(lines)-1 || x >= lineLen-1 {
		return false
	}

	lr := string(lines[y-1][x-1]) + string(lines[y][x]) + string(lines[y+1][x+1])

	rl := string(lines[y-1][x+1]) + string(lines[y][x]) + string(lines[y+1][x-1])

	return (lr == "MAS" || lr == "SAM") && (rl == "MAS" || rl == "SAM")
}
