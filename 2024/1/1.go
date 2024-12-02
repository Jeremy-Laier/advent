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
	fmt.Println("hello advent 2024!")

	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	leftList, rightList := []string{}, []string{}

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
		pairs := strings.Split(line, "   ")
		leftList = append(leftList, pairs[0])
		rightList = append(rightList, pairs[1])
	}

	pt1(leftList, rightList)
	pt2(leftList, rightList)

	return
}
func pt1(ll, rl []string) {
	slices.Sort(ll)
	slices.Sort(rl)

	distance := 0

	for i := 0; i < len(ll); i++ {
		if ll[i] == "" || rl[i] == "" {
			fmt.Println("breaking because of ", ll[i], rl[i])
			break
		}

		l, err := strconv.Atoi(ll[i])
		if err != nil {
			panic(err)
		}
		r, err := strconv.Atoi(rl[i])
		if err != nil {
			panic(err)
		}
		if l < r {
			distance += r - l
		} else {
			distance += l - r
		}
	}

	fmt.Println("pt 1 answer: ", distance)
	return
}

func pt2(ll, rl []string) {
	score := 0

	cache := map[int]int{}

	for i := 0; i < len(ll); i++ {
		if ll[i] == "" || rl[i] == "" {
			fmt.Println("breaking because of ", ll[i], rl[i])
			break
		}

		l, err := strconv.Atoi(ll[i])
		if err != nil {
			panic(err)
		}

		if cache[l] != 0 {
			// fmt.Println(fmt.Sprintf("l: %d, occurances: %d, score: %d", l, cache[l], cache[l]*l))
			score += cache[l] * l
			continue
		}

		for j := 0; j < len(rl); j++ {
			r, err := strconv.Atoi(rl[j])
			if err != nil {
				panic(err)
			}

			if l == r {
				cache[l] += 1
				continue
			}
		}

		// fmt.Println(fmt.Sprintf("l: %d, occurances: %d, score: %d", l, cache[l], cache[l]*l))
		score += cache[l] * l
	}

	fmt.Println("pt 2 answer: ", score)
	return
}
