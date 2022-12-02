package main

import (


	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	game := map[string]string {
		"A": "Y",
		"B": "Z",
		"C": "X",

		"X": "A",
		"Y": "B",
		"Z": "C",
	}

	win := map[string]string {
		"A": "Y",
		"B": "Z",
		"C": "X",
	}
	draw := map[string]string {
		"A": "X",
		"B": "Y",
		"C": "Z",
	}
	loss := map[string]string {
		"A": "Z",
		"B": "X",
		"C": "Y",
	}

	scoring := map[string]int {
		"X": 1,
		"Y": 2,
		"Z": 3,
	}

	/**
	x = rock   x beats c
	y = paper  y beats a
	z = scis   z beats a

	a = rock   A loses to Y
	b = paper  B loses to Z
	c = scis   C loses to X
	*/
	score := 0
	score2 := 0
	for scanner.Scan() {
		round := strings.Split(scanner.Text()," ")
		l, r := round[0], round[1]

		if game[l] == r {
			score += 6 + scoring[r]
		} else if l == game[r] {
			score += 3 + scoring[r]
		} else {
			score += 0 + scoring[r]
		}

		if r == "X" { // loss
			score2 += 0 + scoring[loss[l]]
		} else if r == "Y" { // draw
			score2 += 3 + scoring[draw[l]]
		} else { // win
			score2 += 6 + scoring[win[l]]
		}
	}
	fmt.Println("part 1 answer is", score)
	fmt.Println("part 1 answer is", score2)
}
