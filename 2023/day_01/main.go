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

	result := solve("inputs/inputs.txt")
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
	match := func(start int, matcher string) bool {
		end := start + len(matcher)
		return end <= n && str[start:end] == matcher
	}

	var i int
	if !scanForward {
		i = n - 1
	}
	for i >= 0 && i < n {
		v := str[i]
		if v >= '0' && v <= '9' {
			return int(v - '0')
		}

		switch v {
		case 'o':
			if match(i, "one") {
				return 1
			}
		case 't':
			if match(i, "two") {
				return 2
			}
			if match(i, "three") {
				return 3
			}
		case 'f':
			if match(i, "four") {
				return 4
			}
			if match(i, "five") {
				return 5
			}
		case 's':
			if match(i, "six") {
				return 6
			}
			if match(i, "seven") {
				return 7
			}
		case 'e':
			if match(i, "eight") {
				return 8
			}
		case 'n':
			if match(i, "nine") {
				return 9
			}
		}

		if scanForward {
			i++
		} else {
			i--
		}
	}

	return -1
}
