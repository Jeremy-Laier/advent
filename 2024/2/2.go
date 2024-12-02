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

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	reports := [][]int{}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		report := []int{}
		line = strings.TrimSpace(line)
		levels := strings.Split(line, " ")
		for _, levelStr := range levels {
			level, err := strconv.Atoi(levelStr)
			if err != nil {
				panic(err)
			}
			report = append(report, level)
		}

		reports = append(reports, report)
	}

	pt1(reports)
	pt2(reports)

	return
}

func pt1(reports [][]int) {
	numSafe := 0

	for _, report := range reports {
		lastLevel := report[0]
		isSafe := true

		for i := 1; i < len(report); i++ {
			if report[i] > lastLevel {
				isSafe = false
				break
			}
			if report[i] == lastLevel {
				isSafe = false
				break
			}

			if lastLevel-report[i] > 3 {
				isSafe = false
				break
			}

			lastLevel = report[i]
		}
		if isSafe {
			numSafe++
		}

		lastLevel = report[0]
		isSafe = true
		for i := 1; i < len(report); i++ {
			if report[i] < lastLevel {
				isSafe = false
				break
			}
			if report[i] == lastLevel {
				isSafe = false
				break
			}

			if report[i]-lastLevel > 3 {
				isSafe = false
				break
			}

			lastLevel = report[i]
		}
		if isSafe {
			numSafe++
		}
	}

	fmt.Println(fmt.Sprintf("pt1: %d", numSafe))
}

func pt2(reports [][]int) {
	numSafe := 0

	for _, report := range reports {
		numSafe += _pt2(report)
	}
	fmt.Println(fmt.Sprintf("pt2: %d", numSafe))
}

func _pt2(report []int) int {
	isValid := isSafe(report)
	if isValid {
		return 1
	}

	slices.Reverse(report)
	isValid = isSafe(report)
	if isValid {
		return 1
	}
	slices.Reverse(report)

	for i := 0; i < len(report); i++ {
		r := []int{}

		for j, level := range report {
			if i == j {
				continue
			}

			r = append(r, level)
		}

		isValid := isSafe(r)
		if isValid {
			return 1
		}

		slices.Reverse(r)
		isValid = isSafe(r)

		if isValid {
			return 1
		}
	}
	return 0
}

func isSafe(list []int) bool {
	for i, elem := range list {
		if i == 0 {
			continue
		}

		if elem == list[i-1] {
			return false
		}

		if elem > list[i-1] {
			return false
		}

		if (list[i-1] - elem) > 3 {
			return false
		}
	}
	return true
}
