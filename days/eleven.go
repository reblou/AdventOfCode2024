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

	fmt.Printf("After 25 blinks: %v\n", lenAfterXBlinks(input, 25))

	fmt.Printf("After 75 blinks: %v\n", dpLenAfterXBlinks(input, 75))
}

func lenAfterXBlinks(input []int, blinks int) int {
	for range blinks {
		input = blink(input)
	}

	return len(input)
}

func dpLenAfterXBlinks(input []int, blinks int) int {
	var sum int
	store := make(map[int]map[int]int)
	for _, n := range input {
		count := blinkXTimes(n, blinks, store)
		sum += count
	}

	return sum
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

func blinkXTimes(n int, steps int, store map[int]map[int]int) int {
	// if result is stored already
	l, exists := store[n][steps]
	if exists {
		return l
	}
	_, mapExists := store[n]
	if !mapExists {
		store[n] = make(map[int]int)
	}

	if steps == 0 {
		store[n][steps] = 1
		return 1
	}

	var count int
	if n == 0 {
		// output i = 1
		count = blinkXTimes(1, steps-1, store)
		store[n][steps] = count
		return count
	}
	str := strconv.Itoa(n)
	strLen := len(str)
	if strLen%2 == 0 {
		mid := strLen / 2
		a, _ := strconv.Atoi(str[:mid])
		b, _ := strconv.Atoi(str[mid:])

		count = blinkXTimes(a, steps-1, store) + blinkXTimes(b, steps-1, store)
	} else {
		count = blinkXTimes(n*2024, steps-1, store)
	}

	store[n][steps] = count
	return count
}
