package main

import (
	"andrewsaputra/adventofcode2025/solutions/day01"
	"fmt"
	"os"
	"time"
)

func main() {
	var day string
	if len(os.Args) > 1 {
		day = os.Args[1]
	} else {
		day = "day01"
	}

	startTime := time.Now().UnixMilli()

	switch day {
	case "day01":
		day01.Solve()
	default:
		fmt.Println("unregistered solution args")
	}

	fmt.Printf("Duration : %vms\n", time.Now().UnixMilli()-startTime)
}
