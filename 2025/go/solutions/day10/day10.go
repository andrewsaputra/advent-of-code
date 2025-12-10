package day10

import (
	"andrewsaputra/adventofcode2025/helper"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day10.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day10.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	var result int
	for _, line := range helper.ReadLines(path) {
		targetLights, buttons, _ := parseInput(line)

		var count int
		stateQueue := []string{strings.Repeat(".", len(targetLights))}
		visited := make(map[string]bool)
		for {
			var found bool
			var newQueue []string
			for _, state := range stateQueue {
				if state == targetLights {
					found = true
					break
				}

				for _, but := range buttons {
					nextState := []byte(state)
					for _, idx := range but {
						if nextState[idx] == '.' {
							nextState[idx] = '#'
						} else {
							nextState[idx] = '.'
						}
					}

					strNextState := string(nextState)
					if visited[strNextState] {
						continue
					}

					visited[strNextState] = true
					newQueue = append(newQueue, strNextState)
				}
			}

			if found {
				break
			}

			stateQueue = newQueue
			count++
		}

		result += count
	}

	return result
}

func solvePart2(path string) int {
	var tmp int
	var result int
	for _, line := range helper.ReadLines(path) {
		tmp++
		fmt.Println("tmp", tmp)

		_, buttons, targetJoltage := parseInput(line)
		var count int
		queue := [][]int{make([]int, len(targetJoltage))}
		visited := make(map[string]bool)

		for {
			var found bool
			var newQueue [][]int

			for _, curr := range queue {
				if slices.Equal(curr, targetJoltage) {
					found = true
					break
				}

				for _, but := range buttons {
					nextState := append([]int{}, curr...)
					valid := true
					for _, idx := range but {
						nextState[idx]++
						if nextState[idx] > targetJoltage[idx] {
							valid = false
							break
						}
					}

					if !valid {
						continue
					}

					key := fmt.Sprint(nextState)
					if visited[key] {
						continue
					}

					visited[key] = true
					newQueue = append(newQueue, nextState)
				}
			}

			if found {
				break
			}

			queue = newQueue
			count++
		}

		result += count
	}

	return result
}

func parseInput(inputLine string) (string, [][]int, []int) {
	re := regexp.MustCompile(`^\[(.+?)\]\s+(.*?)\s+\{(.+?)\}$`)
	matches := re.FindStringSubmatch(inputLine)

	targetLights := matches[1]
	strButtons := matches[2]
	strJoultage := matches[3]

	var buttons [][]int
	re2 := regexp.MustCompile(`\((.*?)\)`)
	for _, match := range re2.FindAllStringSubmatch(strButtons, -1) {
		var button []int
		for _, str := range strings.Split(match[1], ",") {
			num, _ := strconv.Atoi(str)
			button = append(button, num)
		}

		buttons = append(buttons, button)
	}

	var targetJoltage []int
	for _, str := range strings.Split(strJoultage, ",") {
		num, _ := strconv.Atoi(str)
		targetJoltage = append(targetJoltage, num)
	}

	return targetLights, buttons, targetJoltage
}
