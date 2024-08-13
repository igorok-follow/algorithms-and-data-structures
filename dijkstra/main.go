package main

import (
	"log"
	"math"
)

func main() {
	graph := make(map[string]map[string]int)
	graph["a"] = map[string]int{"b": 2, "c": 1}
	graph["b"] = map[string]int{"f": 7}
	graph["c"] = map[string]int{"d": 5, "e": 2}
	graph["d"] = map[string]int{"f": 2}
	graph["e"] = map[string]int{"f": 1}
	graph["f"] = map[string]int{"g": 1}
	graph["g"] = nil

	log.Println(searchShortPath(graph, "a", "g"))
}

func searchShortPath(graph map[string]map[string]int, start, end string) int {
	costs := make(map[string]int)
	processed := make(map[string]struct{}, 0)
	var neighbors map[string]int

	for k := range graph {
		if k != start {
			value, ok := graph[start][k]
			if !ok {
				costs[k] = math.MaxInt32
			} else {
				costs[k] = value
			}
		}
	}

	node, ok := findLowestCostNode(costs, processed)
	for ok {
		cost := costs[node]
		neighbors = graph[node]
		for k, v := range neighbors {
			newCost := cost + v
			if newCost < costs[k] {
				costs[k] = newCost
			}
		}

		processed[node] = struct{}{}
		node, ok = findLowestCostNode(costs, processed)
	}

	return costs[end]
}

func findLowestCostNode(costs map[string]int, processed map[string]struct{}) (string, bool) {
	var (
		node   string
		lowest = math.MaxInt32
	)

	for k, v := range costs {
		if _, ok := processed[k]; !ok && v < lowest {
			lowest = v
			node = k
		}
	}

	if len(node) == 0 {
		return "", false
	}

	return node, true
}
