package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Calibration struct {
	testValue int64
	nums      []int64
}

//go:embed inputs/seven
var input7 string

var operators = []string{"+", "*"}

func Seven() {
	fmt.Println("-Day seven-")
	fmt.Printf("Total Calibration Result: %v\n", totalCalibrationResult(input7))
	operators = append(operators, "||")
	fmt.Printf("Total Calibration Result plus ||: %v\n", totalCalibrationResult(input7))
}

func totalCalibrationResult(input string) int64 {
	calibrations := parse7(input)

	// foreach calibration, iterate all possible operations between each num return when success,
	var total int64
	for _, calibration := range calibrations {
		if trueEquation(calibration) {
			total += calibration.testValue
		}

	}

	return total
}

func parse7(input string) []Calibration {
	lines := strings.Split(input, "\r\n")

	calibrations := make([]Calibration, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, " ")

		testVal, _ := strconv.ParseInt(strings.Trim(nums[0], ":"), 10, 64) // minus :

		calibrations[i] = Calibration{testVal, strSliceToint64(nums[1:])}
	}

	return calibrations
}

func strSliceToint64(s []string) []int64 {
	int64s := make([]int64, len(s))
	for i, e := range s {
		int64s[i], _ = strconv.ParseInt(e, 10, 64)
	}

	return int64s
}

func trueEquation(c Calibration) bool {
	var opPerms [][]string
	calcPermutations(len(c.nums)-1, 0, []string{}, &opPerms)

	for _, perm := range opPerms {
		res := calculateTotal(c.nums, perm, c.testValue)
		if res == c.testValue {
			return true
		}
	}

	return false
}

func calcPermutations(size int, n int, curPerm []string, permutations *[][]string) {
	if n >= size {
		*permutations = append(*permutations, curPerm)
		return
	}

	for o := 0; o < len(operators); o++ {
		curPerm = append(curPerm, operators[o])

		calcPermutations(size, n+1, curPerm, permutations)

		var newPerm []string
		if len(curPerm)-1 > 0 {
			newPerm = make([]string, len(curPerm)-1)
		}
		copy(newPerm, curPerm[:len(curPerm)-1])
		curPerm = newPerm
	}
}

func calculateTotal(nums []int64, ops []string, target int64) int64 {
	total := nums[0]
	n := 0
	for i := 1; i < len(nums); i++ {
		total = calc(total, ops[n], nums[i])
		if total > target {
			return total
		}
		n++
	}
	return total
}

func calc(a int64, op string, b int64) int64 {
	switch op {
	case "+":
		return a + b
	case "*":
		return a * b
	case "||":
		res, err := strconv.ParseInt(strconv.FormatInt(a, 10)+strconv.FormatInt(b, 10), 10, 64)
		if err != nil {
			panic("couldn't || ints together")
		}
		return res
	}

	panic("unknown operation")
}
