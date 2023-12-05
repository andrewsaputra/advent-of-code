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

	result := solve("inputs.txt")
	fmt.Println("Result:", result)

	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("Duration: %vms\n", duration)
}

func solve(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	seeds, mappers := parseInputs(scanner)

	ch := make(chan int)
	wg := &sync.WaitGroup{}
	for i := 0; i < len(seeds); i += 2 {
		min, max := seeds[i], seeds[i]+seeds[i+1]-1

		wg.Add(1)
		go getSmallestFinalMapping(ch, wg, min, max, mappers)
	}

	result := make(chan int)
	go func(ch chan int, result chan int) {
		smallest := math.MaxInt32
		for res := range ch {
			if res < smallest {
				smallest = res
			}

		}

		result <- smallest
	}(ch, result)

	wg.Wait()
	close(ch)

	return <-result
}

type Mapper struct {
	SrcStart  int
	Length    int
	DestStart int
}

func parseInputs(scanner *bufio.Scanner) (seeds []int, mappers [][]Mapper) {
	var mapper []Mapper
	numRows := 0
	for scanner.Scan() {
		line := scanner.Text()

		if numRows == 0 {
			seeds = strToInts(line[7:])
		} else {
			if strings.Contains(line, ":") {
				mapper = []Mapper{}
			} else if line == "" {
				if len(mapper) > 0 {
					mappers = append(mappers, mapper)
				}
			} else {
				numbers := strToInts(line)
				mapper = append(mapper, Mapper{
					SrcStart:  numbers[1],
					Length:    numbers[2],
					DestStart: numbers[0],
				})
			}
		}

		numRows++
	}

	if len(mapper) > 0 {
		mappers = append(mappers, mapper)
	}

	return
}

func strToInts(str string) []int {
	result := []int{}
	digits := strings.Split(str, " ")
	for _, digit := range digits {
		num, err := strconv.Atoi(digit)
		if err != nil {
			log.Panic(err)
		}

		result = append(result, num)
	}

	return result
}

func getSmallestFinalMapping(ch chan int, wg *sync.WaitGroup, min int, max int, mappers [][]Mapper) {
	defer wg.Done()

	smallest := math.MaxInt32
	for num := min; num <= max; num++ {
		number := num
		for _, mapper := range mappers {
			for _, data := range mapper {
				diff := number - data.SrcStart
				if diff >= 0 && diff < data.Length {
					number = data.DestStart + diff
					break
				}
			}
		}

		if number < smallest {
			smallest = number
		}
	}

	ch <- smallest
}
