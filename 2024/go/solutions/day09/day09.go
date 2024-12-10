package day09

import (
	"andrewsaputra/adventofcode2024/helper"
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day09.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day09.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int64 {
	input := transcribeInput(path)

	l, r := 0, len(input)-1
	for l < r {
		if input[l] != -1 {
			l++
			continue
		}

		if input[r] == -1 {
			r--
			continue
		}

		input[l], input[r] = input[r], -1
		l++
		r--
	}

	var result int64
	for idx, val := range input {
		if val == -1 {
			break
		}
		result += int64(idx * val)
	}

	return result
}

func solvePart2(path string) int64 {
	input := transcribeInput2(path)

	for r := len(input) - 1; r >= 0; r-- {
		if input[r].ID == -1 {
			continue
		}

		for l := 0; l < r; l++ {
			if input[l].ID != -1 || input[l].Count < input[r].Count {
				continue
			}

			newInput := append([]Item{}, input[:l]...)
			newInput = append(newInput, Item{ID: input[r].ID, Count: input[r].Count})

			input[r].ID = -1
			input[l].Count -= input[r].Count
			if input[l].Count > 0 {
				newInput = append(newInput, input[l:]...)
				r++
			} else {
				newInput = append(newInput, input[l+1:]...)
			}

			input = newInput
			break
		}
	}

	var result int64
	var idx int
	for _, data := range input {
		for i := 0; i < data.Count; i++ {
			if data.ID != -1 {
				result += int64(idx * data.ID)
			}
			idx++
		}
	}

	return result

}

func transcribeInput(path string) []int {
	var result []int
	var id int
	for idx, count := range helper.ReadLines(path)[0] {
		if idx%2 == 0 {
			for i := 0; i < int(count-'0'); i++ {
				result = append(result, id)
			}
			id++
		} else {
			for i := 0; i < int(count-'0'); i++ {
				result = append(result, -1)
			}
		}
	}

	return result
}

type Item struct {
	ID    int
	Count int
}

func transcribeInput2(path string) []Item {
	var result []Item
	var id int
	for idx, count := range helper.ReadLines(path)[0] {
		numCount := int(count - '0')
		if numCount == 0 {
			continue
		}

		if idx%2 == 0 {
			result = append(result, Item{ID: id, Count: int(count - '0')})
			id++
		} else {
			result = append(result, Item{ID: -1, Count: int(count - '0')})
		}
	}

	return result
}
