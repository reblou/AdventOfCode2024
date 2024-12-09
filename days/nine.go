package days

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed inputs/nine
var input9 string

type file struct {
	id   int
	size int
}

func Nine() {
	fmt.Println("-Day 9-")

	input := strings.Trim(input9, "\r\n")
	fmt.Printf("Part 1 Checksum: %v\n", calcChecksum(input, false))
	fmt.Printf("Part 2 Checksum: %v\n", calcChecksum(input, true))
}

func calcChecksum(input string, wholeFiles bool) int {
	var defragged []int
	if wholeFiles {
		defragged = defragWholeFiles(input)
	} else {

		defragged = defragFromEncoding(input)
	}

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
	p2FileLen := charToInt(input[p2])

	var defragged []int

	for p1 < p2+p2FileLen {
		if p1File {
			if p1 >= p2 {
				for range p2FileLen {
					defragged = append(defragged, p1/2)
					p2FileLen--
				}
				break
			}
			// add x ints of id
			for range charToInt(input[p1]) {
				defragged = append(defragged, p1/2)
			}
		} else {
			// fill gap with p2
			for range charToInt(input[p1]) {
				if p2FileLen <= 0 {
					p2 -= 2
					p2FileLen = charToInt(input[p2])
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

func defragWholeFiles(input string) []int {
	files := getFiles(input)

	p := 0
	var defragged []int

	file := true
	for p < len(input) {
		if file {
			id := p / 2
			index, exists := sliceContainsFileId(id, files)
			if !exists {
				// this is now a gap because the file was moved
				file = !file
				continue
			}
			for range charToInt(input[p]) {
				defragged = append(defragged, id)
			}
			files = removeFromSlice(index, files)
		} else {
			gapSize := charToInt(input[p])

			for gapSize > 0 {
				gapMem, gapLeft := fillGap(&files, gapSize)
				defragged = append(defragged, gapMem...)
				gapSize = gapLeft
			}
		}

		p++
		file = !file
	}

	return defragged
}

func getFiles(input string) []file {
	files := make([]file, len(input)/2+1)

	n := 0
	gap := false
	for i := range input {
		if gap {
			gap = !gap
			continue
		}

		files[n] = file{n, charToInt(input[i])}
		n++
		gap = !gap
	}

	return files
}

func fillGap(files *[]file, gapSize int) ([]int, int) {
	for i := len(*files) - 1; i >= 0; i-- {
		fileSize := (*files)[i].size
		if fileSize <= gapSize {
			mem := make([]int, fileSize)
			for n := range fileSize {
				mem[n] = (*files)[i].id
			}

			gapLeft := gapSize - fileSize

			*files = removeFromSlice(i, *files)
			return mem, gapLeft
		}
	}

	// else fill with 0s to represent empty
	return make([]int, gapSize), 0
}

func calculateCheckSum(defragged []int) int {
	var checksum int

	for i, n := range defragged {
		checksum += i * n
	}
	return checksum
}

func removeFromSlice(i int, s []file) []file {
	return append(s[:i], s[i+1:]...)
}

func sliceContainsFileId(id int, files []file) (int, bool) {
	for i, f := range files {
		if f.id == id {
			return i, true
		}
	}

	return -1, false
}

func charToInt(c byte) int {
	i, _ := strconv.Atoi(string(c))
	return i
}
