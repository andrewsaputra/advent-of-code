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

	sumIds := 0
	for scanner.Scan() {
		line := scanner.Text()
		gameId, valid := checkGame(line)

		if valid {
			sumIds += gameId
		}
	}

	return sumIds
}

func checkGame(line string) (int, bool) {
	maxRed, maxGreen, maxBlue := 12, 13, 14

	data := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Split(data[0], " ")[1])

	contentsBatch := strings.Split(data[1], "; ")
	for _, batch := range contentsBatch {
		groups := strings.Split(batch, ", ")
		for _, group := range groups {
			balls := strings.Split(group, " ")
			amount, _ := strconv.Atoi(balls[0])
			color := balls[1]
			switch color {
			case "red":
				if amount > maxRed {
					return gameId, false
				}
			case "green":
				if amount > maxGreen {
					return gameId, false
				}
			case "blue":
				if amount > maxBlue {
					return gameId, false
				}
			}
		}
	}

	return gameId, true
}
