// O(n)

package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	//var (
	//	array = []int{
	//		8, 1, 4, 2, 5, 11, 4, 5, 52, 132,
	//	}
	//	elem = 132
	//)
	//
	//index := searchInt(array, elem)
	//if index != -1 {
	//	log.Printf("index of %d in array: %d", elem, index)
	//} else {
	//	log.Printf("number %d not found in array", elem)
	//}
	//
	//value := findMinEvenX(array)
	//if value != -1 {
	//	log.Printf("value %d is min even in array", value)
	//} else {
	//	log.Printf("min even not found in array")
	//}
	//
	//log.Println(isleFlood([]int{3, 1, 4, 3, 5, 1, 5, 1, 1, 3, 1}))

	log.Println(rle("AAAABBBBBCCCCDDD222111AAA"))
}

func searchInt(array []int, elem int) int {
	for i, v := range array {
		if v == elem {
			return i
		}
	}

	return -1
}

// найти минимальное четное чило, если нет то -1
func findMinEvenX(array []int) int {
	minV := -1
	for i := 0; i < len(array); i++ {
		if (array[i] < minV || minV == -1) && array[i]%2 == 0 {
			minV = array[i]
		}
	}

	return minV
}

func minLenWords(array []string) string {
	if len(array) == 0 {
		return ""
	}

	minLen := len(array[0])
	for i := 1; i < len(array); i++ {
		if len(array[i]) < minLen {
			minLen = len(array[i])
		}
	}

	res := make([]string, 0)
	for _, v := range array {
		if len(v) == minLen {
			res = append(res, v)
		}
	}

	return strings.Join(array, " ")
}

func isleFlood(array []int) int {
	highpos := 0
	for i, v := range array {
		if v > highpos {
			highpos = i
		}
	}

	result := 0
	currentHighPos := 0
	for _, v := range array[:highpos] {
		if v > currentHighPos {
			currentHighPos = v
		}

		result += currentHighPos - v
	}
	currentHighPos = 0
	for i := len(array) - 1; i > highpos; i-- {
		if array[i] > currentHighPos {
			currentHighPos = array[i]
		}

		result += currentHighPos - array[i]
	}

	return result
}

func rle(input string) string {
	arr := []rune(input)

	lastSymbol := arr[0]
	lastPos := 0
	result := make([]string, 0)

	for i, v := range arr {
		if v != lastSymbol {
			result = append(result, string(lastSymbol)+strconv.Itoa(i-lastPos))
			lastPos = i
			lastSymbol = v
		}
	}
	result = append(result, string(arr[lastPos])+strconv.Itoa(len(arr)-lastPos))

	return strings.Join(result, " ")
}
