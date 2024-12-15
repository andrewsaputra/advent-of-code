package main

import (
	"andrewsaputra/adventofcode2024/solutions/day01"
	"andrewsaputra/adventofcode2024/solutions/day02"
	"andrewsaputra/adventofcode2024/solutions/day03"
	"andrewsaputra/adventofcode2024/solutions/day04"
	"andrewsaputra/adventofcode2024/solutions/day05"
	"andrewsaputra/adventofcode2024/solutions/day06"
	"andrewsaputra/adventofcode2024/solutions/day07"
	"andrewsaputra/adventofcode2024/solutions/day08"
	"andrewsaputra/adventofcode2024/solutions/day09"
	"andrewsaputra/adventofcode2024/solutions/day10"
	"andrewsaputra/adventofcode2024/solutions/day11"
	"andrewsaputra/adventofcode2024/solutions/day13"
	"andrewsaputra/adventofcode2024/solutions/day14"
	"andrewsaputra/adventofcode2024/solutions/day15"
	"fmt"
	"os"
	"time"
)

func main() {
	var day string
	if len(os.Args) > 1 {
		day = os.Args[1]
	} else {
		day = "day15"
	}

	startTime := time.Now().UnixMilli()

	switch day {
	case "day01":
		day01.Solve()
	case "day02":
		day02.Solve()
	case "day03":
		day03.Solve()
	case "day04":
		day04.Solve()
	case "day05":
		day05.Solve()
	case "day06":
		day06.Solve()
	case "day07":
		day07.Solve()
	case "day08":
		day08.Solve()
	case "day09":
		day09.Solve()
	case "day10":
		day10.Solve()
	case "day11":
		day11.Solve()
	case "day13":
		day13.Solve()
	case "day14":
		day14.Solve()
	case "day15":
		day15.Solve()
	default:
		fmt.Println("unregistered solution args")
	}

	fmt.Printf("Duration : %vms\n", time.Now().UnixMilli()-startTime)
}
