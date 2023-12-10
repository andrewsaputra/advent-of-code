package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	result := solve("inputs.txt")
	fmt.Println("Result:", result)

	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("Duration: %vms\n", duration)
}

func solve(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instruction, paths := parseInputs(scanner)

	numSteps := []int64{}
	for node, _ := range paths {
		if node[2] == 'A' {
			numSteps = append(numSteps, findSteps(node, paths, instruction))
		}
	}

	return findLCM(numSteps)
}

func parseInputs(scanner *bufio.Scanner) (instruction string, paths map[string][]string) {
	numRows := 0
	paths = make(map[string][]string)
	for scanner.Scan() {
		line := scanner.Text()
		if numRows == 0 {
			instruction = line
		} else if line != "" {
			re := regexp.MustCompile(`([A-Z]+) = \(([A-Z]+), ([A-Z]+)\)`)
			matches := re.FindStringSubmatch(line)
			paths[matches[1]] = []string{matches[2], matches[3]}
		}

		numRows++
	}

	return instruction, paths
}

func findSteps(start string, paths map[string][]string, instruction string) int64 {
	currPos := start
	numSteps := int64(0)
	for currPos[2] != 'Z' {
		for _, dir := range instruction {
			if dir == 'L' {
				currPos = paths[currPos][0]
			} else {
				currPos = paths[currPos][1]
			}

			numSteps++
			if currPos[2] == 'Z' {
				break
			}
		}
	}

	return numSteps

}

func findLCM(numbers []int64) int64 {
	lcm := numbers[0]

	for i := 1; i < len(numbers); i++ {
		lcm = findTwoNumbersLCM(lcm, numbers[i])
	}

	return lcm
}

func findTwoNumbersLCM(a int64, b int64) int64 {
	if a == b {
		return a
	}

	valA, valB := a, b
	for valA != valB {
		if valA < valB {
			valA += a
		} else {
			valB += b
		}
	}

	return valA
}
