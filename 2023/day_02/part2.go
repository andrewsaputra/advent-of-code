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

	result := solve("inputs-real.txt")
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
		sum += findGamePower(line)
	}

	return sum
}

func findGamePower(line string) int {
	data := strings.Split(line, ": ")
	contentsBatch := strings.Split(data[1], "; ")
	red, green, blue := 0, 0, 0
	for _, batch := range contentsBatch {
		groups := strings.Split(batch, ", ")
		for _, group := range groups {
			balls := strings.Split(group, " ")
			amount, _ := strconv.Atoi(balls[0])
			color := balls[1]

			switch color {
			case "red":
				if amount > red {
					red = amount
				}
			case "green":
				if amount > green {
					green = amount
				}
			case "blue":
				if amount > blue {
					blue = amount
				}
			}
		}
	}

	return red * green * blue
}
