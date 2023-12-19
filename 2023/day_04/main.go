package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res1 := solvePart1("inputs.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs.txt")
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solvePart1(path string) int {
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
		points := calculatePoints(line)
		if points == 0 {
			continue
		}

		sum += int(math.Pow(2, float64(points-1)))
	}

	return sum
}

func solvePart2(path string) int {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	counters := make(map[int]int)
	cardIdx := 0
	for scanner.Scan() {
		line := scanner.Text()
		points := calculatePoints(line)
		counters[cardIdx]++
		for i := 1; i <= points; i++ {
			counters[cardIdx+i] += counters[cardIdx]
		}

		cardIdx++
	}

	sum := 0
	for i := 0; i < cardIdx; i++ {
		sum += counters[i]
	}

	return sum
}

func calculatePoints(line string) int {
	data := strings.Split(line, ": ")
	data = strings.Split(data[1], " | ")

	winners, owned := toNumbers(data[0]), toNumbers(data[1])
	points := 0

	for _, num := range owned {
		for _, num2 := range winners {
			if num == num2 {
				points++
				break
			}
		}
	}

	return points
}

func toNumbers(str string) []int {
	result := []int{}
	for _, s := range strings.Split(str, " ") {
		if s == "" {
			continue
		}

		num, _ := strconv.Atoi(s)
		result = append(result, num)
	}

	return result
}
