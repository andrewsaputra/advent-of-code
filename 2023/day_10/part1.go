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

	var res int
	for _, direction := range []string{"north", "south", "east", "west"} {
		res = findLoop(matrix, start, start, direction, 1)
		if res > 0 {
			break
		}
	}

	return res / 2
}

func findLoop(matrix [][]byte, start []int, prev []int, direction string, length int) int {
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

	return findLoop(matrix, start, currPos, newDirection, length+1)
}
