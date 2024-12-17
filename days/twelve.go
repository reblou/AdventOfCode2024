package days

import (
	_ "embed"
	"fmt"
	"sort"
)

//go:embed inputs/twelve
var input12 string

func Twelve() {
	fmt.Println("-Day 12-")

	garden := parseInput2dStrSlice(input12)
	fmt.Printf("Total Perimeter: %v\n", totalPermiter(garden, false))
	fmt.Printf("Total Discount Perimeter: %v\n", totalPermiter(garden, true))
}

func totalPermiter(garden [][]string, discount bool) int {
	seen := make(map[string]bool)

	var total int
	var regions []map[string]coord
	for y, _ := range garden {
		for x, _ := range garden[y] {
			c := coord{x, y}
			if seen[c.GetHashKey()] {
				continue
			}

			// else calculate full area
			region := make(map[string]coord)
			findPatch(garden, c, seen, region)
			// add patch to list
			regions = append(regions, region)
		}
	}

	// foreach region, find perimiter and * area, add to total
	for _, region := range regions {
		if discount {
			sides := calcSides(region)
			total += len(region) * sides
		} else {
			perimeter := calcPerimeter(region)
			total += len(region) * perimeter
		}
	}

	return total
}

func findPatch(garden [][]string, c coord, seen map[string]bool, patch map[string]coord) {
	seen[c.GetHashKey()] = true
	patch[c.GetHashKey()] = c

	// search all directions for matching
	veg := garden[c.y][c.x]
	dirs := []coord{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
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

func calcPerimeter(region map[string]coord) int {
	var perimeter int

	dirs := []coord{
		{1, 0},
		{0, 1},
		{-1, 0},
		{0, -1},
	}
	for _, c := range region {
		// check surroundings, if also in patch, don't add perimeter
		for _, dir := range dirs {
			search := coord{c.x + dir.x, c.y + dir.y}
			_, inRegion := region[search.GetHashKey()]

			if !inRegion {
				perimeter += 1
			}
		}
	}

	return perimeter
}

func calcSides(region map[string]coord) int {
	yAxisSides := make(map[int][]int)
	xAxisSides := make(map[int][]int)

	xdirs := []coord{
		{1, 0},
		{-1, 0},
	}
	ydirs := []coord{
		{0, 1},
		{0, -1},
	}
	for _, c := range region {
		for _, xdir := range xdirs {
			search := coord{c.x + xdir.x, c.y + xdir.y}
			_, inRegion := region[search.GetHashKey()]

			if !inRegion {
				sideI := c.x * 2
				if (xdir.x) > 0 {
					sideI += xdir.x
				}
				yAxisSides[sideI] = append(yAxisSides[sideI], c.y)
			}
		}
		for _, ydir := range ydirs {
			search := coord{c.x + ydir.x, c.y + ydir.y}
			_, inRegion := region[search.GetHashKey()]

			if !inRegion {
				sideI := c.y * 2
				if (ydir.y) > 0 {
					sideI += ydir.y
				}
				xAxisSides[sideI] = append(xAxisSides[sideI], c.x)
			}
		}
	}

	return numSides(yAxisSides) + numSides(xAxisSides)
}

func numSides(sides map[int][]int) int {
	var sum int
	for _, side := range sides {
		sum++
		// if numbers are continuous, 1 side, else foreach gap +1
		// sort ascending order
		sort.Slice(side, func(a, b int) bool {
			return side[a] < side[b]
		})
		for i := 0; i < len(side)-1; i++ {
			// if gap bigger than 1, new side
			if side[i+1]-side[i] > 1 {
				sum++
			}
		}
	}
	return sum
}
