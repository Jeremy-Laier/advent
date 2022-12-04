package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)
func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	total, total2 := 0, 0

	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), ",")
		l, r := strings.Split(coords[0], "-"), strings.Split(coords[1], "-")

		li1, _ := strconv.Atoi(l[0])
		li2, _ := strconv.Atoi(l[1])

		ri1, _ := strconv.Atoi(r[0])
		ri2, _ := strconv.Atoi(r[1])

		if li1 >= ri1 && li2 <= ri2 || ri1 >= li1 && ri2 <= li2 {
			total += 1
		}

		if li1 >= ri1 && li1 <= ri2 || li2 >= ri1 && li2 <= ri2 {
			total2 += 1
		} else if ri1 >= li1 && ri1 <= li2 || ri2 >= li1 && ri2 <= li2 {
			total2 += 1
		}
	}

	println("part 1 answer: ", total)
	println("part 2 answer: ", total2)
}
