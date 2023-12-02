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

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		sum += scanFirstDigit(line, true)*10 + scanFirstDigit(line, false)
	}

	return sum
}

func scanFirstDigit(str string, scanForward bool) int {
	n := len(str)
	var i int
	if !scanForward {
		i = n - 1
	}
	for i >= 0 && i < n {
		v := str[i]
		if v >= '0' && v <= '9' {
			return int(v - '0')
		}

		if scanForward {
			i++
		} else {
			i--
		}
	}

	return -1
}
