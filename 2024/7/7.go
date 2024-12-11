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
	defer file.Close()

	equations := map[uint64][]uint64{}
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
		parts := strings.Split(line, ":")
		parts[1] = strings.TrimSpace(parts[1])
		values := strings.Split(parts[1], " ")

		value, err := strconv.ParseUint(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}
		// fmt.Println(parts[0], value)
		equations[value] = Map(values, func(e string) uint64 {
			val, err := strconv.ParseUint(e, 10, 64)
			if err != nil {
				panic(err)
			}
			return val
		})
	}

	pt1(equations)
	pt2(equations)
	return
}

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

func pt1(equations map[uint64][]uint64) {
	var result uint64

	for k, v := range equations {
		slices.Reverse(v)
		answers := _pt1(0, v)
		for _, answer := range answers {
			if answer == k {
				result += k
				break
			}
		}
		slices.Reverse(v)
	}

	fmt.Println("pt1: ", result)
}

func _pt1(n int, list []uint64) []uint64 {
	if n == len(list)-1 {
		return []uint64{list[n]}
	}

	iters := _pt1(n+1, list)

	ret := []uint64{}
	for _, iter := range iters {
		ret = append(ret, []uint64{
			list[n] + iter,
			list[n] * iter,
		}...,
		)
	}

	return ret
}

func pt2(equations map[uint64][]uint64) {
	var result uint64

	for k, v := range equations {
		slices.Reverse(v)
		answers := _pt2(0, v)
		for _, answer := range answers {
			if answer == k {
				result += k
				break
			}
		}
		slices.Reverse(v)
	}

	fmt.Println("pt2: ", result)
}

func _pt2(n int, list []uint64) []uint64 {
	if n == len(list)-1 {
		return []uint64{list[n]}
	}

	iters := _pt2(n+1, list)

	ret := []uint64{}
	for _, iter := range iters {
		concatStr := fmt.Sprintf("%d%d", iter, list[n])
		concat, err := strconv.ParseUint(concatStr, 10, 64)
		if err != nil {
			panic(err)
		}

		ret = append(ret,
			[]uint64{
				list[n] + iter,
				list[n] * iter,
				concat,
			}...,
		)
	}

	return ret
}
