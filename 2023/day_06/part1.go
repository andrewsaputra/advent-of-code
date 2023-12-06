package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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
		fmt.Println("line", line)

		lines = append(lines, line)
	}

	times, distances := strToInts(lines[0]), strToInts(lines[1])
	total := 1
	for i, maxTime := range times {
		distance := distances[i]
		num := 0
		for time := 1; time < maxTime; time++ {
			if (maxTime-time)*time > distance {
				num++
			}
		}

		total *= num
	}

	return total
}

func strToInts(str string) []int {
	re := regexp.MustCompile(`\d+`)

	numbers := []int{}
	for _, val := range re.FindAllString(str, -1) {
		num, _ := strconv.Atoi(val)
		numbers = append(numbers, num)
	}

	return numbers
}
