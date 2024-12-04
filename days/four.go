package days

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed inputs/four
var input4 string

var debug = "MMMSXXMASM\r\nMSAMXMSMSA\r\nAMXSXMAAMM\r\nMSAMASMSMX\r\nXMASAMXAMM\r\nXXAMMXXAMA\r\nSMSMSASXSS\r\nSAXAMASAAA\r\nMAMMMXMMMM\r\nMXMXAXMASX"

func Four() {
	fmt.Println("-Day Four-")

	count := countMultDirectionalXmas(input4)
	fmt.Printf("Total: %v\n", count)
}

func countMultDirectionalXmas(input string) int {
	count := 0

	lr := countXmas(input)
	fmt.Printf("LR Count: %v\n", lr)
	count += lr

	manip := manipulateInput(input)

	v := countXmas(manip[0])
	count += v
	fmt.Printf("Vert Count: %v\n", v)

	dlr := countXmas(manip[1])
	count += dlr
	fmt.Printf("DiagLR Count: %v\n", dlr)

	drl := countXmas(manip[2])
	count += drl
	fmt.Printf("DiagRL Count: %v\n", drl)
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
