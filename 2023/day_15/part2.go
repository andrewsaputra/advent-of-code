package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	res := solve("inputs.txt")
	fmt.Println(res)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

func solve(path string) int {
	file, _ := os.Open(path)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	contents := []byte{}
	boxes := make([][]Lens, 256)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "," || line == "\r" || line == "\n" {
			if len(contents) > 0 {
				operate(boxes, contents)
				contents = []byte{}
			}
		} else {
			contents = append(contents, line...)
		}
	}

	total := 0
	for i, lenses := range boxes {
		boxMod := 1 + i
		for j, lens := range lenses {
			val := boxMod * (1 + j) * lens.Focal
			total += val
		}
	}

	return total
}

type Lens struct {
	Label string
	Focal int
}

func doHash(content []byte) int {
	res := 0
	for _, v := range content {
		res = ((res + int(v)) * 17) % 256
	}

	return res
}

func operate(boxes [][]Lens, contents []byte) {
	n := len(contents)
	var label string
	var boxNumber, focal int
	if contents[n-1] == '-' {
		label = string(contents[:n-1])
		boxNumber = doHash(contents[:n-1])
		for i, data := range boxes[boxNumber] {
			if data.Label == label {
				boxes[boxNumber] = append(boxes[boxNumber][:i], boxes[boxNumber][i+1:]...)
				break
			}
		}
	} else {
		label = string(contents[:n-2])
		boxNumber = doHash(contents[:n-2])
		focal, _ = strconv.Atoi(string(contents[n-1]))
		exists := false
		for i := range boxes[boxNumber] {
			data := &boxes[boxNumber][i]
			if data.Label == label {
				data.Focal = focal
				exists = true
				break
			}
		}

		if !exists {
			boxes[boxNumber] = append(boxes[boxNumber], Lens{label, focal})
		}
	}
}
