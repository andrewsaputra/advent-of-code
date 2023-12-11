package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res := solve("inputs.txt")
	fmt.Println(res)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string) int64 {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	matrix := [][]byte{}
	emptyRows := []int{}
	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, "#") == -1 {
			emptyRows = append(emptyRows, len(matrix))
		}

		matrix = append(matrix, []byte(line))
	}

	galaxies := [][]int{}
	emptyCols := []int{}
	numRows, numCols := len(matrix), len(matrix[0])
	for col := 0; col < numCols; col++ {
		empty := true
		for row := 0; row < numRows; row++ {
			val := matrix[row][col]
			if val == '#' {
				empty = false

				galaxies = append(galaxies, []int{row, col})
			}
		}

		if empty {
			emptyCols = append(emptyCols, col)
		}
	}

	numGalaxies := len(galaxies)
	totalDistance := int64(0)
	for i := 0; i < numGalaxies-1; i++ {
		src := galaxies[i]
		for j := i + 1; j < numGalaxies; j++ {
			dest := galaxies[j]

			totalDistance += calculateDistance(src, dest, emptyRows, emptyCols)
		}
	}

	return totalDistance
}

func calculateDistance(src []int, dest []int, emptyRows []int, emptyCols []int) int64 {
	expansionFactor := int64(999999)
	result := int64(0)
	if src[0] > dest[0] {
		result += int64(src[0] - dest[0])

		for _, v := range emptyRows {
			if v > dest[0] && v < src[0] {
				result += expansionFactor
			}
		}
	} else {
		result += int64(dest[0] - src[0])

		for _, v := range emptyRows {
			if v > src[0] && v < dest[0] {
				result += expansionFactor
			}
		}
	}

	if src[1] > dest[1] {
		result += int64(src[1] - dest[1])

		for _, v := range emptyCols {
			if v > dest[1] && v < src[1] {
				result += expansionFactor
			}
		}
	} else {
		result += int64(dest[1] - src[1])

		for _, v := range emptyCols {
			if v > src[1] && v < dest[1] {
				result += expansionFactor
			}
		}
	}

	return result
}
