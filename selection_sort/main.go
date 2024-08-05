// O(n*n)

package main

import "log"

func main() {
	var array = []int{
		15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 1, 0,
	}

	log.Println(sortInt(array))
}

func sortInt(array []int) []int {
	var minIndex int
	for i := range array {
		minIndex = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[i] {
				minIndex = j
			}
		}

		array[i], array[minIndex] = array[minIndex], array[i]
	}

	return array
}
