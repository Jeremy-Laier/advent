package main

import (
	"advent/lib/grid"
	"advent/lib/higherorder"
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

func main() {

	testGrid := grid.Grid[int]{}
	inputGrid := grid.Grid[int]{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

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
			slic := strings.Split(line, "")
			row := higherorder.Map(slic, func(e string) int {
				elem, err := strconv.Atoi(e)
				if err != nil {
					panic(err)
				}
				return elem
			})
			inputGrid.Append(row)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open("testinput.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

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
			slic := strings.Split(line, "")
			row := higherorder.Map(slic, func(e string) int {
				elem, err := strconv.Atoi(e)
				if err != nil {
					panic(err)
				}
				return elem
			})
			testGrid.Append(row)
		}
	}()

	wg.Wait()

	pt1(testGrid)
	pt1(inputGrid)

	pt2(testGrid)
	pt2(inputGrid)
}

func pt1(g grid.Grid[int]) {
	scores := 0

	for y, rows := range g.Li {
		for x := range rows {
			if val, err := g.SafeGet(grid.Coord{X: x, Y: y}); err == nil && val == 0 {
				paths := dfs(grid.Coord{X: x, Y: y}, g, []grid.Coord{})
				paths = higherorder.Filter(paths, func(path []grid.Coord) bool {
					return len(path) == 10
				})

				trailheads := []grid.Coord{}
				for _, path := range paths {
					top := path[0]
					if slices.Contains(trailheads, top) {
						continue
					}
					trailheads = append(trailheads, top)
				}
				scores += len(trailheads)
			}
		}
	}
	fmt.Println("pt1: ", scores)
}

func pt2(g grid.Grid[int]) {
	scores := 0

	for y, rows := range g.Li {
		for x := range rows {
			if val, err := g.SafeGet(grid.Coord{X: x, Y: y}); err == nil && val == 0 {
				paths := dfs(grid.Coord{X: x, Y: y}, g, []grid.Coord{})
				paths = higherorder.Filter(paths, func(path []grid.Coord) bool {
					return len(path) == 10
				})

				// trailheads := []grid.Coord{}
				// for _, path := range paths {
				// 	top := path[0]
				// 	if slices.Contains(trailheads, top) {
				// 		continue
				// 	}
				// 	trailheads = append(trailheads, top)
				// }
				scores += len(paths)
			}
		}
	}
	fmt.Println("pt2: ", scores)
}

// dfs depth first search
//
//	for each node in queue:
//		if node != explored
//			mark node as explored
//			add its edges to the queue
func dfs(point grid.Coord, g grid.Grid[int], explored []grid.Coord) [][]grid.Coord {
	paths := [][]grid.Coord{{point}}
	if slices.Contains(explored, point) {
		return paths
	}

	val := g.Get(point)

	explored = append(explored, point)

	edges := []grid.Coord{
		{
			X: point.X - 1,
			Y: point.Y,
		},
		{
			X: point.X + 1,
			Y: point.Y,
		},

		{
			X: point.X,
			Y: point.Y - 1,
		},

		{
			X: point.X,
			Y: point.Y + 1,
		},
	}

	for _, edge := range edges {
		edgeVal, err := g.SafeGet(edge)
		if err != nil {
			continue
		}

		if edgeVal != val+1 {
			continue
		}
		edgePaths := dfs(edge, g, explored)

		for _, edgePath := range edgePaths {
			edgePath = append(edgePath, point)

			paths = append(paths, edgePath)
		}
	}

	return paths
}
