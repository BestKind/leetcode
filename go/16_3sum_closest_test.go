package main

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
提示：
3 <= nums.length <= 10^3
-10^3 <= nums[i] <= 10^3
-10^4 <= target <= 10^4
链接：https://leetcode-cn.com/problems/3sum-closest
*/

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	var closest = math.MaxInt64

	length := len(nums)
	for l := 0; l < length-2; l++ {
		var (
			m = l + 1
			r = length - 1
		)
		for m < r || closest == target {
			sum := nums[r] + nums[m] + nums[l]
			distance := sum - target
			lastDistance := closest - target
			if math.Abs(float64(distance)) < math.Abs(float64(lastDistance)) {
				closest = sum
			}

			if distance > 0 {
				for cur := r; r > m && nums[cur] == nums[r]; r-- {
				}
			} else if distance < 0 {
				for cur := m; r > m && nums[cur] == nums[m]; m++ {
				}
			} else {
				return sum
			}
		}
	}

	return closest
}

func TestThreeSumClosest(t *testing.T) {
	nums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(nums, 1))
}
