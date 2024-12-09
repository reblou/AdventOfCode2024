package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/nine
var input9 string

type pointer struct {
	i            int
	curFileCount int
}

// return index
func (p *pointer) Pop() int {
	if p.curFileCount > 0 {
		p.curFileCount--
	} else {
		p.i -= 2
		//TODO: need to get file count
	}

	return p.i
}

func Nine() {
	fmt.Println("-Day 9-")

	input := strings.Trim(input9, "\r\n")
	fmt.Printf("Checksum: %v\n", calcChecksum(input))
}

func calcChecksum(input string) int {
	defragged := defragFromEncoding(input)

	// fmt.Printf("Defragged: %v\n", defragged)
	return calculateCheckSum(defragged)
}

func defragFromEncoding(input string) []int {
	p1 := 0
	p2 := len(input) - 1
	p1File := true
	p2File := len(input)%2 == 1
	if !p2File {
		p2 -= 1
		p2File = true
	}
	p2FileLen := strToInt(string(input[p2]))

	var defragged []int

	for p1 < p2+p2FileLen { // todo re-evaluate this condition
		if p1File {
			if p1 >= p2 {
				for range p2FileLen {
					defragged = append(defragged, p1/2)
					p2FileLen--
				}
				break
			}
			// add x ints of id
			for range strToInt(string(input[p1])) {
				defragged = append(defragged, p1/2)
			}
		} else {
			// fill gap with p2
			for range strToInt(string(input[p1])) {
				if p2FileLen <= 0 {
					p2 -= 2
					p2FileLen = strToInt(string(input[p2]))
				}
				if p2 <= p1 {
					p2FileLen = 0
					break
				}
				defragged = append(defragged, p2/2)
				p2FileLen--
			}
		}

		p1++
		p1File = !p1File
	}

	return defragged
}

func calculateCheckSum(defragged []int) int {
	var checksum int

	for i, n := range defragged {
		checksum += i * n
	}
	return checksum
}

func strToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
