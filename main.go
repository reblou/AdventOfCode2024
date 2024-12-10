package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/adventofcode2024/days"
)

func main() {
	fmt.Println("===AdventOfCode2024===")

	if len(os.Args) < 2 {
		fmt.Println("Too few arguments. Run with day as int argument")
		return
	}

	day, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Couldn't parse day argument.")
		return
	}

	daySelector(day)
}

func daySelector(day int) {
	switch day {
	case 1:
		days.One()
	case 2:
		days.Two()
	case 3:
		days.Three()
	case 4:
		days.Four()
	case 5:
		days.Five()
	case 6:
		days.Six()
	case 7:
		days.Seven()
	case 8:
		days.Eight()
	case 9:
		days.Nine()
	case 10:
		days.Ten()
	default:
		fmt.Println("Unknown day argument")
	}
}
