package main

import "fmt"

/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。
示例 1:
输入: [3,2,3]
输出: 3
示例 2:
输入: [2,2,1,1,1,2,2]
输出: 2
https://leetcode-cn.com/explore/interview/card/top-interview-questions/261/before-you-start/1107/
*/

func majorityElement(nums []int) int {
	length := len(nums)
	if length <= 0 {
		return 0
	}
	res := map[int]int{}
	for _, v := range nums {
		res[v] += 1
	}
	for index, v := range res {
		if v > (length / 2) {
			return index
		}
	}
	return 0
}

func main() {
	nums := []int{3, 2, 3}
	res := majorityElement(nums)
	fmt.Println(res)

	nums = []int{2, 2, 1, 1, 1, 2, 2}
	res = majorityElement(nums)
	fmt.Println(res)
}
