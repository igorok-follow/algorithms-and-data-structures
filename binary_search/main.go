// O(log2n)

package main

import (
	"log"
)

func main() {
	var (
		array = []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15,
		}
		elem = -1
	)

	index := searchInt(array, elem)
	if index != -1 {
		log.Printf("index of %d in array: %d", elem, index)
	} else {
		log.Printf("number %d not found in array", elem)
	}

	index = recursivelySearchInt(array, elem, 0, len(array))
	if index != -1 {
		log.Printf("index of %d in array: %d", elem, index)
	} else {
		log.Printf("number %d not found in array", elem)
	}
}

func searchInt(array []int, elem int) int {
	start := 0
	end := len(array)

	for start <= end {
		middle := (start + end) / 2
		if middle >= len(array) || middle < 0 {
			return -1
		}

		if array[middle] == elem {
			return middle
		}

		if array[middle] > elem {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}

	return -1
}

func recursivelySearchInt(array []int, elem, start, end int) int {
	if start > end {
		return -1
	}

	middle := (start + end) / 2
	if middle >= len(array) || middle < 0 {
		return -1
	}

	if array[middle] == elem {
		return middle
	}

	if array[middle] > elem {
		end = middle - 1
	} else {
		start = middle + 1
	}

	return recursivelySearchInt(array, elem, start, end)
}
