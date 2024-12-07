package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	ID   string
	Sets []Set
}

type Set struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	var games []Game
	for scanner.Scan() {
		line := scanner.Text()

		game := Game{ID: stringArr[0]}

		sets := strings.Split(stringArr[1], ";")

		for _, set := range sets {
			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				green := strings.ReplaceAll(cube, "green", "")
				if len(green) != len(cube) {
					if numOfGreen, err := strconv.Atoi(green); err != nil {
						panic(err)
					} else {
						subset.Green = numOfGreen
					}
				}

				blue := strings.ReplaceAll(cube, "blue", "")
				if len(blue) != len(cube) {
					if numOfBlue, err := strconv.Atoi(blue); err != nil {
						panic(err)
					} else {
						subset.Blue = numOfBlue
					}
				}

				red := strings.ReplaceAll(cube, "red", "")
				if len(red) != len(cube) {
					if numOfRed, err := strconv.Atoi(red); err != nil {
						panic(err)
					} else {
						subset.Red = numOfRed
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return
}
