// O(n*log2n)

package main

import "log"

func main() {
	var array = []int{
		15, 14, 13, 12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 1, 0,
	}

	log.Println(sortInt(array))
}

func sortInt(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	var (
		middleIndex = len(array) / 2
		middle      = array[middleIndex]
		less        []int
		greater     []int
	)
	for i, v := range array {
		if i == middleIndex {
			continue
		}

		if v > middle {
			greater = append(greater, v)
		} else {
			less = append(less, v)
		}
	}

	return append(append(sortInt(less), middle), sortInt(greater)...)
}
