package day21

import (
	"fmt"
)

func Solve() {
	res1 := solvePart1("inputs/day21-test.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs/day21-test.txt")
	fmt.Println("Part 2:", res2)
}

func solvePart1(path string) int {
	//inputs := helper.ReadLines(path)
	return 0
}

func solvePart2(path string) int {
	return 0
}

type Pos struct {
	Row int
	Col int
}

var (
	numpad = [][]byte{
		{'7', '8', '9'},
		{'4', '5', '6'},
		{'1', '2', '3'},
		{' ', '0', 'A'},
	}
	numpadPos = Pos{Row: 3, Col: 2}

	dpad = [][]byte{
		{' ', '^', 'A'},
		{'<', 'v', '>'},
	}
	dpadPos = Pos{Row: 0, Col: 2}
)

func shortestNumpad(input string) []string {
	//drow := []int{-1, 0, 1, 0}
	//dcol := []int{0, 1, 0, -1}

	return nil
}
