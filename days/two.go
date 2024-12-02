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
		increasing := true

		if report[0] > report[1] {
			increasing = false
		}

		rowSafe := true
		dampener := true

		for i := 1; i < len(report); i++ {
			diff := absDiff(report[i], report[i-1])

			if diff == 0 {
				if !dampener {
					rowSafe = false
					break
				}
				dampener = false
				report[i] = report[i-1]
			} else if increasing && report[i] <= report[i-1] {
				if i < len(report)-2 && report[i] >= report[i+1] {
					//pretend i-1 replaced
					dampener = false
					increasing = !increasing
					continue
				} else if i < len(report)-2 && report[i-1] == report[i+1] {
					// replace report[i-1] instead
					dampener = false
					continue
				}

				report[i] = report[i-1]

				if !dampener {
					rowSafe = false
					break
				}
				dampener = false
			} else if !increasing && report[i] >= report[i-1] {
				if i < len(report)-2 && report[i] <= report[i+1] {
					//pretend i-1 replaced
					dampener = false
					increasing = !increasing
					continue
				} else if i < len(report)-2 && report[i-1] == report[i+1] {
					// replace report[i-1] instead
					dampener = false
					continue
				}

				report[i] = report[i-1]
				if !dampener {
					rowSafe = false
					break
				}

				dampener = false
			} else if diff > 3 {
				if !dampener {
					rowSafe = false
					break
				}

				// removing just i will never create a safe report, unless it's last in report
				// "remove" i-1 need to check i-2
				if i >= 2 && i != len(report)-1 {
					report[i-1] = report[i-2]
					i -= 1
				}
				dampener = false
			}
		}

		if rowSafe {
			safe += 1
		} else {
			fmt.Printf("%v\n", report)
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
