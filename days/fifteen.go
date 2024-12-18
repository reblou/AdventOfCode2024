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
	case "O":
		free := moveBoxes(n, g, d)
		if free {
			g[n.y][n.x] = g[r.y][r.x]
			g[r.y][r.x] = "."
			return n
		} else {
			return r
		}
	}

	if g[n.y][n.x] == "[" || g[n.y][n.x] == "]" {
		if d.x != 0 {
			// reuse old move on horiz axis
			free := moveBoxes(n, g, d)
			if !free {
				return r
			}
			g[n.y][n.x] = g[r.y][r.x]
			g[r.y][r.x] = "."
			return n
		} else {
			// else moving[b.y][b.x+1]g vertically, need extra logic
			b := n
			if g[n.y][n.x] == "]" {
				b = coord{n.x - 1, n.y}
			}
			stack := make([]coord, 0)
			free := moveBigBoxes(b, g, d, &stack)
			if !free {
				return r
			}
			// remove duplicates that could occur from a box being pushed twice
			stack = removeDupeBoxes(stack)
			// go through stack and move
			for _, b := range stack {
				g[b.y+d.y][b.x+d.x] = g[b.y][b.x]
				g[b.y+d.y][b.x+d.x+1] = g[b.y][b.x+1]
				g[b.y][b.x] = "."
				g[b.y][b.x+1] = "."
			}

			g[n.y][n.x] = g[r.y][r.x]
			g[r.y][r.x] = "."
			return n

		}
	}

	panic("unknown obsticle")
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

func moveBigBoxes(l coord, g [][]string, d coord, stack *[]coord) bool {
	nl := coord{l.x + d.x, l.y + d.y}
	nr := coord{l.x + d.x + 1, l.y + d.y}

	if !inBounds(nl, len(g[0]), len(g)) || !inBounds(nr, len(g[0]), len(g)) || g[nl.y][nl.x] == "#" || g[nr.y][nr.x] == "#" {
		return false
	}

	if g[nl.y][nl.x] == "." && g[nr.y][nr.x] == "." {
		*stack = append(*stack, l)
		return true
	}

	if g[nl.y][nl.x] == "[" {
		// box directly above us, don't need to check alternatives
		free := moveBigBoxes(nl, g, d, stack)
		if !free {
			return false
		}

		*stack = append(*stack, l)
		return true
	}

	lfree := true
	rfree := true
	if g[nl.y][nl.x] == "]" {
		// box to the left,
		lfree = moveBigBoxes(coord{nl.x - 1, nl.y}, g, d, stack)
	}

	if g[nr.y][nr.x] == "[" {
		// box ofset right
		rfree = moveBigBoxes(nr, g, d, stack)
	}

	if !lfree || !rfree {
		// if either side has blockers
		return false
	}

	*stack = append(*stack, l)
	return true
}

func printGrid(g [][]string) {
	var boxes int
	fmt.Println()
	for y := range g {
		for x := range g[y] {
			if g[y][x] == "O" || g[y][x] == "[" {
				boxes++
			}
			fmt.Print(g[y][x])
		}
		fmt.Print("\r\n")
	}
	fmt.Printf("--%v boxes--\n", boxes)
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

func calcBigGPS(g [][]string) int {
	var sum int
	for y := range g {
		for x := range g[y] {
			if g[y][x] != "[" {
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
	robot := FindRobot(g)

	for _, m := range moves {
		dir, err := moveToCoord(string(m))
		if err != nil {
			// non directional character, newline etc.
			continue
		}
		robot = moveRobot(robot, g, dir)
	}

	return calcBigGPS(g)
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

func removeDupeBoxes(stack []coord) []coord {
	var out []coord
	seen := make(map[string]bool)

	for _, s := range stack {
		if seen[s.GetHashKey()] {
			continue
		}

		seen[s.GetHashKey()] = true
		out = append(out, s)
	}

	return out
}
