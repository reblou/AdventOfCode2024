package days

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type clawmachine struct {
	A     coord
	B     coord
	prize coord
}

//go:embed inputs/thirteen
var input13 string

func Thirteen() {
	fmt.Println("-Day 13-")

	machines := parseInput13(input13)
	fmt.Printf("Min total tokens: %v\n", lowestTotalTokens(machines))

	fmt.Printf("Min total tokens 2: %v\n", lowestTokensConversion(machines))
}

func parseInput13(input string) []clawmachine {
	var output []clawmachine

	lines := strings.Split(input, "\r\n")
	var claw clawmachine
	r := regexp.MustCompile("(?P<a>[0-9]+)[^0-9]*(?P<b>[0-9]+)")

	for _, line := range lines {
		if line == "" {
			output = append(output, claw)
			claw = clawmachine{}
			continue
		}
		m := r.FindStringSubmatch(line)
		a, _ := strconv.Atoi(m[r.SubexpIndex("a")])
		b, _ := strconv.Atoi(m[r.SubexpIndex("b")])

		if strings.Contains(line, "A") {
			claw.A.x = a
			claw.A.y = b
		} else if strings.Contains(line, "B") {
			claw.B.x = a
			claw.B.y = b
		} else if strings.Contains(line, "Prize") {
			claw.prize.x = a
			claw.prize.y = b
		} else {
			panic("invalid line contents")
		}
	}

	return output
}

func lowestTotalTokens(input []clawmachine) int {
	var total int
	for _, machine := range input {
		total += lowestTokens(machine.A, machine.B, machine.prize)
	}

	return total
}

func lowestTokens(a coord, b coord, prize coord) int {
	lowest := 0
	for i := 0; i <= 100; i++ {
		// calc b possible?
		tx := prize.x - (i * a.x)
		ty := prize.y - (i * a.y)
		if tx%b.x == 0 && ty%b.y == 0 {
			// division possible
			// check if same factor times x and y = tx / ty
			n := (tx / b.x)
			if n != (ty / b.y) {
				continue
			}

			cost := (3 * i) + n

			if lowest == 0 || cost < lowest {
				lowest = cost
			}
		}
	}

	return lowest
}

func lowestTokensConversion(input []clawmachine) int {
	var total int
	for _, machine := range input {
		machine.prize.x += 10000000000000
		machine.prize.y += 10000000000000

		total += cost(machine)
	}

	return total
}

func cost(m clawmachine) int {
	ca := float64(m.prize.x*m.B.y-(m.prize.y*m.B.x)) / float64((m.A.x*m.B.y)-(m.A.y*m.B.x))

	cb := (float64(m.prize.x) - float64(m.A.x)*ca) / float64(m.B.x)

	if ca == float64(int64(ca)) && cb == float64(int64(cb)) {
		return (3*int(ca) + int(cb))
	} else {
		return 0
	}
}
