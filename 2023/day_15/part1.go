package main

import (
	"bufio"
	"fmt"
	"os"
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
	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "," || line == "\r" || line == "\n" {
			if len(contents) > 0 {
				total += doHash(contents)
				contents = []byte{}
			}
		} else {
			contents = append(contents, line...)
		}
	}

	return total
}

func doHash(content []byte) int {
	res := 0
	for _, v := range content {
		res = ((res + int(v)) * 17) % 256
	}

	return res
}
