package main

import (
	"log"
	"os"
	"strings"
)

var paths map[string][]string

type Graph struct {
	Nodes map[string]*Cave
	Edges map[string][]string
}

type Cave struct {
	Big         bool
	Special     bool
	Connections []string
}

func main() {
	f, err := os.ReadFile("./input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(f), "\n")
	graph := Graph{
		Nodes: map[string]*Cave{},
		Edges: map[string][]string{},
	}
	for i := 0; i < len(input); i++ {
		line := strings.Split(input[i], "-")
		graph.Edges[line[0]] = append(graph.Edges[line[0]], line[1])
		graph.Edges[line[1]] = append(graph.Edges[line[1]], line[0])
	}
	for k, v := range graph.Edges {
		graph.Nodes[k] = &Cave{
			Connections: v,
			Big:         isUpper(k),
			Special:     false,
		}
	}

	paths = map[string][]string{}
	for k := range graph.Nodes {
		if k == "start" || k == "end" || graph.Nodes[k].Big {
			continue
		}
		graph.Nodes[k].Special = true
		findPaths("start", []string{}, map[string]int{}, &graph)
		graph.Nodes[k].Special = false
	}

	for _, v := range paths {
		log.Println(v)
	}
	log.Println(len(paths))
}

func findPaths(currentNode string, visited []string, visits map[string]int, graph *Graph) {
	visited = append(visited, currentNode)
	copy := map[string]int{}
	for k, v := range visits {
		copy[k] = v
	}
	copy[currentNode]++
	if currentNode != "end" {
		for _, node := range graph.Nodes[currentNode].Connections {
			if !contains(visited, node) || graph.Nodes[node].Big || (graph.Nodes[node].Special && copy[node] < 2) {
				findPaths(node, visited, copy, graph)
			}
		}
	} else {
		paths[strings.Join(visited, ",")] = visited
	}
}

func isUpper(s string) bool {
	return s == strings.ToUpper(string(s))
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
