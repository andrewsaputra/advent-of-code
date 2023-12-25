package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res1 := solvePart1("inputs.txt", 200000000000000, 400000000000000)
	fmt.Println("Part 1:", res1)

	//res2 := solvePart2(grid, start)
	//fmt.Println("Part 2:", res2)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Item struct {
	Position []float64
	Velocity []float64
	M        float64 //formula : y = mx + b
	B        float64
}

func solvePart1(path string, min float64, max float64) int {
	items := parseInput(path)
	numIntersect := 0
	n := len(items)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			intersection := findFutureIntersection(items[i], items[j])
			if intersection == nil {
				continue
			}

			if intersection[0] < min || intersection[0] > max || intersection[1] < min || intersection[1] > max {
				continue
			}

			numIntersect++
		}
	}

	return numIntersect
}

func parseInput(path string) []Item {
	strToNumbers := func(str string) []float64 {
		re := regexp.MustCompile(`-?\d+`)

		numbers := []float64{}
		for _, val := range re.FindAllString(str, -1) {
			num, _ := strconv.ParseFloat(val, 64)
			numbers = append(numbers, num)
		}

		return numbers
	}

	lineEquation := func(position []float64, velocity []float64) (m float64, b float64) {
		if velocity[0] == 0 {
			return math.MaxInt32, math.MaxInt32
		}

		m = velocity[1] / velocity[0]
		b = position[1] - (m * position[0])
		return m, b
	}

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := []Item{}
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, " @ ")
		position := strToNumbers(data[0])
		velocity := strToNumbers(data[1])
		m, b := lineEquation(position, velocity)

		item := Item{
			Position: position,
			Velocity: velocity,
			M:        m,
			B:        b,
		}
		result = append(result, item)
	}

	return result
}

func findFutureIntersection(a, b Item) []float64 {
	if a.M == b.M {
		return nil
	}

	posX := (b.B - a.B) / (a.M - b.M)
	posY := (a.M * posX) + a.B

	var time1, time2 float64
	if a.M == 0 {
		time1 = (posX - a.Position[0]) / a.Velocity[0]
	} else {
		time1 = (posY - a.Position[1]) / a.Velocity[1]
	}

	if b.M == 0 {
		time2 = (posX - b.Position[0]) / b.Velocity[0]
	} else {
		time2 = (posY - b.Position[1]) / b.Velocity[1]
	}

	//fmt.Println(a, b, []float64{posX, posY}, time1, time2)

	if time1 < 0 || time2 < 0 {
		return nil
	}
	return []float64{posX, posY}
}
