package days

import (
	"bufio"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/two
var input2 string

func Two() {
	fmt.Println("-Day two-")

	input := parseInput2()
	safeReports := safeReports(input)
	fmt.Printf("Safe reports: %v\n", safeReports)

	dampenedSafeReports := dampenerSafeReports(input)
	fmt.Printf("Dampener adjusted safe reports: %v\n", dampenedSafeReports)
}

func parseInput2() [][]int {
	scanner := bufio.NewScanner(strings.NewReader(input2))
	var output [][]int
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")

		newrow := make([]int, len(splits))
		for i, split := range splits {
			newrow[i], _ = strconv.Atoi(split)
		}
		output = append(output, newrow)
	}
	return output
}

func safeReports(input [][]int) int {
	/*
	   The levels are either all increasing or all decreasing.
	   Any two adjacent levels differ by at least one and at most three.
	*/
	safe := 0

	for _, report := range input {
		increasing := true

		if report[0] > report[1] {
			increasing = false
		}
		rowSafe := true

		for i := 1; i < len(report); i++ {
			diff := report[i] - report[i-1]
			if diff < 0 {
				diff = -diff
			}

			if diff < 1 || diff > 3 {
				rowSafe = false
				break
			}

			if (increasing && report[i] <= report[i-1]) ||
				(!increasing && report[i] >= report[i-1]) {
				rowSafe = false
				break
			}
		}

		if rowSafe {
			safe += 1
		}
	}

	return safe
}

func dampenerSafeReports(input [][]int) int {
	// if removing one level in report makes it safe, count it as safe.
	safe := 0

	for _, report := range input {
		rowSafe := reportSafe(report)

		if !rowSafe {
			// loop and try removing all rows to find a safe iteration
			rowSafe = trySafeAlternatives(report)
		}

		if rowSafe {
			safe += 1
		}
	}

	return safe
}

func absDiff(a int, b int) int {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}

	return diff
}

func reportSafe(report []int) bool {
	increasing := true

	if report[0] > report[1] {
		increasing = false
	}

	for i := 1; i < len(report); i++ {
		diff := absDiff(report[i], report[i-1])

		if (diff < 1 || diff > 3) ||
			(increasing && report[i] <= report[i-1]) ||
			(!increasing && report[i] >= report[i-1]) {
			return false
		}
	}
	return true
}

func trySafeAlternatives(report []int) bool {
	for i := range report {
		reportCopy := make([]int, len(report))
		copy(reportCopy, report)

		newReport := append(reportCopy[:i], reportCopy[i+1:]...)
		if reportSafe(newReport) {
			return true
		}
	}
	return false
}
