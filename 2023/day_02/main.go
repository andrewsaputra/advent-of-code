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

	maxBalls := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := lineToGame(line)

		if isValidGame(game, maxBalls) {
			sum += game.Id
		}
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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		game := lineToGame(line)
		sum += calculatePower(game)
	}

	return sum
}

type Game struct {
	Id      int
	Records []map[string]int
}

func lineToGame(line string) Game {
	data := strings.Split(line, ": ")
	gameId, _ := strconv.Atoi(strings.Split(data[0], " ")[1])

	records := []map[string]int{}
	for _, roundStr := range strings.Split(data[1], "; ") {
		record := make(map[string]int)
		for _, ballStr := range strings.Split(roundStr, ", ") {
			ball := strings.Split(ballStr, " ")
			amount, _ := strconv.Atoi(ball[0])
			color := ball[1]
			record[color] = amount
		}

		records = append(records, record)
	}

	return Game{gameId, records}
}

func isValidGame(game Game, maxBalls map[string]int) bool {
	for _, record := range game.Records {
		for color, amount := range record {
			if amount > maxBalls[color] {
				return false
			}
		}
	}

	return true
}

func calculatePower(game Game) int {
	maxBalls := map[string]int{}
	for _, record := range game.Records {
		for color, amount := range record {
			currAmount, ok := maxBalls[color]
			if !ok || amount > currAmount {
				maxBalls[color] = amount
			}
		}
	}

	total := 1
	for _, amount := range maxBalls {
		total *= amount
	}

	return total
}
