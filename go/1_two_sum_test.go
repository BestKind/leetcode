package main

import (
	"fmt"
	"testing"
)

/**
 * 给定一个整数数组和一个目标值，找出数组中和为目标值的两个数。
 * 你可以假设每个输入只对应一种答案，且同样的元素不能被重复利用。
 *
 * 示例:
 *
 * 给定 nums = [2, 7, 11, 15], target = 9
 *
 * 因为 nums[0] + nums[1] = 2 + 7 = 9
 * 所以返回 [0, 1]
 *
 * 解题思路：a + b = target => a = target - b
 * 所以用一个 map[value] = index 就可以查询到 a 对应的 index
 */
func twoSum(num []int, target int) []int {
	maps := make(map[int]int)

	for i, value := range num {
		member := target - value

		if j, ok := maps[member]; ok {
			return []int{j, i}
		} else {
			maps[value] = i
		}
	}

	return []int{}
}

func TestTwoSum(t *testing.T) {
	num := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(num, target)
	fmt.Println(res)
}
