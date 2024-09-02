package main

import "log"

func main() {
	log.Println(Task2([]int{1, 3, 5, 7}, []int{2, 4, 6, 8}))
}

func Task2(nums1, nums2 []int) []int {
	var (
		c1   int
		c2   int
		nums = make([]int, 0, len(nums1)+len(nums2))
	)
	for i := 0; i < len(nums1)+len(nums2); i++ {
		if c1 != len(nums1) && (c2 == len(nums2) || nums1[c1] <= nums2[c2]) {
			nums = append(nums, nums1[c1])
			c1++
		} else {
			nums = append(nums, nums2[c2])
			c2++
		}
	}

	return nums
}

// Найти количество пар A и B чисел, при которых B-A>K, nums - отсортирован
func Task1(nums []int, k int) int {
	var (
		res = 0
		end = 0
	)
	for i := 0; i < len(nums); i++ {
		for end < len(nums) && nums[end]-nums[i] <= k {
			end += 1
		}
		res += len(nums) - end
	}

	return res
}
