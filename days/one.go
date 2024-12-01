package days

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs/one
var input string

func One() {
	fmt.Println("-Day one-")
	var a []int
	var b []int

	parseInput(&a, &b)
	difference := totalDistance(a, b)

	fmt.Printf("Total Difference: %v\n", difference)
	similarity := similarityScore(a, b)

	fmt.Printf("Similarity Score: %v\n", similarity)
}

func totalDistance(a []int, b []int) int {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})

	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	diffTotal := 0
	for i, val := range a {
		diff := val - b[i]
		if diff < 0 {
			diff = -diff
		}

		diffTotal += diff
	}

	return diffTotal
}

func similarityScore(a []int, b []int) int {
	// BUILD map of ints and how often they occur in list b
	// foreach in a, * by freq in map + add to total
	bFreqMap := make(map[int]int)

	for _, val := range b {
		bFreqMap[val] += 1
	}

	similarity := 0
	for _, val := range a {
		similarity += val * bFreqMap[val]
	}

	return similarity
}

func readFile(path string, a *[]int, b *[]int) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")

		inta, _ := strconv.Atoi(split[0])
		*a = append(*a, inta)

		intb, _ := strconv.Atoi(split[1])
		*b = append(*b, intb)
	}
}

func parseInput(a *[]int, b *[]int) {
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "   ")

		inta, _ := strconv.Atoi(split[0])
		*a = append(*a, inta)

		intb, _ := strconv.Atoi(split[1])
		*b = append(*b, intb)
	}
}
