package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()
	seeds, allCodecs := parseInputs("inputs.txt")

	res1 := solvePart1(seeds, allCodecs)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2(seeds, allCodecs)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Codec struct {
	SrcStart  int
	SrcEnd    int
	DestStart int
}

func solvePart1(seeds []int, allCodecs [][]Codec) int {
	smallest := math.MaxInt32
	for _, seed := range seeds {
		value := smallestMapping(seed, seed, allCodecs)
		if value < smallest {
			smallest = value
		}
	}
	return smallest
}

func solvePart2(seeds []int, allCodecs [][]Codec) int {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	go func() {
		for i := 0; i < len(seeds); i += 2 {
			minSeed := seeds[i]
			maxSeed := minSeed + seeds[i+1] - 1

			wg.Add(1)
			go func(min int, max int) {
				defer wg.Done()
				ch <- smallestMapping(min, max, allCodecs)
			}(minSeed, maxSeed)
		}

		wg.Wait()
		close(ch)
	}()

	smallest := math.MaxInt32
	for val := range ch {
		if val < smallest {
			smallest = val
		}
	}
	return smallest
}

func parseInputs(path string) ([]int, [][]Codec) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	rows := 0
	allCodecs := [][]Codec{}
	var seeds []int
	var codecs []Codec
	for scanner.Scan() {
		line := scanner.Text()
		if rows == 0 {
			seeds = toNumbers(line[7:])
		} else if line != "" {
			if strings.Contains(line, ":") {
				if len(codecs) > 0 {
					allCodecs = append(allCodecs, codecs)
				}

				codecs = []Codec{}
			} else {
				numbers := toNumbers(line)
				codecs = append(codecs, Codec{
					SrcStart:  numbers[1],
					SrcEnd:    numbers[1] + numbers[2] - 1,
					DestStart: numbers[0],
				})
			}
		}

		rows++
	}

	if len(codecs) > 0 {
		allCodecs = append(allCodecs, codecs)
	}

	return seeds, allCodecs
}

func toNumbers(str string) []int {
	result := []int{}
	for _, s := range strings.Split(str, " ") {
		num, _ := strconv.Atoi(s)
		result = append(result, num)
	}

	return result
}

func smallestMapping(minSeed int, maxSeed int, allCodecs [][]Codec) int {
	smallest := math.MaxInt32
	for original := minSeed; original <= maxSeed; original++ {
		value := original
		for _, codecs := range allCodecs {
			for _, codec := range codecs {
				if value < codec.SrcStart || value > codec.SrcEnd {
					continue
				}

				value = codec.DestStart + (value - codec.SrcStart)
				break
			}
		}

		if value < smallest {
			smallest = value
		}
	}

	return smallest
}
