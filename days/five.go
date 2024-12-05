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

	fmt.Printf("Part 1: %v\n", middleNumSum(input5))

	fmt.Printf("Part 2: %v\n", reorderedMidNumSum(input5))
}

func middleNumSum(input string) int {
	orders := make(map[int][]int)
	updates := parseInput5(input, orders)

	total := 0
	for _, update := range updates {
		if !correctlyOrdered(update, orders) {
			continue
		}

		// if valid get middle value add to total
		total += getMidElem(update)
	}

	return total
}

func reorderedMidNumSum(input string) int {
	orders := make(map[int][]int)
	updates := parseInput5(input, orders)

	total := 0
	for _, update := range updates {
		if correctlyOrdered(update, orders) {
			continue
		}

		update = sortOrder(update, orders)
		total += getMidElem(update)
	}

	return total

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

		orders[x] = append(orders[x], y)
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

func correctlyOrdered(update []int, orders map[int][]int) bool {
	seen := make(map[int]bool, len(update))
	for _, num := range update {
		if !checkPrereqs(num, orders, seen) {
			// invalid update
			return false
		}

		seen[num] = true
	}

	return true
}

func sortOrder(update []int, order map[int][]int) []int {
	breakFlag := false
	for !correctlyOrdered(update, order) {
		indexes := make(map[int]int, len(update))

		for i, num := range update {
			if breakFlag {
				breakFlag = false
				break
			}

			indexes[num] = i
			prereqs := order[num]
			if prereqs == nil {
				continue
			}

			for _, prereq := range prereqs {
				pi, seen := indexes[prereq]
				if seen {
					// prereq needs to be after current num
					tmp := prereq
					update[pi] = update[i]
					update[i] = tmp
					breakFlag = true
					break
				}
			}
		}
	}

	return update
}

func checkPrereqs(num int, order map[int][]int, seen map[int]bool) bool {
	prereqs := order[num]
	if prereqs == nil {
		return true
	}

	for _, prereq := range prereqs {
		if seen[prereq] {
			// prereq needs to be after current num
			return false
		}
	}
	return true
}

func getMidElem(s []int) int {
	m := len(s) / 2

	return s[m]
}
