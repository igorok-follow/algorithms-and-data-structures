package main

import (
	"log"
	"strings"
)

func main() {
	log.Println(Task1([]int{1, 2, 3, 4, 5, 5}, 9))
	Task2([]string{
		"hello",
		"world",
	}, "helo wold")
}

// дана последовательность чисел длиной N и число X
// найти 2 различных A и B сумма которых будет равна X
// если нет то вернуть 0, 0
// O(N)
func Task1(arr []int, x int) (int, int) {
	set := make(map[int]struct{})

	for _, num := range arr {
		if _, ok := set[x-num]; ok {
			return x - num, num
		} else {
			set[num] = struct{}{}
		}
	}

	return 0, 0
}

// Дан текст и словарь, в словах в тексте может быть пропущена одна буква,
// сопоставить их к правильным словам и сказать есть ли они
// O(NK^2+M) где N - словарь K каждое слово M проход по тексту
func Task2(dictionary []string, text string) {
	fullDictionary := make(map[string]struct{})
	for _, v := range dictionary {
		fullDictionary[v] = struct{}{}
	}

	for _, word := range dictionary {
		for i := 1; i < len(word); i++ {
			w := word[:i] + word[i+1:]
			if _, ok := fullDictionary[w]; !ok {
				fullDictionary[w] = struct{}{}
			}
		}
	}

	for _, word := range strings.Split(text, " ") {
		_, ok := fullDictionary[word]
		log.Println(word, ok)
	}
}
