package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := [][]string{}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		line = strings.TrimSpace(line)
		row := strings.Split(line, "")
		row = Map(row, func(elem string) string {
			return strings.TrimSpace(elem)
		})
		grid = append(grid, row)
	}
	// grid1 := slices.Clone(grid)
	grid2 := slices.Clone(grid)
	// pt1(grid1)
	pt2(grid2)
	return
}

func pt1(grid [][]string) {
	positions := 0

	var x, y, x1, y1 int
	var direction string

	var rowL, colL int
	rowL = len(grid)
	for i, row := range grid {
		colL = len(row)
		for j, pos := range row {
			if string(pos) == "^" {
				direction = string(pos)
				x, y = j, i
				break
			}

			if string(pos) == "v" {
				direction = string(pos)
				x, y = j, i
				break
			}

			if string(pos) == "<" {
				direction = string(pos)
				x, y = j, i
				break
			}

			if string(pos) == ">" {
				direction = string(pos)
				x, y = j, i
				break
			}
		}
	}

	// move the guard until he walks out of bounds
	for {
		if x < 0 || x >= colL || y < 0 || y >= rowL {
			break
		}
		tile := grid[y][x]
		if tile == "#" {
			x, y = x1, y1
			switch direction {
			case "^":
				direction = ">"

			case ">":
				direction = "v"

			case "v":
				direction = "<"

			case "<":
				direction = "^"
			}
		}
		x1, y1 = x, y
		x, y = moveGuard(grid, x, y, direction)
		grid[y1][x1] = "X"
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[j][i] == "X" {
				positions++
			}
		}
	}
	fmt.Println("pt1: ", positions)
}

func moveGuard(grid [][]string, x, y int, direction string) (int, int) {
	switch direction {
	case "^":
		return x, y - 1

	case ">":
		return x + 1, y

	case "v":
		return x, y + 1

	case "<":
		return x - 1, y
	default:
		return -1, -1
	}
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func pt2(grid [][]string) {
	positions := 0

	var x, y, x1, y1 int
	var direction string

	var rowL, colL int
	rowL = len(grid)
	for i, row := range grid {
		colL = len(row)
		for j, pos := range row {
			if string(pos) == "^" {
				direction = string(pos)
				x, y = j, i
				break
			}

			if string(pos) == "v" {
				direction = string(pos)
				x, y = j, i
				break
			}

			if string(pos) == "<" {
				direction = string(pos)
				x, y = j, i
				break
			}

			if string(pos) == ">" {
				direction = string(pos)
				x, y = j, i
				break
			}
		}
	}

	x0, y0 := x, y
	d0 := direction
	// stack of path
	// pop, push, clear
	path := Set[struct {
		x, y int
	}]{}
	// move the guard until he walks out of bounds
	for {
		if x < 0 || x >= colL || y < 0 || y >= rowL {
			break
		}
		tile := grid[y][x]
		if tile == "#" {
			x, y = x1, y1
			switch direction {
			case "^":
				direction = ">"

			case ">":
				direction = "v"

			case "v":
				direction = "<"

			case "<":
				direction = "^"
			}
		}
		x1, y1 = x, y
		x, y = moveGuard(grid, x, y, direction)
		grid[y1][x1] = "X"
		path.Append(struct {
			x int
			y int
		}{x: x1, y: y1})
	}

	for _, coord := range path.set {
		grid[coord.y][coord.x] = "O"

		if willLoop(grid, x0, y0, d0) {
			positions++
		}

		grid[coord.y][coord.x] = "X"
	}

	fmt.Println("pt2: ", positions)
}

type Set[T comparable] struct {
	set []T
}

func (s *Set[T]) Append(e T) bool {
	if slices.Contains(s.set, e) {
		return false
	}

	s.set = append(s.set, e)

	return true
}

func (s *Set[T]) Len() int {
	return len(s.set)
}

func (s *Set[T]) Contains(e T) bool {
	return slices.Contains(s.set, e)
}

func (s *Set[T]) Reverse() {
	slices.Reverse(s.set)
	return
}

func _willLoop(grid [][]string, x, y int, direction string) bool {
	var x1, y1 int

	// move the guard until he walks out of bounds
	// or, if he loops return true
	loops := 0
	for {
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid) {
			return false
		}
		// once we stop adding new we know we have started looping
		if loops == 10000 {
			return true
		}
		tile := grid[y][x]
		if tile == "#" || tile == "O" {
			x, y = x1, y1
			switch direction {
			case "^":
				direction = ">"

			case ">":
				direction = "v"

			case "v":
				direction = "<"

			case "<":
				direction = "^"
			}
		}
		x1, y1 = x, y
		x, y = moveGuard(grid, x, y, direction)
		loops++
	}
}
func willLoop(grid [][]string, x, y int, d string) bool {
	return _willLoop(grid, x, y, d)
}
