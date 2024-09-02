package main

import (
	"log"
	"math"
	"sort"
	"strings"
)

func main() {
	Task5([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
}

// сгруппировать слова с одинаковым набором букв
func Task5(words []string) {
	result := make(map[string][]string)

	for _, word := range words {
		sortedWord := countSortString(word)
		if _, ok := result[sortedWord]; !ok {
			result[sortedWord] = make([]string, 0)
		}
		result[sortedWord] = append(result[sortedWord], word)
	}

	for k, v := range result {
		log.Println(k, v)
	}
}

func countSortString(s string) string {
	dict := make(map[string]int)

	for _, r := range s {
		dict[string(r)] += 1
	}

	keys := make([]string, 0, len(dict))

	for k := range dict {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	word := make([]string, 0)
	for _, key := range keys {
		for i := 0; i < dict[key]; i++ {
			word = append(word, key)
		}
	}

	return strings.Join(word, "")
}

// вывести гистограмму количества букв в строке
func Task4(s string) string {
	maxv := 0
	letters := make(map[string]int)
	for _, v := range s {
		letters[string(v)] += 1
		if letters[string(v)] > maxv {
			maxv = letters[string(v)]
		}
	}

	histogram := ""
	for k, v := range letters {
		histogram += k + " "
		for i := 0; i < v; i++ {
			histogram += "#"
		}
		histogram += "\n"
	}

	return histogram
}

type Cord struct {
	row int
	col int
}

// Даны координаты N ладей, найти количество пар по их ходу на условно бесконечном поле
func Task3(cords []Cord) int {
	var (
		rows = make(map[int]int)
		cols = make(map[int]int)
	)

	for _, v := range cords {
		rows[v.row] += 1
		cols[v.col] += 1
	}

	return countPairs(rows) + countPairs(cols)
}

func countPairs(input map[int]int) int {
	pairs := 0
	for _, v := range input {
		pairs += v - 1
	}

	return pairs
}

// даны 2 числа без ведущих нулей, проверить можно ли получить второе чисто перестановкой чисел в первом
func Task2(num1 int, num2 int) bool {
	var (
		dnum1 = countDigits(num1)
		dnum2 = countDigits(num2)
	)

	for i := 0; i < 10; i++ {
		if dnum1[i] != dnum2[i] {
			return false
		}
	}

	return true
}

func countDigits(num int) []int {
	count := make([]int, 0, 10)
	for i := 0; i < 10; i++ {
		count = append(count, i)
	}

	for num != 0 {
		count[num%10] += 1
		num /= 10
	}

	return count
}

// отсортировать оценки учащегося
func Task1(arr []int) []int {
	maxv := getMax(arr)
	minv := getMin(arr)

	k := (maxv - minv) + 1
	count := make([]int, 0, k)
	for i := 0; i < k; i++ {
		count = append(count, 0)
	}

	for _, v := range arr {
		count[v-minv] += 1
	}

	pos := 0
	for i, v := range count {
		for j := 0; j < v; j++ {
			arr[pos] = i + minv
			pos++
		}
	}

	return arr
}

func getMax(arr []int) int {
	maxv := -math.MaxInt
	for _, v := range arr {
		if v > maxv {
			maxv = v
		}
	}

	return maxv
}

func getMin(arr []int) int {
	minv := math.MaxInt
	for _, v := range arr {
		if v < minv {
			minv = v
		}
	}

	return minv
}
