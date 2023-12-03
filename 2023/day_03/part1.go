package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	result := solve("inputs.txt")
	fmt.Println("Result:", result)

	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("Duration: %vms\n", duration)
}

type Numbers struct {
	Indexes []int
	Value   int
}

func solve(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	numbers := [][]Numbers{}
	symbols := [][]int{}
	row := 0
	for scanner.Scan() {
		line := scanner.Text()
		numbers = append(numbers, []Numbers{})
		numStartIdx, numEndIdx := -1, -1

		for col, v := range line {
			if v >= '0' && v <= '9' {
				if numStartIdx == -1 {
					numStartIdx = col
				}
			} else {
				if numStartIdx != -1 {
					numEndIdx = col - 1
					num, _ := strconv.Atoi(line[numStartIdx : numEndIdx+1])
					numbers[row] = append(numbers[row], Numbers{
						Indexes: []int{numStartIdx, numEndIdx},
						Value:   num,
					})

					numStartIdx = -1
				}

				if v != '.' {
					symbols = append(symbols, []int{row, col})
				}
			}
		}

		if numStartIdx != -1 {
			numEndIdx = len(line) - 1
			num, _ := strconv.Atoi(line[numStartIdx : numEndIdx+1])
			numbers[row] = append(numbers[row], Numbers{
				Indexes: []int{numStartIdx, numEndIdx},
				Value:   num,
			})
		}

		row++
	}

	sum := 0
	numRows := len(numbers)
	for _, symbol := range symbols {
		scanRows := []int{symbol[0] - 1, symbol[0] + 1}
		scanCols := []int{symbol[1] - 1, symbol[1] + 1}

		row := scanRows[0]
		if row < 0 {
			row = 0
		}

		for row <= scanRows[1] && row < numRows {
			for _, num := range numbers[row] {
				if scanCols[0] <= num.Indexes[1] && scanCols[1] >= num.Indexes[0] {
					sum += num.Value
				}
			}

			row++
		}
	}

	return sum
}
