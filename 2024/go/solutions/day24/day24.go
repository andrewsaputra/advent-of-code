package day24

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day24-test.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day24-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	_, operations := toInputs(path)

	for len(operations) > 0 {

	}

	return 0
}

func solvePart2(path string) int {
	return 0
}

type Item struct {
	Wire1     string
	Wire2     string
	Operation string
}

func toInputs(filepath string) (wires map[string]int, operations map[Item]string) {
	wires = make(map[string]int)
	operations = make(map[Item]string)

	var parseItems bool
	for _, line := range helper.ReadLines(filepath) {
		if !parseItems && line[3] != ':' {
			parseItems = true
		}

		if parseItems {
			tmp1 := strings.Split(line, " -> ")
			targetWire := tmp1[1]

			tmp2 := strings.Split(tmp1[0], " ")
			item := Item{
				Wire1:     tmp2[0],
				Wire2:     tmp2[2],
				Operation: tmp2[1],
			}

			operations[item] = targetWire
		} else {
			tmp := strings.Split(line, ": ")
			wireID := tmp[0]
			num, _ := strconv.Atoi(tmp[1])
			wires[wireID] = num
		}
	}

	fmt.Println(wires)
	fmt.Println(operations)

	return
}
