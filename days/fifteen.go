package days

import (
	"fmt"
	"os"
)

func Fifteen() {
	fmt.Println("-Day 15-")

	grid := ReadFile("days/inputs/fifteen_grid")
	moves := ReadFile("days/inputs/fifteen_moves")

	fmt.Printf("GPS Sum: %v\n", gpsSum(grid, moves))
}

func gpsSum(grid string, movements string) int {
	g := parseInput2dStrSlice(grid)
	//move robot and push boxes around.
	robot := FindRobot(g)
	for _, m := range movements {
		var dir coord
		switch string(m) {
		case "^":
			dir = coord{0, -1}
		case "<":
			dir = coord{-1, 0}
		case "v":
			dir = coord{0, 1}
		case ">":
			dir = coord{1, 0}
		default:
			continue
		}

		// move robot + boxes
		robot = moveRobot(robot, g, dir)
	}

	printGrid(g)
	// foreach box, 100 * x * y and sum for total

	return calcGPS(g)
}

func ReadFile(path string) string {
	b, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	return string(b)
}

func FindRobot(g [][]string) coord {
	for y := range g {
		for x := range g[y] {
			if g[y][x] == "@" {
				return coord{x, y}
			}
		}
	}

	panic("no robot")
}

func moveRobot(r coord, g [][]string, d coord) coord {
	n := coord{r.x + d.x, r.y + d.y}
	if !inBounds(n, len(g[0]), len(g)) {
		return r
	}

	switch g[n.y][n.x] {
	case "#":
		return r
	case ".":
		g[n.y][n.x] = g[r.y][r.x]
		g[r.y][r.x] = "."
		return n
	case "O":
		free := moveBoxes(n, g, d)
		if free {
			g[n.y][n.x] = g[r.y][r.x]
			g[r.y][r.x] = "."
			return n
		} else {
			return r
		}
	default:
		panic("unkown obsticle")
	}
}

func moveBoxes(b coord, g [][]string, d coord) bool {
	n := coord{b.x + d.x, b.y + d.y}
	if !inBounds(n, len(g[0]), len(g)) {
		return false
	}

	switch g[n.y][n.x] {
	case "#":
		return false
	case ".":
		g[n.y][n.x] = g[b.y][b.x]
		// fill prev with next in chain
		return true
	case "O":
		free := moveBoxes(n, g, d)
		if free {
			g[n.y][n.x] = g[b.y][b.x]
			return true
		} else {
			return false
		}
	default:
		panic("unkown obstruction")
	}
}

func printGrid(g [][]string) {
	fmt.Println()
	for y := range g {
		for x := range g[y] {
			fmt.Print(g[y][x])
		}
		fmt.Print("\r\n")
	}
	fmt.Println()
}

func calcGPS(g [][]string) int {
	var sum int
	for y := range g {
		for x := range g[y] {
			if g[y][x] != "O" {
				continue
			}
			sum += 100*y + x
		}
	}
	return sum
}
