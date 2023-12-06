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

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	maxTime, maxDistance := strToInt(lines[0]), strToInt(lines[1])
	min, max := -1, -1

	//find min valid time
	left, right := 1, maxTime-1
	for left <= right {
		time := left + (right-left)/2
		distance := (maxTime - time) * time

		if distance > maxDistance {
			min = time
			right = time - 1
		} else {
			left = time + 1
		}
	}

	//find max valid time
	left, right = min, maxTime-1
	for left <= right {
		time := left + (right-left)/2
		distance := (maxTime - time) * time

		if distance > maxDistance {
			max = time
			left = time + 1
		} else {
			right = time - 1
		}
	}

	return max - min + 1
}

func strToInt(str string) int {
	num := 0
	for _, v := range str {
		if v >= '0' && v <= '9' {
			num = num*10 + int(v-'0')
		}
	}
	return num
}
