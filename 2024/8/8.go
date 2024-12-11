package main

import (
	"advent/lib/grid"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"slices"
	"strings"
	"sync"
)

func main() {

	testGrid := grid.Grid[string]{}
	testGrid2 := grid.Grid[string]{}
	inputGrid := grid.Grid[string]{}
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
			row := strings.Split(line, "")
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
			row := strings.Split(line, "")
			testGrid.Append(row)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open("testinput2.txt")
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
			row := strings.Split(line, "")
			testGrid2.Append(row)
		}
	}()

	wg.Wait()

	wg.Wait()

	pt1(testGrid)
	pt1(inputGrid)

	pt2(testGrid)
	pt2(testGrid2)
	pt2(inputGrid)
	return
}

func pt1(g grid.Grid[string]) {
	antennasPerSymbol := map[string][]grid.Coord{}
	for i := range g.Li {
		for j := range g.Li {
			coord := grid.Coord{
				X: j,
				Y: i,
			}
			symbol := g.Get(coord)
			if symbol == "." {
				continue
			}

			antennasPerSymbol[symbol] = append(antennasPerSymbol[symbol], coord)
		}
	}

	antinodes := []grid.Coord{}
	for _, antennas := range antennasPerSymbol {
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}

				nodes := generateAntinodes(antennas[i], antennas[j])
				for _, node := range nodes {
					if node.X < 0 || node.Y < 0 || node.Y >= len(g.Li) || node.X >= len(g.Li) {
						continue
					}

					if !slices.Contains(antinodes, node) {
						antinodes = append(antinodes, node)
					}
				}
			}
		}
	}

	fmt.Println("pt1: ", len(antinodes))
}

func pt2(g grid.Grid[string]) {
	antennasPerSymbol := map[string][]grid.Coord{}
	for i := range g.Li {
		for j := range g.Li {
			coord := grid.Coord{
				X: j,
				Y: i,
			}
			symbol := g.Get(coord)
			if symbol == "." {
				continue
			}

			antennasPerSymbol[symbol] = append(antennasPerSymbol[symbol], coord)
		}
	}

	antinodes := []grid.Coord{}
	for _, antennas := range antennasPerSymbol {
		antiNodesPerSymbol := []grid.Coord{}
		for i := 0; i < len(antennas); i++ {
			for j := 0; j < len(antennas); j++ {
				if i == j {
					continue
				}

				nodes := generateAntinodes(antennas[i], antennas[j])
				for _, node := range nodes {
					if node.X < 0 || node.Y < 0 || node.Y >= len(g.Li) || node.X >= len(g.Li) {
						continue
					}

					if !slices.Contains(antinodes, node) {
						antinodes = append(antinodes, node)
						antiNodesPerSymbol = append(antiNodesPerSymbol, node)

						for {
							l := len(antiNodesPerSymbol)
							nodes2 := generateAntinodes(antennas[i], node)
							for _, node2 := range nodes2 {
								if node2.X < 0 || node2.Y < 0 || node2.Y >= len(g.Li) || node2.X >= len(g.Li) {
									continue
								}
								if !slices.Contains(antinodes, node2) {
									antinodes = append(antinodes, node2)
								}
								if !slices.Contains(antiNodesPerSymbol, node2) {
									antiNodesPerSymbol = append(antiNodesPerSymbol, node2)
								}
							}

							nodes2 = generateAntinodes(antennas[j], node)
							for _, node2 := range nodes2 {
								if node2.X < 0 || node2.Y < 0 || node2.Y >= len(g.Li) || node2.X >= len(g.Li) {
									continue
								}
								if !slices.Contains(antinodes, node2) {
									antinodes = append(antinodes, node2)
								}
								if !slices.Contains(antiNodesPerSymbol, node2) {
									antiNodesPerSymbol = append(antiNodesPerSymbol, node2)
								}
							}
							if l == len(antiNodesPerSymbol) {
								break
							}
						}
					}
				}
			}
		}
	}

	fmt.Println("pt2: ", len(antinodes))
}

func generateAntinodes(antenna1, antenna2 grid.Coord) []grid.Coord {
	if antenna1 == antenna2 {
		return []grid.Coord{}
	}

	// order such that a1 is always either above for up left or
	// below for down left case

	xTravel := math.Abs(float64(antenna1.X - antenna2.X))
	yTravel := math.Abs(float64(antenna2.Y - antenna1.Y))

	// generate 2 anti nodes on either side of antennas
	// must figure out which direction to pick

	// first case
	// # . . .
	// . A . . -> (1, 1)
	// . . A . -> (2, 2)
	// . . . #
	if antenna1.Y < antenna2.Y && antenna1.X < antenna2.X {
		return []grid.Coord{
			{
				X: antenna1.X - int(xTravel),
				Y: antenna1.Y - int(yTravel),
			},
			{
				X: antenna2.X + int(xTravel),
				Y: antenna2.Y + int(yTravel),
			},
		}

	}

	if antenna2.Y < antenna1.Y && antenna2.X < antenna1.X {
		return []grid.Coord{
			{
				X: antenna2.X - int(xTravel),
				Y: antenna2.Y - int(yTravel),
			},
			{
				X: antenna1.X + int(xTravel),
				Y: antenna1.Y + int(yTravel),
			},
		}
	}

	// second case
	// . . . #
	// . . A . -> (2, 1)
	// . A . . -> (1, 2)
	// # . . .
	// antenna 1 => (1, 2)

	if antenna1.Y < antenna2.Y && antenna1.X > antenna2.X {
		return []grid.Coord{
			{
				X: antenna1.X + int(xTravel),
				Y: antenna1.Y - int(yTravel),
			},
			{
				X: antenna2.X - int(xTravel),
				Y: antenna2.Y + int(yTravel),
			},
		}
	}

	if antenna2.Y < antenna1.Y && antenna2.X > antenna1.X {
		return []grid.Coord{
			{
				X: antenna2.X + int(xTravel),
				Y: antenna2.Y - int(yTravel),
			},
			{
				X: antenna1.X - int(xTravel),
				Y: antenna1.Y + int(yTravel),
			},
		}
	}

	return []grid.Coord{}
}
