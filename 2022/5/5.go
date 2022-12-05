package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	stacks := make(map[int][]string)
	lineNum := 0
	for scanner.Scan() {
		// first 10 lines are cranes setup
		// parse objects into lists
		line := scanner.Text()

		if lineNum < 8 {
			for i, ch := range line {
				if ch != '[' && ch != ']' && ch != ' ' {
					if i != 1 {
						stacks[i/4] = append(stacks[i/4], string(ch))
					} else {
						stacks[0] = append(stacks[0], string(ch))
					}
				}
			}
			lineNum += 1
		}

		if lineNum == 9 || lineNum == 10 {
			lineNum += 1
			continue
		}


		// 11->EOF is input

	}
	fmt.Printf("%v", stacks)
}
