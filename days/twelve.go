package days

import (
	_ "embed"
	"fmt"
)

//go:embed inputs/twelve
var input12 string

func Twelve() {
	fmt.Println("-Day 12-")

	garden := parseInput6(input12)
	fmt.Printf("Total Perimeter: %v\n", totalPermiter(garden))
}

func totalPermiter(garden [][]string) int {
	seen := make(map[string]bool)

	var patches []map[string]coord
	for y, _ := range garden {
		for x, _ := range garden[y] {
			c := coord{x, y}
			if seen[c.GetHashKey()] {
				continue
			}

			// else calculate full area
			patch := make(map[string]coord)
			findPatch(garden, c, seen, patch)
			// add patch to list
			patches = append(patches, patch)
			fmt.Printf("Patch: %v\n", patch)
		}
	}

	// TODO : foreach region, find perimiter and * area, add to total

	return -1
}

func findPatch(garden [][]string, c coord, seen map[string]bool, patch map[string]coord) {
	seen[c.GetHashKey()] = true
	patch[c.GetHashKey()] = c

	// search all directions for matching
	veg := garden[c.y][c.x]
	dirs := []coord{
		coord{1, 0},
		coord{0, 1},
		coord{-1, 0},
		coord{0, -1},
	}

	for _, dir := range dirs {
		search := coord{c.x + dir.x, c.y + dir.y}
		if !inBounds(search, len(garden[0]), len(garden)) {
			continue
		}

		// check if not seen and same veg type
		if !seen[search.GetHashKey()] && garden[search.y][search.x] == veg {
			// add to patch, add to seen
			seen[search.GetHashKey()] = true
			patch[search.GetHashKey()] = search
			findPatch(garden, search, seen, patch)
		}
	}
}
