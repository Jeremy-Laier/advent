package main

import (
	_ "embed"
	"fmt"
)

//go:embed testinput.txt
var testInput string

//go:embed input.txt
var input string

func main() {
	wordSearch := [][]string{}

	line := []string{}
	for _, char := range input {
		if char == '\n' {
			wordSearch = append(wordSearch, line)
			line = []string{}
			continue
		}
		line = append(line, string(char))
	}

	pt1(wordSearch)
	pt2(wordSearch)
	return
}

// pt1 find all occurances of XMAS in all directions in the grid
// horizontal, vertical, diagonal
// h backwards, v backwords, d backwards
func pt1(wordSearch [][]string) {
	// counting procedure:
	//  search horizonztal
	//  search vertical
	//  search diagnonal down right
	//  search diagnonal down left
	// match procedure:
	//  grab indeces
	//  match whether the word is XMAS or SAMX
	//    count++
	// final tally procedure:
	//  len(found)

	h, v, dr, dl := 0, 0, 0, 0
	l := len(wordSearch)
	for rowPos, row := range wordSearch {
		// rowPos => what row we are in currently
		for colPos := range row {
			// colPos => what column we are in currently
			// can we search horizontal
			if l-colPos >= 4 {
				word := row[colPos] + row[colPos+1] + row[colPos+2] + row[colPos+3]
				if word == "XMAS" || word == "SAMX" {
					h++
				}
			}

			// can we search vertical
			if l-rowPos >= 4 {
				word := wordSearch[rowPos][colPos] + wordSearch[rowPos+1][colPos] + wordSearch[rowPos+2][colPos] + wordSearch[rowPos+3][colPos]

				if word == "XMAS" || word == "SAMX" {
					v++
				}
			}

			// can we search diagonal down right
			if l-rowPos >= 4 && l-colPos >= 4 {
				word := wordSearch[rowPos][colPos] + wordSearch[rowPos+1][colPos+1] + wordSearch[rowPos+2][colPos+2] + wordSearch[rowPos+3][colPos+3]
				if word == "XMAS" || word == "SAMX" {
					dr++
				}
			}

			// can we search diagonal down left
			if l-rowPos >= 4 && colPos >= 3 {
				word := wordSearch[rowPos][colPos] + wordSearch[rowPos+1][colPos-1] + wordSearch[rowPos+2][colPos-2] + wordSearch[rowPos+3][colPos-3]
				if word == "XMAS" || word == "SAMX" {
					dl++
				}
			}
		}
	}

	fmt.Println("pt1: ", h+v+dr+dl)
}

func pt2(wordSearch [][]string) {
	occurances := 0

	for rowPos, row := range wordSearch {
		for colPos := range row {
			if rowPos >= 1 && rowPos < len(wordSearch)-1 && colPos >= 1 && colPos < len(row)-1 {
				dr := wordSearch[rowPos-1][colPos-1] + wordSearch[rowPos][colPos] + wordSearch[rowPos+1][colPos+1]

				dl := wordSearch[rowPos-1][colPos+1] + wordSearch[rowPos][colPos] + wordSearch[rowPos+1][colPos-1]

				if (dr == "MAS" || dr == "SAM") && (dl == "MAS" || dl == "SAM") {
					occurances++
				}
			}
		}
	}
	fmt.Println("pt2: ", occurances)
}
