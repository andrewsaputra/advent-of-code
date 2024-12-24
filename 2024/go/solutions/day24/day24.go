package day24

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day24.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day24-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	wires, operations := toInputs(path)

	for len(operations) > 0 {
		for item, targets := range operations {
			if val1, ok := wires[item.Wire1]; ok {
				if val2, ok2 := wires[item.Wire2]; ok2 {
					var res int
					switch item.Operation {
					case "AND":
						if val1 == 1 && val2 == 1 {
							res = 1
						} else {
							res = 0
						}
					case "OR":
						if val1 == 1 || val2 == 1 {
							res = 1
						} else {
							res = 0
						}
					case "XOR":
						if (val1 == 1 && val2 == 0) || (val1 == 0 && val2 == 1) {
							res = 1
						} else {
							res = 0
						}
					}

					for _, targetWire := range targets {
						wires[targetWire] = res
					}

					delete(operations, item)
				}
			}
		}
	}

	return parseResult(wires)
}

func solvePart2(path string) int {
	return 0
}

type Item struct {
	Wire1     string
	Wire2     string
	Operation string
}

type ZItem struct {
	Wire  string
	Order int
	Value int
}

func toInputs(filepath string) (wires map[string]int, operations map[Item][]string) {
	wires = make(map[string]int)
	operations = make(map[Item][]string)

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

			operations[item] = append(operations[item], targetWire)
		} else {
			tmp := strings.Split(line, ": ")
			wireID := tmp[0]
			num, _ := strconv.Atoi(tmp[1])
			wires[wireID] = num
		}
	}

	return
}

func parseResult(wires map[string]int) int64 {
	var zItems []ZItem
	for wire, val := range wires {
		if !strings.HasPrefix(wire, "z") {
			continue
		}

		order, _ := strconv.Atoi(wire[1:])
		zItems = append(zItems, ZItem{Wire: wire, Value: val, Order: order})
	}

	sort.Slice(zItems, func(i, j int) bool {
		return zItems[i].Order < zItems[j].Order
	})

	var result int64
	factor := 1
	for _, item := range zItems {
		result += int64(item.Value * factor)
		factor *= 2
	}

	return result
}
