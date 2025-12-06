package main

import (
	"andrewsaputra/adventofcode2025/solutions/day01"
	"andrewsaputra/adventofcode2025/solutions/day02"
	"andrewsaputra/adventofcode2025/solutions/day03"
	"andrewsaputra/adventofcode2025/solutions/day04"
	"andrewsaputra/adventofcode2025/solutions/day05"
	"andrewsaputra/adventofcode2025/solutions/day06"
	"fmt"
	"os"
	"time"
)

func main() {
	var day string
	if len(os.Args) > 1 {
		day = os.Args[1]
	} else {
		day = "day06"
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
	default:
		fmt.Println("unregistered solution args")
	}

	fmt.Printf("Duration : %vms\n", time.Now().UnixMilli()-startTime)
}
