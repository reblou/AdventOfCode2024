package days

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed inputs/three
var input3 string

func Three() {
	fmt.Println("-Day Three-")

	multTotal := multTotal(input3)
	fmt.Printf("Mult total: %v\n", multTotal)

	fmt.Printf("Conditional Mult total: %v\n", multTotalConditionals(input3))
}

func multTotal(input string) int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	multTotal := 0
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])

		multTotal += a * b
	}

	return multTotal
}

func multTotalConditionals(input string) int {
	re := regexp.MustCompile(`(mul)\((\d{1,3}),(\d{1,3})\)|(do)\(\)|(don't)\(\)`)

	matches := re.FindAllStringSubmatch(input, -1)

	do := true
	multTotal := 0
	for _, match := range matches {
		switch match[0] {
		case "don't()":
			do = false
		case "do()":
			do = true
		default:
			if do {
				multTotal += multStrings(match[2], match[3])
			}
		}
	}

	return multTotal
}

func multStrings(a string, b string) int {
	ai, _ := strconv.Atoi(a)
	bi, _ := strconv.Atoi(b)

	return ai * bi
}
