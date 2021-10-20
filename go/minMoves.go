package main

import (
	"fmt"
	"sort"
)

/**
 * https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements/
 * 给你一个长度为 n 的整数数组，每次操作将会使 n - 1 个元素增加 1 。返回让数组所有元素相等的最小操作次数。
 * 示例 1：
 * 输入：nums = [1,2,3]
 * 输出：3
 * 解释：
 * 只需要3次操作（注意每次操作会增加两个元素的值）：
 * [1,2,3]  =>  [2,3,3]  =>  [3,4,3]  =>  [4,4,4]
 */
func minMoves(nums []int) int {
	sort.Ints(nums)
	min := nums[0]
	ans := 0
	for _, num := range nums {
		ans += num - min
	}
	return ans
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))
}
