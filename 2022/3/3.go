package main

import (
	"bufio"
	"os"
	"unicode"

	"strings"
)

func main() {

	file, _ := os.Open("input.txt")

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var sacks []string

	for scanner.Scan() {
		sacks = append(sacks, scanner.Text())
	}


	var duplicates []int
	for _, sack := range sacks {
		l, r := sack[0:len(sack)/2], sack[len(sack)/2:]

		for _, ch := range l {
			if strings.Contains(r, string(ch)) {
				subtractor := 0
				if unicode.IsUpper(ch) {
					subtractor = 38
				} else {
					subtractor = 96
				}
				duplicates = append(duplicates, int(ch) - subtractor)
				break
			}
		}
	}

	sum := 0
	for _, dup := range duplicates {
		sum += dup
	}

	println("part 1 answer: ", sum)

	var duplicates2 []int
	for i := 0; i < len(sacks); i+= 3 {
		one, two, three := sacks[i], sacks[i + 1], sacks[i + 2]

		for _, ch := range one {
			if strings.Contains(two, string(ch)) && strings.Contains(three, string(ch)) {
				subtractor := 0
				if unicode.IsUpper(ch) {
					subtractor = 38
				} else {
					subtractor = 96
				}
				duplicates2 = append(duplicates2, int(ch) - subtractor)
				break
			}
		}
	}
	sum = 0
	for _, dup := range duplicates2 {
		sum += dup
	}

	println("part 2 answer: ", sum)

}
