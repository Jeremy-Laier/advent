package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	l, r int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rules := []rule{}
	isRules := true
	pages := [][]int{}

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

		if line == "" {
			isRules = false
			continue
		}

		if isRules {
			ruleArr := strings.Split(line, "|")
			l, err := strconv.Atoi(ruleArr[0])
			if err != nil {
				panic(err)
			}

			r, err := strconv.Atoi(ruleArr[1])
			if err != nil {
				panic(err)
			}

			rules = append(rules, rule{l: l, r: r})
		}

		if !isRules {
			page := []int{}

			line := strings.Split(line, ",")
			for _, num := range line {
				digit, err := strconv.Atoi(num)
				if err != nil {
					panic(err)
				}

				page = append(page, digit)
			}
			pages = append(pages, page)
		}
	}

	pt1(rules, pages)
	pt2(rules, pages)

	return
}

func pt1(rules []rule, pages [][]int) {
	sumOfMedians := 0

	for _, page := range pages {
		isValid := true
		for _, rule := range rules {
			lIdx, rIdx := slices.Index(page, rule.l), slices.Index(page, rule.r)
			if lIdx == -1 || rIdx == -1 {
				continue
			}

			if lIdx < rIdx {
				continue
			}

			if lIdx >= rIdx {
				isValid = false
				break
			}
		}
		if isValid {
			sumOfMedians += page[len(page)/2]
		}
	}

	fmt.Println("pt1: ", sumOfMedians)
}

// pt2 find the pages that are not in the correct order
// reorder the incorrect pages
// sum their medians
func pt2(rules []rule, pages [][]int) {
	sumOfMedians := 0

	incorrectPages := [][]int{}
	for _, page := range pages {
		brokenRules := findBrokenRules(page, rules)
		if len(brokenRules) > 0 {
			incorrectPages = append(incorrectPages, page)
		}
	}

	// for each page, continuously swap indeces until there are no broken rules left
	for _, page := range incorrectPages {
		for {
			brokenRules := findBrokenRules(page, rules)
			if len(brokenRules) == 0 {
				break
			}

			flipUpdates(page, brokenRules[0])
		}
	}

	// add up the medians of the fixed, intially wrong pags
	for _, page := range incorrectPages {
		sumOfMedians += page[len(page)/2]
	}

	fmt.Println("pt2: ", sumOfMedians)
}

// flipUpdates swap the elements from the broken rule
// this works because slices in go are passed by reference, we can modify in place and not return
func flipUpdates(page []int, rule rule) {
	lIdx, rIdx := slices.Index(page, rule.l), slices.Index(page, rule.r)

	l := page[lIdx]
	r := page[rIdx]

	page[lIdx] = r
	page[rIdx] = l
}

// findBrokenRules given a page, return a subset of the rules slice with all broken rules
func findBrokenRules(page []int, rules []rule) []rule {
	brokenRules := []rule{}
	for _, rule := range rules {
		lIdx, rIdx := slices.Index(page, rule.l), slices.Index(page, rule.r)
		if lIdx == -1 || rIdx == -1 {
			continue
		}

		if lIdx < rIdx {
			continue
		}

		brokenRules = append(brokenRules, rule)
	}

	return brokenRules
}
