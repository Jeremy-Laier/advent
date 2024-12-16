package main

import (
	"fmt"
	"strconv"
)

func main() {
	testStones, stones := []string{"125", "17"}, []string{
		"5",
		"127",
		"680267",
		"39260",
		"0",
		"26",
		"3553",
		"5851995",
	}

	pt1(testStones)
	pt1(stones)

	pt2(testStones)
	pt2(stones)

}

func pt1(stones []string) {
	lastBlink := []string{}
	for _, stone := range stones {
		_pt1(stone, 0, 25, &lastBlink)
	}

	fmt.Println("pt1: ", len(lastBlink))
}

func _pt1(stone string, numBlinks, targetBlinks int, lastBlink *[]string) {
	if numBlinks == targetBlinks {
		// log.Printf("stone: %s, blink: %d", stone, numBlinks)
		*lastBlink = append(*lastBlink, stone)

		return
	}

	num, err := strconv.Atoi(stone)
	if err != nil {
		return
	}

	if num == 0 {
		_pt1("1", numBlinks+1, targetBlinks, lastBlink)
	} else if len(stone)%2 == 0 {
		l, r := "", ""

		for i := range stone {
			if i < len(stone)/2 {
				l += string(stone[i])
			} else {
				r += string(stone[i])
			}
		}
		lNum, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		rNum, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}

		_pt1(fmt.Sprintf("%d", lNum), numBlinks+1, targetBlinks, lastBlink)
		_pt1(fmt.Sprintf("%d", rNum), numBlinks+1, targetBlinks, lastBlink)
	} else {
		_pt1(fmt.Sprintf("%d", num*2024), numBlinks+1, targetBlinks, lastBlink)
	}
}

func pt2(stones []string) {
	var totalStones uint64
	totalStones = 0

	for _, stone := range stones {
		totalStones += _pt22(stone, 1000)
	}

	fmt.Println("pt2: ", totalStones)
}

// instead of passing a pointer around we just make a global cache variable
// ...sorry
// Not really :)
var cache = map[struct {
	s string
	b int
}]uint64{}

func _pt22(stone string, blinksLeft int) uint64 {
	if blinksLeft == 0 {
		return 1
	}

	if v, ok := cache[struct {
		s string
		b int
	}{s: stone, b: blinksLeft}]; ok {
		return v
	}

	num, err := strconv.Atoi(stone)
	if err != nil {
		panic(err)
	}

	totalblinkers := uint64(0)

	if num == 0 {
		totalblinkers += _pt22("1", blinksLeft-1)
	} else if len(stone)%2 == 0 {
		l, r := "", ""

		for i := range stone {
			if i < len(stone)/2 {
				l += string(stone[i])
			} else {
				r += string(stone[i])
			}
		}
		lNum, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		rNum, err := strconv.Atoi(r)
		if err != nil {
			panic(err)
		}

		totalblinkers += _pt22(fmt.Sprintf("%d", rNum), blinksLeft-1)
		totalblinkers += _pt22(fmt.Sprintf("%d", lNum), blinksLeft-1)
	} else {
		totalblinkers += _pt22(fmt.Sprintf("%d", num*2024), blinksLeft-1)
	}

	cache[struct {
		s string
		b int
	}{s: stone, b: blinksLeft}] = totalblinkers

	return totalblinkers
}
