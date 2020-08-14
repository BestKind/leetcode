package main

import (
	"fmt"
	"sort"
)

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：
答案中不可以包含重复的四元组。
示例：
给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。
满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]
链接：https://leetcode-cn.com/problems/4sum
 */

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	if length < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			k, l := j+1, length-1
			for k < l {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				distance := sum - target
				if 0 == distance {
					res = append(res, []int{nums[i], nums[j], nums[k], nums[l]})
				}
				if distance > 0 {
					for cur := l; l > k && nums[l] == nums[cur]; l-- {
					}
				} else {
					for cur := k; k < l && nums[k] == nums[cur]; k++ {
					}
				}
			}
		}
	}

	return res
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	fmt.Println(fourSum(nums, target))
}
