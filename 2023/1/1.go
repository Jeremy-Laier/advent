package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	summation := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		words := []struct {
			index  int
			number string
		}{
			{index: strings.Index(line, "one"), number: "1"},
			{index: strings.Index(line, "two"), number: "2"},
			{index: strings.Index(line, "three"), number: "3"},
			{index: strings.Index(line, "four"), number: "4"},
			{index: strings.Index(line, "five"), number: "5"},
			{index: strings.Index(line, "six"), number: "6"},
			{index: strings.Index(line, "seven"), number: "7"},
			{index: strings.Index(line, "eight"), number: "8"},
			{index: strings.Index(line, "nine"), number: "9"},
			{index: strings.Index(line, "1"), number: "1"},
			{index: strings.Index(line, "2"), number: "2"},
			{index: strings.Index(line, "3"), number: "3"},
			{index: strings.Index(line, "4"), number: "4"},
			{index: strings.Index(line, "5"), number: "5"},
			{index: strings.Index(line, "6"), number: "6"},
			{index: strings.Index(line, "7"), number: "7"},
			{index: strings.Index(line, "8"), number: "8"},
			{index: strings.Index(line, "9"), number: "9"},

			{index: strings.LastIndex(line, "one"), number: "1"},
			{index: strings.LastIndex(line, "two"), number: "2"},
			{index: strings.LastIndex(line, "three"), number: "3"},
			{index: strings.LastIndex(line, "four"), number: "4"},
			{index: strings.LastIndex(line, "five"), number: "5"},
			{index: strings.LastIndex(line, "six"), number: "6"},
			{index: strings.LastIndex(line, "seven"), number: "7"},
			{index: strings.LastIndex(line, "eight"), number: "8"},
			{index: strings.LastIndex(line, "nine"), number: "9"},
			{index: strings.LastIndex(line, "1"), number: "1"},
			{index: strings.LastIndex(line, "2"), number: "2"},
			{index: strings.LastIndex(line, "3"), number: "3"},
			{index: strings.LastIndex(line, "4"), number: "4"},
			{index: strings.LastIndex(line, "5"), number: "5"},
			{index: strings.LastIndex(line, "6"), number: "6"},
			{index: strings.LastIndex(line, "7"), number: "7"},
			{index: strings.LastIndex(line, "8"), number: "8"},
			{index: strings.LastIndex(line, "9"), number: "9"},
		}

		words = slices.DeleteFunc(words, func(s struct {
			index  int
			number string
		}) bool {
			return s.index == -1
		})

		slices.SortFunc(words,
			func(a, b struct {
				index  int
				number string
			}) int {
				return cmp.Compare(a.index, b.index)
			})

		if num, err := strconv.Atoi(words[0].number + words[len(words)-1].number); err != nil {
			panic(err)
		} else {
			summation += num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(summation)

	return
}

func p1(line string) int {

	c1, c2 := 0, 0

	for _, ch := range line {
		if digit, err := strconv.Atoi(string(ch)); err != nil {
			continue
		} else {
			if c1 == 0 {
				c1 = digit
			} else {
				c2 = digit
			}
		}
	}

	if c2 == 0 {
		c2 = c1
	}

	if num, err := strconv.Atoi(fmt.Sprintf("%d%d", c1, c2)); err != nil {
		panic(err)
	} else {
		return num
	}
	return 0
}
