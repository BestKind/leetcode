package main

import (
	"fmt"
	"sort"
	"testing"
)

/**
 * 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？
 * 请你找出所有满足条件且不重复的三元组。
 * 答案中不可以包含重复的三元组。
 * 示例：
 * 	给定数组 nums = [-1, 0, 1, 2, -1, -4]，
 * 	满足要求的三元组集合为：
 * 	[
 * 	 [-1, 0, 1],
 * 	 [-1, -1, 2]
 * 	]
 */

func threeSum(nums []int) [][]int {
	length := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		k := length - 1
		target := -1 * nums[i]
		for j := i + 1; j < length; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for j < k && nums[j]+nums[k] > target {
				k--
			}
			if j == k {
				break
			}
			if nums[j]+nums[k] == target {
				res = append(res, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return res
}

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	res := threeSum(nums)
	fmt.Println(res)

	nums = []int{0, 0, 0, 0, 0, 0}
	res = threeSum(nums)
	fmt.Println(res)
}
