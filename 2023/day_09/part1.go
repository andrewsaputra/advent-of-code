package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	result := solve("inputs.txt")
	fmt.Println("Result:", result)

	duration := time.Now().UnixMilli() - startTime
	fmt.Printf("Duration: %vms\n", duration)
}

func solve(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	sum := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println("line", line)

		val := findNext(line)
		//fmt.Println("val", val)

		sum += val
	}

	return sum
}

func findNext(str string) int64 {
	numbers := []int64{}
	for _, s := range strings.Split(str, " ") {
		num, _ := strconv.Atoi(s)
		numbers = append(numbers, int64(num))
	}

	result := []int64{}
	result = append(result, numbers[len(numbers)-1])

	allZero := false
	for !allZero {
		allZero = true
		newNumbers := []int64{}
		for i := 1; i < len(numbers); i++ {
			num := numbers[i] - numbers[i-1]
			newNumbers = append(newNumbers, num)
			if num != 0 {
				allZero = false
			}
		}

		numbers = newNumbers
		result = append(result, numbers[len(numbers)-1])
	}

	sum := int64(0)
	for _, num := range result {
		sum += num
	}

	return sum

}
