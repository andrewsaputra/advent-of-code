package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()
	symbols, rowNumbers, numRows := parseInputs("inputs.txt")

	res1 := solvePart1(symbols, rowNumbers, numRows)
	fmt.Println("Part 1:", res1)

	res2 := solvePart2(symbols, rowNumbers, numRows)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solvePart1(symbols []Symbol, rowNumbers map[int][]Number, numRows int) int {
	sum := 0
	for _, symbol := range symbols {
		minCol, maxCol := symbol.Col-1, symbol.Col+1
		for row := symbol.Row - 1; row <= symbol.Row+1; row++ {
			if row < 0 || row >= numRows {
				continue
			}

			for _, number := range rowNumbers[row] {
				if number.StartCol <= maxCol && number.EndCol >= minCol {
					sum += number.Value
				}
			}
		}
	}

	return sum
}

func solvePart2(symbols []Symbol, rowNumbers map[int][]Number, numRows int) int {
	sum := 0
	for _, symbol := range symbols {
		if symbol.Value != '*' {
			continue
		}

		minCol, maxCol := symbol.Col-1, symbol.Col+1
		adjacentNumbers := []Number{}
		for row := symbol.Row - 1; row <= symbol.Row+1; row++ {
			if row < 0 || row >= numRows {
				continue
			}

			for _, number := range rowNumbers[row] {
				if number.StartCol <= maxCol && number.EndCol >= minCol {
					adjacentNumbers = append(adjacentNumbers, number)
				}
			}
		}

		if len(adjacentNumbers) == 2 {
			sum += adjacentNumbers[0].Value * adjacentNumbers[1].Value
		}
	}

	return sum
}

type Number struct {
	Value    int
	StartCol int
	EndCol   int
}

type Symbol struct {
	Value byte
	Row   int
	Col   int
}

func parseInputs(path string) ([]Symbol, map[int][]Number, int) {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	grid := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	symbols := gridToSymbols(grid)
	rowNumbers := gridToRowNumbers(grid)

	return symbols, rowNumbers, len(grid)
}

func gridToSymbols(grid [][]byte) []Symbol {
	result := []Symbol{}
	for row := range grid {
		for col, val := range grid[row] {
			if val == '.' || (val >= '0' && val <= '9') {
				continue
			}

			result = append(result, Symbol{val, row, col})
		}
	}

	return result
}

func gridToRowNumbers(grid [][]byte) map[int][]Number {
	isValidNumber := func(row, col int) bool {
		val := grid[row][col]
		if val < '0' || val > '9' {
			return false
		}

		prevCol := col - 1
		if prevCol >= 0 && grid[row][prevCol] >= '0' && grid[row][prevCol] <= '9' {
			return false
		}

		return true
	}

	numCols := len(grid[0])
	result := make(map[int][]Number)
	for row := range grid {
		numbers := []Number{}
		for col, val := range grid[row] {
			if !isValidNumber(row, col) {
				continue
			}

			num := int(val - '0')
			endCol := col
			for mod := col + 1; mod < numCols; mod++ {
				next := grid[row][mod]
				if next < '0' || next > '9' {
					break
				}

				num = num*10 + int(next-'0')
				endCol++
			}

			numbers = append(numbers, Number{num, col, endCol})
		}

		if len(numbers) > 0 {
			result[row] = numbers
		}
	}

	return result
}
