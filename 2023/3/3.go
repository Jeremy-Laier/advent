package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

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
	var engine [][]string
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for _, part := range line {
			row = append(row, string(part))
		}
		engine = append(engine, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pt1(engine))
	return
}

var symbols = []string{"@", "#", "$", "%", "&", "*", "/", "+", "-"}

func isSymbol(s string) bool {
	return slices.Contains(symbols, s)
}

// for each row build
func pt1(engine [][]string) int {
	for i, row := range engine {

		for _, part := range row {

		}
	}

	return 0
}
