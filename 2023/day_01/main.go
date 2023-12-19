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

	res1 := solve("inputs.txt", false)
	fmt.Println("Part 1:", res1)

	res2 := solve("inputs.txt", true)
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string, isAlphanumeric bool) int {
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
		sum += findFirstDigit(line, isAlphanumeric, true)*10 + findFirstDigit(line, isAlphanumeric, false)
	}

	return sum
}

func findFirstDigit(line string, isAlphanumeric bool, scanForward bool) int {
	n := len(line)
	idx, modIdx := 0, 1
	if !scanForward {
		idx = n - 1
		modIdx = -1
	}

	for idx >= 0 && idx < n {
		v := line[idx]
		if v >= '0' && v <= '9' {
			return int(v - '0')
		}

		if isAlphanumeric {
			if digit := getAlphanumericDigit(line, idx); digit != -1 {
				return digit
			}
		}

		idx += modIdx
	}

	return -1
}

func getAlphanumericDigit(line string, idx int) int {
	n := len(line)
	match := func(start int, matcher string) bool {
		end := start + len(matcher)
		return end <= n && line[start:end] == matcher
	}

	switch line[idx] {
	case 'o':
		if match(idx, "one") {
			return 1
		}
	case 't':
		if match(idx, "two") {
			return 2
		}
		if match(idx, "three") {
			return 3
		}
	case 'f':
		if match(idx, "four") {
			return 4
		}
		if match(idx, "five") {
			return 5
		}
	case 's':
		if match(idx, "six") {
			return 6
		}
		if match(idx, "seven") {
			return 7
		}
	case 'e':
		if match(idx, "eight") {
			return 8
		}
	case 'n':
		if match(idx, "nine") {
			return 9
		}
	}

	return -1
}
