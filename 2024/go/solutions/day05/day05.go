package day05

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
	"strconv"
	"strings"
)

func Solve() {
	res1 := solvePart1("inputs/day05.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day05.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	orderingRules, pages := toInputs(path)
	var result int
	for _, page := range pages {
		existing := make(map[string]bool)
		valid := true
		for _, num := range page {
			for keyAfter := range orderingRules[num] {
				if existing[keyAfter] {
					valid = false
					break
				}
			}

			if !valid {
				break
			}
			existing[num] = true
		}

		if valid {
			num, _ := strconv.Atoi(page[len(page)/2])
			result += num
		}
	}

	return result
}

func solvePart2(path string) int {
	orderingRules, pages := toInputs(path)
	var result int
	for _, page := range pages {
		existing := make(map[string]bool)
		valid := true

		for _, num := range page {
			for keyAfter := range orderingRules[num] {
				if existing[keyAfter] {
					valid = false
					break
				}
			}

			if !valid {
				break
			}
			existing[num] = true
		}

		if !valid {
			result += middleOfFixedPage(orderingRules, page)
		}

	}

	return result
}

func toInputs(path string) (orderingRules map[string]map[string]bool, pages [][]string) {
	var isPages bool
	orderingRules = make(map[string]map[string]bool)
	for _, line := range helper.ReadLines(path) {
		if !isPages && strings.Contains(line, ",") {
			isPages = true
		}

		if isPages {
			pages = append(pages, strings.Split(line, ","))
		} else {
			tmp := strings.Split(line, "|")
			before, after := tmp[0], tmp[1]
			if _, ok := orderingRules[before]; ok {
				orderingRules[before][after] = true
			} else {
				orderingRules[before] = map[string]bool{after: true}
			}
		}
	}

	return
}

func middleOfFixedPage(orderingRules map[string]map[string]bool, page []string) int {
	newPage := []string{}
	for _, val := range page {
		insertIdx := -1

		for idx := len(newPage) - 1; idx >= 0; idx-- {
			curr := newPage[idx]
			if orderingRules[val][curr] {
				insertIdx = idx
			}
		}

		if insertIdx == -1 {
			newPage = append(newPage, val)
		} else {
			tmp := append([]string{}, newPage[:insertIdx]...)
			tmp = append(tmp, val)
			tmp = append(tmp, newPage[insertIdx:]...)
			newPage = tmp
		}
	}

	num, _ := strconv.Atoi(newPage[len(newPage)/2])
	return num

}
