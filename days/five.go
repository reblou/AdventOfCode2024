package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/five
var input5 string

func Five() {
	fmt.Println("-Day five-")

	orders := make(map[int][]int)
	//seen := make(map[int]bool)
	updates := parseInput5(input5, orders)

	fmt.Printf("order: %v\n", orders)
	fmt.Printf("updates: %v\n", updates)

}

func parseInput5(input string, orders map[int][]int) [][]int {
	lines := strings.Split(input, "\r\n")

	spliti := -1
	for i, line := range lines {
		if line == "" {
			spliti = i
			break
		}

		x, y := parseXBeforeY(line)

		orders[y] = append(orders[y], x)
	}

	var updates [][]int
	for n := spliti + 1; n < len(lines); n++ {
		// parse updates to [][]int
		updates = append(updates, parseUpdate(lines[n]))
	}

	return updates
}

func parseXBeforeY(input string) (int, int) {
	bits := strings.Split(input, "|")

	x, _ := strconv.Atoi(bits[0])
	y, _ := strconv.Atoi(bits[1])

	return x, y
}

func parseUpdate(input string) []int {
	output := make([]int, 0)
	nums := strings.Split(input, ",")

	for _, num := range nums {
		n, _ := strconv.Atoi(num)
		output = append(output, n)
	}

	return output
}
