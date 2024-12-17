package days

import (
	"errors"
	"fmt"
	"os"
)

func Fifteen() {
	fmt.Println("-Day 15-")

	grid := ReadFile("days/inputs/fifteen_grid")
	moves := ReadFile("days/inputs/fifteen_moves")

	fmt.Printf("GPS Sum: %v\n", gpsSum(grid, moves))
	fmt.Printf("Bigger GPS Sum: %v\n", biggerGpsSum(grid, moves))
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
	case "[":
		free := moveBigBoxes(n, coord{n.x + 1, n.y}, g, d)
		if !free {
			return r
		}
		g[n.y][n.x] = g[r.y][r.x]
		g[r.y][r.x] = "."
		return n
	case "]":
		free := moveBigBoxes(coord{n.x - 1, n.y}, n, g, d)
		if !free {
			return r
		}
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
	case "[":
		fallthrough
	case "]":
		fallthrough
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

func moveBigBoxes(l coord, r coord, g [][]string, d coord) bool {
	nl := coord{l.x + d.x, l.y + d.y}
	nr := coord{r.x + d.x, r.y + d.y}
	if d.x != 0 {
		// reuse old move on horiz axis
		var f coord
		if d.x == 1 {
			f = l
		} else {
			f = r
		}
		return moveBoxes(f, g, d)
	}

	// rest for vert axis only

	if !inBounds(nl, len(g[0]), len(g)) || !inBounds(nr, len(g[0]), len(g)) || g[nl.y][nl.x] == "#" || g[nr.y][nr.x] == "#" {
		return false
	}

	// only move if both left and right free
	// this only works on vert axis, for horiz one of these will be the other part of the box
	// TODO: bug here where we're duplicating boxes
	if g[nl.y][nl.x] == "." && g[nr.y][nr.x] == "." {
		g[nl.y][nl.x] = g[l.y][l.x]
		g[nr.y][nr.x] = g[r.y][r.x]
		return true
	}

	// TODO: left move makes right blocked when it otherwise wouldn't be
	var lfree bool
	switch g[nl.y][nl.x] {
	case "]":
		lfree = moveBigBoxes(coord{nl.x - 1, nl.y}, nl, g, d)
	case "[":
		lfree = moveBigBoxes(nl, coord{nl.x + 1, nl.y}, g, d)
	}

	var rfree bool
	switch g[nr.y][nr.x] {
	case "]":
		rfree = moveBigBoxes(coord{nr.x - 1, nr.y}, nr, g, d)
	case "[":
		rfree = moveBigBoxes(nr, coord{nr.x + 1, nr.y}, g, d)
	}

	if !lfree || !rfree {
		return false
	}

	g[nl.y][nl.x] = g[l.y][l.x]
	g[nr.y][nr.x] = g[r.y][r.x]
	return true
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

func biggerGpsSum(grid string, moves string) int {
	g := parseInput2dStrSlice(grid)
	g = embiggenGrid(g)
	printGrid(g)
	robot := FindRobot(g)

	for _, m := range moves {
		dir, err := moveToCoord(string(m))
		if err != nil {
			continue
		}
		robot = moveRobot(robot, g, dir)
		printGrid(g)
	}

	// calc new gps score
	printGrid(g)

	return -1
}

func embiggenGrid(g [][]string) [][]string {
	ng := make([][]string, len(g))
	for i := range ng {
		ng[i] = make([]string, len(g[0])*2)
	}

	nx := 0
	for y := range g {
		for x := range g[y] {
			switch g[y][x] {
			case "#":
				ng[y][nx] = "#"
				ng[y][nx+1] = "#"
			case ".":
				ng[y][nx] = "."
				ng[y][nx+1] = "."
			case "O":
				ng[y][nx] = "["
				ng[y][nx+1] = "]"
			case "@":
				ng[y][nx] = "@"
				ng[y][nx+1] = "."
			default:
				panic("unkown element")
			}
			nx += 2
		}
		nx = 0
	}

	return ng
}

func moveToCoord(m string) (coord, error) {
	var dir coord
	switch m {
	case "^":
		dir = coord{0, -1}
	case "<":
		dir = coord{-1, 0}
	case "v":
		dir = coord{0, 1}
	case ">":
		dir = coord{1, 0}
	default:
		return dir, errors.New("invalid move")
	}
	return dir, nil
}
