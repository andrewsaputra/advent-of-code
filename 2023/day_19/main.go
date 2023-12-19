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

	res1 := solvePart1("inputs.txt")
	fmt.Println("Part 1:", res1)

	res2 := solvePart2("inputs.txt")
	fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solvePart1(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ruleSet := make(map[string][]Rule)
	isInputString := false
	total := int64(0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isInputString = true
			continue
		}

		if isInputString {
			input := lineToInput(line)
			total += calculateInput(input, ruleSet, "in")
		} else {
			label, rules := lineToRule(line)
			ruleSet[label] = rules
		}
	}

	return total
}

func solvePart2(path string) int64 {
	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	ruleSet := make(map[string][]Rule)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		label, rules := lineToRule(line)
		ruleSet[label] = rules
	}

	inputRange := map[string]Range{
		"x": {Min: 1, Max: 4000},
		"m": {Min: 1, Max: 4000},
		"a": {Min: 1, Max: 4000},
		"s": {Min: 1, Max: 4000},
	}
	return calculateDistinct(inputRange, ruleSet, "in")
}

type Rule struct {
	Char       string
	Comparator string
	MatcherVal int64
	NextLabel  string
}

type Range struct {
	Min int64
	Max int64
}

func lineToRule(line string) (string, []Rule) {
	separator := strings.Index(line, "{")
	label := line[:separator]

	rules := []Rule{}
	for _, data := range strings.Split(line[separator+1:len(line)-1], ",") {
		var newRule Rule

		if len(data) > 1 && (data[1] == '<' || data[1] == '>') {
			destIdx := strings.Index(data, ":")
			matcherVal, _ := strconv.ParseInt(data[2:destIdx], 10, 64)
			newRule = Rule{
				Char:       string(data[0]),
				Comparator: string(data[1]),
				MatcherVal: matcherVal,
				NextLabel:  data[destIdx+1:],
			}
		} else {
			newRule = Rule{NextLabel: data}
		}

		rules = append(rules, newRule)
	}

	return label, rules
}

func lineToInput(line string) map[string]int64 {
	input := make(map[string]int64)
	for _, valueSet := range strings.Split(line[1:len(line)-1], ",") {
		values := strings.Split(valueSet, "=")
		input[values[0]], _ = strconv.ParseInt(values[1], 10, 64)
	}

	return input
}

func calculateInput(input map[string]int64, ruleSet map[string][]Rule, ruleLabel string) int64 {
	if ruleLabel == "R" {
		return 0
	} else if ruleLabel == "A" {
		total := int64(0)
		for _, val := range input {
			total += val
		}
		return total
	}

	for _, rule := range ruleSet[ruleLabel] {
		if rule.Char == "" {
			return calculateInput(input, ruleSet, rule.NextLabel)
		}

		inputVal := input[rule.Char]
		switch rule.Comparator {
		case "<":
			if inputVal < rule.MatcherVal {
				return calculateInput(input, ruleSet, rule.NextLabel)
			}
		case ">":
			if inputVal > rule.MatcherVal {
				return calculateInput(input, ruleSet, rule.NextLabel)
			}
		}
	}

	return 0
}

func copyInputRange(src map[string]Range) map[string]Range {
	dest := make(map[string]Range)
	for key, val := range src {
		dest[key] = val
	}

	return dest
}

func calculateDistinct(inputRange map[string]Range, ruleSet map[string][]Rule, ruleLabel string) int64 {
	if ruleLabel == "R" {
		return 0
	} else if ruleLabel == "A" {
		total := int64(1)
		for _, data := range inputRange {
			tmp := data.Max - data.Min + 1
			total *= tmp
		}
		return total
	}

	total := int64(0)

ForLoop:
	for _, rule := range ruleSet[ruleLabel] {
		if rule.Char == "" {
			total += calculateDistinct(inputRange, ruleSet, rule.NextLabel)
			break
		}

		min, max := inputRange[rule.Char].Min, inputRange[rule.Char].Max
		switch rule.Comparator {
		case "<":
			if min < rule.MatcherVal {
				newMax := rule.MatcherVal - 1
				if newMax >= max {
					total += calculateDistinct(inputRange, ruleSet, rule.NextLabel)
					break ForLoop
				} else {
					subRange := copyInputRange(inputRange)
					subRange[rule.Char] = Range{Min: min, Max: newMax}
					total += calculateDistinct(subRange, ruleSet, rule.NextLabel)
					inputRange[rule.Char] = Range{Min: newMax + 1, Max: max}
				}
			}
		case ">":
			if max > rule.MatcherVal {
				newMin := rule.MatcherVal + 1
				if newMin <= min {
					total += calculateDistinct(inputRange, ruleSet, rule.NextLabel)
					break ForLoop
				} else {
					subRange := copyInputRange(inputRange)
					subRange[rule.Char] = Range{Min: newMin, Max: max}
					total += calculateDistinct(subRange, ruleSet, rule.NextLabel)
					inputRange[rule.Char] = Range{Min: min, Max: newMin - 1}
				}
			}
		}
	}

	return total
}
