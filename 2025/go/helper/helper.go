package helper

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		result = append(result, line)
	}
	return result
}

func StringToInts(line string) []int {
	var nums []int
	for _, str := range strings.Fields(line) {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}
	return nums
}

func AbsDiff(v1 int, v2 int) int {
	if v1 > v2 {
		return v1 - v2
	}
	return v2 - v1
}

func ToMatrix(filepath string) [][]byte {
	result := [][]byte{}
	for _, line := range ReadLines(filepath) {
		result = append(result, []byte(line))
	}
	return result
}

func PrintMatrix(matrix [][]byte) {
	for _, rowData := range matrix {
		fmt.Println(string(rowData))
	}
	fmt.Println()
}
