package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/eleven
var input11 string

func Eleven() {
	fmt.Println("-Day 11-")

	input := parseInputIntSlice(input11)

	fmt.Printf("After blinks: %v\n", lenAfterXBlinks(input, 25))
}

func lenAfterXBlinks(input []int, blinks int) int {
	for range blinks {
		input = blink(input)
	}

	return len(input)
}

func parseInputIntSlice(input string) []int {
	splits := strings.Split(input, " ")
	output := make([]int, len(splits))

	for i, s := range splits {
		v, _ := strconv.Atoi(s)
		output[i] = v
	}

	return output
}

func blink(data []int) []int {
	var output []int
	for _, n := range data {
		if n == 0 {
			// output i = 1
			output = append(output, 1)
			continue
		}
		str := strconv.Itoa(n)
		if len(str)%2 == 0 {
			// if even number of digits
			// split
			mid := len(str) / 2
			a, _ := strconv.Atoi(str[:mid])
			b, _ := strconv.Atoi(str[mid:])

			// up to i
			output = append(output, a)
			output = append(output, b)
		} else {
			output = append(output, n*2024)
		}
	}

	return output
}
