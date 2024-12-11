package main

import (
	"advent/lib/higherorder"
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	testDiskMap := []int{}
	diskMap := []int{}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open("input.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

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
			slic := strings.Split(line, "")
			diskMap = higherorder.Map(slic, func(e string) int {
				elem, err := strconv.Atoi(e)
				if err != nil {
					panic(err)
				}
				return elem
			})
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		file, err := os.Open("testinput.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := bufio.NewReader(file)
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		line = strings.TrimSpace(line)
		slic := strings.Split(line, "")
		testDiskMap = higherorder.Map(slic, func(e string) int {
			elem, err := strconv.Atoi(e)
			if err != nil {
				panic(err)
			}
			return elem
		})
	}()

	wg.Wait()

	pt1(testDiskMap)
	pt1(diskMap)

	pt2(testDiskMap)
	pt2(diskMap)
}

func pt1(dm []int) {
	format := getFormat(dm)
	li := generateList(format)
	rearrangeList(li)
	checksum := checksum(li)

	fmt.Println("pt1: ", checksum)
}

func pt2(dm []int) {
	format := getFormat(dm)
	li := generateList(format)
	rearrangeListPt2(li, format)
	checksum := checksum(li)

	fmt.Println("pt2: ", checksum)
}

// Chunk is describing a file block + its following free space. Can have 0 freespace
type Chunk struct {
	size      int
	freespace int
	pos       int
}

func (c Chunk) String() string {
	return fmt.Sprintf("[s: %d, f: %d]", c.size, c.freespace)
}

func getFormat(diskMap []int) map[int]Chunk {
	format := make(map[int]Chunk)

	id := 0
	var chunk Chunk
	for i := 0; i < len(diskMap)+1; i++ {
		if i == len(diskMap) {
			format[id] = chunk
			break
		}

		if i%2 != 0 {
			chunk.freespace = diskMap[i]
			format[id] = chunk

			id++
			chunk = Chunk{}
		}
		if i%2 == 0 {
			chunk.size = diskMap[i]
			chunk.pos = -1
		}
	}

	return format
}

func generateList(format map[int]Chunk) []string {
	list := []string{}

	for i := 0; i < len(format); i++ {
		for range format[i].size {
			list = append(list, fmt.Sprintf("%d", i))
		}

		for range format[i].freespace {
			list = append(list, ".")
		}
	}

	for i := range list {
		if list[i] == "." {
			continue
		}

		id, err := strconv.Atoi(list[i])
		if err != nil {
			panic(err)
		}

		if format[id].pos == -1 {
			format[id] = Chunk{
				freespace: format[id].freespace,
				size:      format[id].size,
				pos:       i,
			}
		}
	}

	return list
}

// rearrangeList takes a slice, rearranges in place
func rearrangeList(list []string) {

	// sliding window
	// iterate backwards until i==j
	j := len(list) - 1
	for i := 0; i < len(list); i++ {
		if i == j {
			break
		}

		iPtr := list[i]
		jPtr := list[j]

		if jPtr == "." {
			j--
			i--
			continue
		}

		if iPtr != "." {
			continue
		}

		list[i] = jPtr
		list[j] = iPtr
		j--
	}
}

func checksum(li []string) int {
	checksum := 0

	for pos := range li {
		if li[pos] == "." {
			continue
		}

		id, err := strconv.Atoi(li[pos])
		if err != nil {
			panic(err)
		}

		checksum += pos * id
	}

	return checksum
}

// rearrangeListPt2 takes a slice, rearranges in place only if block fits
func rearrangeListPt2(list []string, format map[int]Chunk) {

	// iterate starting from first chunk in map adding to the curIdx
	rID := len(format) - 1
	// iterating over the format list, not the list list
	// sliding window over the blocks from 0 to rID
	for {
		if rID < 0 {
			return
		}

		// iterate over entire list
		for i := range list {
			if list[i] == "." {
				continue
			}

			if i >= format[rID].pos {
				break
			}

			lID, err := strconv.Atoi(list[i])
			if err != nil {
				panic(err)
			}

			// if the freespace of the current list elem
			if format[rID].size <= format[lID].freespace {
				// we can fit the right block into the free space
				for l := 0; l < format[rID].size; l++ {
					list[i+l+format[lID].size] = fmt.Sprintf("%d", rID)
					list[format[rID].pos+l] = "."
				}

				format[rID] = Chunk{
					size:      format[rID].size,
					freespace: int(math.Abs(float64(format[lID].freespace - format[rID].size))),
					pos:       i + format[lID].size,
				}

				format[lID] = Chunk{
					size:      format[lID].size,
					freespace: 0,
					pos:       format[lID].pos,
				}
				break
			}
		}
		rID--
	}
}
