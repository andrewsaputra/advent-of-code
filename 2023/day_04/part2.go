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

	copies := make(map[int]int)
	numCards := 0
	for scanner.Scan() {
		line := scanner.Text()
		parseCard(copies, numCards, line)
		numCards++
	}

	sum := 0
	for idx := 0; idx < numCards; idx++ {
		curr := copies[idx]
		sum += 1 + curr
	}

	return sum
}

func parseCard(copies map[int]int, currIdx int, str string) {
	cardDetails := strings.Split(str, ": ")[1]
	numbers := strings.Split(cardDetails, " | ")
	winners, owned := strToInts(numbers[0]), strToInts(numbers[1])

	curr := copies[currIdx]
	numMatch := 0
	for _, val := range owned {
		for _, val2 := range winners {
			if val == val2 {
				numMatch++
				break
			}
		}
	}

	for idx := 1; idx <= numMatch; idx++ {
		targetIdx := currIdx + idx
		target := copies[targetIdx]
		target += 1 + curr
		copies[targetIdx] = target
	}
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
