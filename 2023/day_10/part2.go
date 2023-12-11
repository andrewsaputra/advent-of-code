package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

	var start []int
	matrix := [][]byte{}
	for scanner.Scan() {
		line := scanner.Text()
		idx := strings.Index(line, "S")
		if idx != -1 {
			start = []int{len(matrix), idx}
		}

		matrix = append(matrix, []byte(line))
	}

	newMatrix := make([][]byte, len(matrix))
	for row, _ := range newMatrix {
		newMatrix[row] = make([]byte, len(matrix[0]))
		for col := range newMatrix[row] {
			newMatrix[row][col] = '.'
		}
	}

	var res int
	for _, direction := range []string{"north", "south", "east", "west"} {
		res = findLoop(matrix, start, start, direction, 1, nil)
		if res > 0 {
			findLoop(matrix, start, start, direction, 1, newMatrix)
			break
		}
	}

	newMatrix[start[0]][start[1]] = identifyStartPipe(matrix, start)
	total := 0
	for _, data := range newMatrix {
		amount := 0
		isCounting := false

		for _, val := range data {
			if !isCounting {
				if val == '|' || val == 'L' || val == 'J' {
					isCounting = true
					amount = 0
				}
			} else {
				if val == '.' {

					amount++
				} else if val == '|' || val == 'L' || val == 'J' {
					isCounting = false
					total += amount
					amount = 0
				}
			}
		}
	}

	return total
}

func findLoop(matrix [][]byte, start []int, prev []int, direction string, length int, newMatrix [][]byte) int {
	var currPos []int
	switch direction {
	case "north":
		currPos = []int{prev[0] - 1, prev[1]}
	case "south":
		currPos = []int{prev[0] + 1, prev[1]}
	case "west":
		currPos = []int{prev[0], prev[1] - 1}
	case "east":
		currPos = []int{prev[0], prev[1] + 1}
	}

	if currPos[0] == start[0] && currPos[1] == start[1] {
		return length
	}

	m, n := len(matrix), len(matrix[0])
	if currPos[0] < 0 || currPos[1] < 0 || currPos[0] >= m || currPos[1] >= n {
		return -1
	}

	nextDirection := make(map[string]map[byte]string)
	nextDirection["north"] = map[byte]string{'|': "north", '7': "west", 'F': "east"}
	nextDirection["south"] = map[byte]string{'|': "south", 'L': "east", 'J': "west"}
	nextDirection["west"] = map[byte]string{'-': "west", 'L': "north", 'F': "south"}
	nextDirection["east"] = map[byte]string{'-': "east", 'J': "north", '7': "south"}

	val := matrix[currPos[0]][currPos[1]]
	newDirection := nextDirection[direction][val]
	if newDirection == "" {
		return -1
	}

	if newMatrix != nil {
		newMatrix[currPos[0]][currPos[1]] = matrix[currPos[0]][currPos[1]]
	}

	return findLoop(matrix, start, currPos, newDirection, length+1, newMatrix)
}

func identifyStartPipe(matrix [][]byte, start []int) byte {
	m, n := len(matrix), len(matrix[0])
	isValid := func(row int, col int) bool {
		if row < 0 || col < 0 || row >= m || col >= n || matrix[row][col] == '.' {
			return false
		}

		return true
	}

	northValid := isValid(start[0]-1, start[1])
	southValid := isValid(start[0]+1, start[1])
	westValid := isValid(start[0], start[1]-1)
	eastValid := isValid(start[0], start[1]+1)

	switch {
	case northValid && southValid:
		return '|'
	case westValid && eastValid:
		return '-'
	case northValid && eastValid:
		return 'L'
	case northValid && westValid:
		return 'J'
	case southValid && eastValid:
		return 'F'
	case southValid && westValid:
		return '7'
	}

	return 0
}
