package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	memories := []string{}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}

		memory := strings.TrimSpace(line)
		memories = append(memories, memory)
	}

	pt1(memories)
	pt2(memories)

}

func pt1(memories []string) {
	result := 0

	for _, memory := range memories {
		result += _pt1(memory)
	}

	fmt.Println(fmt.Sprintf("pt1: %d", result))
}

func _pt1(memory string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	result := 0

	instructions := re.FindAllString(memory, -1)
	for _, instruction := range instructions {
		ins, ok := strings.CutSuffix(instruction, ")")
		if !ok {
			panic("failed to remove )")
		}
		ins, ok = strings.CutPrefix(ins, "mul(")
		if !ok {
			panic("failed to remove mul(")
		}

		pair := strings.Split(ins, ",")
		x, err := strconv.Atoi(pair[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(pair[1])
		if err != nil {
			panic(err)
		}

		result += x * y
	}

	return result
}

func pt2(memories []string) {
	result := 0

	totalMemory := ""
	for _, memory := range memories {
		totalMemory += memory
	}

	result += _pt2(totalMemory)
	fmt.Println(fmt.Sprintf("pt2: %d", result))
}

func _pt2(memory string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	result := 0

	instructions := re.FindAllString(memory, -1)

	doMultiplication := true
	for _, instruction := range instructions {
		if strings.Contains(instruction, "do()") {
			doMultiplication = true
			continue
		}

		if strings.Contains(instruction, "don't") {
			doMultiplication = false
			continue
		}

		if strings.Contains(instruction, "mul") && doMultiplication {
			ins, ok := strings.CutSuffix(instruction, ")")
			if !ok {
				panic("failed to remove )")
			}
			ins, ok = strings.CutPrefix(ins, "mul(")
			if !ok {
				panic("failed to remove mul(")
			}

			pair := strings.Split(ins, ",")
			x, err := strconv.Atoi(pair[0])
			if err != nil {
				panic(err)
			}

			y, err := strconv.Atoi(pair[1])
			if err != nil {
				panic(err)
			}

			result += x * y
		}
	}

	return result
}
