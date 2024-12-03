package main

import (
	"andrewsaputra/adventofcode2024/solutions/day01"
	"andrewsaputra/adventofcode2024/solutions/day02"
	"andrewsaputra/adventofcode2024/solutions/day03"
	"fmt"
	"os"
	"time"
)

func main() {
	var day string
	if len(os.Args) > 1 {
		day = os.Args[1]
	} else {
		day = "day03"
	}

	startTime := time.Now().UnixMilli()

	switch day {
	case "day01":
		day01.Solve()
	case "day02":
		day02.Solve()
	case "day03":
		day03.Solve()
	default:
		fmt.Println("unregistered solution args")
	}

	fmt.Printf("Duration : %vms\n", time.Now().UnixMilli()-startTime)
}
