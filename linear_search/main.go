// O(n)

package main

import "log"

func main() {
	var (
		array = []int{
			2, 5, 11, 4, 5, 52, 132,
		}
		elem = 132
	)

	index := searchInt(array, elem)
	if index != -1 {
		log.Printf("index of %d in array: %d", elem, index)
	} else {
		log.Printf("number %d not found in array", elem)
	}
}

func searchInt(array []int, elem int) int {
	for i, v := range array {
		if v == elem {
			return i
		}
	}

	return -1
}
