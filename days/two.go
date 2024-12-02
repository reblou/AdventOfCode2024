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
	fmt.Printf("Input: %v", input)
	safeReports := safeReports(input)
	fmt.Printf("Safe reports: %v\n", safeReports)
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

	for i, report := range input {
		for y, value := range report {

		}
	}

	return safe
}
