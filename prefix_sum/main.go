package main

import "log"

func main() {
	log.Println(Task2([]int{1, 0, 0, 0, 1}))
	// 0  1  1  1  1  2
	//log.Println(Task1([]int{1, 0, 1, 1, 0, 0, 1}, 0, 7))
}

// Дана последовательность N и M запросо. Найти количество нулей на [L:R)
func Task1(nums []int, l, r int) int {
	sum := make([]int, 0, len(nums)+1)
	sum = append(sum, 0)

	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] == 0 {
			sum = append(sum, sum[i-1]+1)
		} else {
			sum = append(sum, sum[i-1])
		}
	}

	return sum[r] - sum[l]
}

// насчитать количество нулевых отрезков
func Task2(nums []int) int {
	ranges := make(map[int]int)
	ranges[0] = 1

	sum := 0
	for _, num := range nums {
		sum += num
		ranges[sum] += 1
	}

	res := 0
	for _, s := range ranges {
		res += s * (s - 1) / 2
	}

	return res

}
