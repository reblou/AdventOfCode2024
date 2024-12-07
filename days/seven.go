package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type Calibration struct {
	testValue int
	nums      []int
}

//go:embed inputs/seven
var input7 string

func Seven() {
	fmt.Println("-Day seven-")

	fmt.Printf("Total Calibration Result: %v\n", totalCalibrationResult(input7))
}

func totalCalibrationResult(input string) int {
	calibrations := parse7(input)

	fmt.Printf("%v\n", calibrations)

	return -1
}

func parse7(input string) []Calibration {
	lines := strings.Split(input, "\r\n")

	calibrations := make([]Calibration, len(lines))
	for i, line := range lines {
		nums := strings.Split(line, " ")

		testVal, _ := strconv.Atoi(strings.Trim(nums[0], ":")) // minus :

		calibrations[i] = Calibration{testVal, strSliceToInt(nums[1:])}
	}

	return calibrations
}

func strSliceToInt(s []string) []int {
	ints := make([]int, len(s))
	for i, e := range s {
		ints[i], _ = strconv.Atoi(e)
	}

	return ints
}
