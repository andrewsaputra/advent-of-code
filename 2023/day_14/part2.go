package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res := solve("inputs.txt")
	fmt.Println(res)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	matrix := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		matrix = append(matrix, []byte(line))
	}

	cache := []CacheData{}
	targetCycles := 1000000000
	for cycle := 0; cycle < targetCycles; cycle++ {
		tiltPlatform(matrix, "north")
		tiltPlatform(matrix, "west")
		tiltPlatform(matrix, "south")
		tiltPlatform(matrix, "east")

		cacheData := toCacheData(matrix)
		for i, data := range cache {
			found := true
			for row, matData := range matrix {
				if string(matData) != data.MatrixStr[row] {
					found = false
					break
				}
			}

			if found {
				modSize := cycle - i
				target := (targetCycles - 1 - i) % modSize
				return cache[i+target].Load
			}
		}

		cache = append(cache, cacheData)
	}

	return calculateLoad(matrix)
}

type CacheData struct {
	MatrixStr []string
	Load      int
}

func toCacheData(matrix [][]byte) CacheData {
	matrixStr := []string{}
	for _, data := range matrix {
		matrixStr = append(matrixStr, string(data))
	}
	return CacheData{matrixStr, calculateLoad(matrix)}
}

func calculateLoad(matrix [][]byte) int {
	total, numRows := 0, len(matrix)
	for row, data := range matrix {
		for _, val := range data {
			if val == 'O' {
				total += numRows - row
			}
		}
	}
	return total
}

func tiltPlatform(matrix [][]byte, direction string) {
	switch direction {
	case "north", "west":
		for row, data := range matrix {
			for col, val := range data {
				if val == 'O' {
					newRow, newCol := findNewSpot(matrix, direction, row, col)
					if row != newRow || col != newCol {
						matrix[row][col] = '.'
						matrix[newRow][newCol] = 'O'
					}
				}
			}
		}
	case "south", "east":
		for row := len(matrix) - 1; row >= 0; row-- {
			for col := len(matrix[0]) - 1; col >= 0; col-- {
				if matrix[row][col] == 'O' {
					newRow, newCol := findNewSpot(matrix, direction, row, col)
					if row != newRow || col != newCol {
						matrix[row][col] = '.'
						matrix[newRow][newCol] = 'O'
					}
				}
			}
		}
	}
}

func findNewSpot(matrix [][]byte, direction string, row int, col int) (int, int) {
	numRows, numCols := len(matrix), len(matrix[0])
	for {
		switch direction {
		case "north":
			row--
			if row < 0 {
				return 0, col
			} else if matrix[row][col] == '#' || matrix[row][col] == 'O' {
				return row + 1, col
			}
		case "south":
			row++
			if row >= numRows {
				return numRows - 1, col
			} else if matrix[row][col] == '#' || matrix[row][col] == 'O' {
				return row - 1, col
			}
		case "west":
			col--
			if col < 0 {
				return row, 0
			} else if matrix[row][col] == '#' || matrix[row][col] == 'O' {
				return row, col + 1
			}
		case "east":
			col++
			if col >= numCols {
				return row, numCols - 1
			} else if matrix[row][col] == '#' || matrix[row][col] == 'O' {
				return row, col - 1
			}
		}
	}
}
