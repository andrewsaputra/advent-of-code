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

func solve(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	instruction, paths := parseInputs(scanner)

	currPos := "AAA"
	numSteps := 0
	for currPos != "ZZZ" {
		for _, dir := range instruction {
			if dir == 'L' {
				currPos = paths[currPos][0]
			} else {
				currPos = paths[currPos][1]
			}

			numSteps++
			if currPos == "ZZZ" {
				break
			}
		}
	}

	return numSteps
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
