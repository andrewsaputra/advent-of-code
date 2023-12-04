package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += countPoints(line)
	}

	return sum
}

func countPoints(str string) int {
	cardDetails := strings.Split(str, ": ")[1]
	numbers := strings.Split(cardDetails, " | ")
	winners, owned := strToInts(numbers[0]), strToInts(numbers[1])

	points := 0
	for _, val := range owned {
		for _, val2 := range winners {
			if val == val2 {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
				break
			}
		}
	}

	return points
}

func strToInts(str string) []int {
	result := []int{}
	arr := strings.Split(str, " ")
	for _, s := range arr {
		num, err := strconv.Atoi(s)
		if err != nil {
			continue
		}

		result = append(result, num)
	}
	return result
}
