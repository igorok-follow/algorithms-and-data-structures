package main

import (
	"log"
	"slices"
)

func main() {
	graph := make(map[string][]string)
	graph["a"] = []string{"b", "c"}
	graph["b"] = []string{"f"}
	graph["c"] = []string{"d", "e"}
	graph["d"] = []string{"f"}
	graph["e"] = []string{"f", "g"}
	graph["f"] = []string{"g"}

	log.Println(searchDot(graph, "a", "g"))
}

func searchDot(graph map[string][]string, start, end string) (int, bool) {
	var (
		queue   = []string{start}
		counter int
	)

	var current string
	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]
		if _, ok := graph[current]; !ok {
			graph[current] = []string{}
		}

		if slices.Contains(graph[current], end) {
			return counter, true
		} else {
			queue = append(queue, graph[current]...)
			counter++
		}
	}

	return 0, false
}
