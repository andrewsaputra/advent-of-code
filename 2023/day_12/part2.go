package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res := solve("inputs.txt")
	fmt.Println(res)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string) int64 {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	total := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		total += countSolutions(line)
	}
	return total
}

func countSolutions(str string) int64 {
	data := strings.Split(str, " ")
	records := data[0]

	nums := strings.Split(data[1], ",")
	damages := []int{}
	for _, num := range nums {
		val, _ := strconv.Atoi(num)
		damages = append(damages, val)
	}

	cache := map[string]int64{}
	records, damages = unfold(records, damages, 5)
	res := countSolutionsHelper(records, damages, cache)
	//fmt.Println(records, damages, " = ", res)
	return res
}

func countSolutionsHelper(record string, damages []int, cache map[string]int64) int64 {
	if len(damages) == 0 {
		for i := 0; i < len(record); i++ {
			if record[i] == '#' {
				return 0
			}
		}

		return 1
	}

	cacheKey := fmt.Sprintf("%s-%v", record, damages)
	if val, ok := cache[cacheKey]; ok {
		return val
	}

	maxSize := len(damages) - 1
	for _, v := range damages {
		maxSize += v
	}

	total := int64(0)
	limit := len(record) - maxSize
	for i := 0; i <= limit; i++ {
		if !isValid(record, i, damages[0]) {
			continue
		}

		nextIdx := i + damages[0] + 1
		nextRecord := ""
		if nextIdx < len(record) {
			nextRecord = record[nextIdx:]
		}

		//fmt.Printf("record:%s, i:%d, damages:%v\n", record, i, damages)
		total += countSolutionsHelper(nextRecord, damages[1:], cache)
	}

	cache[cacheKey] = total
	return total
}

func isValid(record string, recordIdx int, damageSize int) bool {
	afterIdx := recordIdx + damageSize
	if afterIdx < len(record) && record[afterIdx] == '#' {
		return false
	}

	for i := recordIdx - 1; i >= 0; i-- {
		if record[i] == '#' {
			return false
		}
	}

	for i := recordIdx; i < afterIdx; i++ {
		if record[i] == '.' {
			return false
		}
	}

	return true
}

func unfold(record string, damages []int, modifier int) (newRecord string, newDamages []int) {
	newDamages = []int{}
	sb := strings.Builder{}
	for i := 0; i < modifier; i++ {
		sb.WriteString(record)
		if i < modifier-1 {
			sb.WriteString("?")
		}

		newDamages = append(newDamages, damages...)
	}

	newRecord = sb.String()
	return
}
