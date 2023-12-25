package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	startTime := time.Now().UnixMilli()

	edges := readInputs("inputs.txt")

	//toGraphVizFile(edges, "graph.gv", []Edge{})
	exclusions := []Edge{ //infer these from GraphViz visualizations
		{Src: "jbx", Dest: "sml"},
		{Src: "vqj", Dest: "szh"},
		{Src: "zhb", Dest: "vxr"},
	}
	//toGraphVizFile(edges, "graph.gv", exclusions)

	//take any starting nodes from each group visuals for starting point
	res1 := solve(edges, exclusions, "lqp", "jqf")
	fmt.Println("Final Part:", res1)

	fmt.Printf("Duration : %vms", time.Now().UnixMilli()-startTime)
}

type Edge struct {
	Src  string
	Dest string
}

func readInputs(input string) []Edge {
	file, err := os.Open(input)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	results := []Edge{}
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ": ")
		for _, dest := range strings.Split(data[1], " ") {
			results = append(results, Edge{Src: data[0], Dest: dest})
		}
	}

	return results
}

func toGraphVizFile(edges []Edge, output string, exclusions []Edge) {
	file, err := os.Create(output)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("graph G {\n")

	for _, edge := range edges {
		skip := false
		for _, exc := range exclusions {
			if exc.Src == edge.Src && exc.Dest == edge.Dest {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		outputLine := fmt.Sprintf("    %s -- %s;\n", edge.Src, edge.Dest)
		_, err := writer.WriteString(outputLine)
		if err != nil {
			log.Panic(err)
		}
	}

	writer.WriteString("}")
	err = writer.Flush()
	if err != nil {
		log.Panic(err)
	}
}

func solve(edges []Edge, exclusions []Edge, start1 string, start2 string) int {
	allNodes := map[string]map[string]bool{}
	for _, edge := range edges {
		skip := false
		for _, exc := range exclusions {
			if exc.Src == edge.Src && exc.Dest == edge.Dest {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		if curr, ok := allNodes[edge.Src]; ok {
			curr[edge.Dest] = true
			allNodes[edge.Src] = curr
		} else {
			allNodes[edge.Src] = map[string]bool{edge.Dest: true}
		}

		if curr, ok := allNodes[edge.Dest]; ok {
			curr[edge.Src] = true
			allNodes[edge.Dest] = curr
		} else {
			allNodes[edge.Dest] = map[string]bool{edge.Src: true}
		}
	}

	numNodes1 := calculateNodes(allNodes, start1)
	numNodes2 := calculateNodes(allNodes, start2)
	return numNodes1 * numNodes2
}

func calculateNodes(allNodes map[string]map[string]bool, start string) int {
	visited := make(map[string]bool)
	queue := []string{start}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if visited[curr] {
			continue
		}

		visited[curr] = true
		currNode := allNodes[curr]
		for nextNode := range currNode {
			queue = append(queue, nextNode)
		}
	}

	return len(visited)
}
