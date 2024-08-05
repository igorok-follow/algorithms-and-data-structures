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
	for range array {
		for j := 0; j < len(array)-1; j++ {
			if array[j+1] < array[j] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}

	return array
}
