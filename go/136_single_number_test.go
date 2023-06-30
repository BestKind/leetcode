package main

import (
	"fmt"
	"testing"
)

/**
给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
说明：
你的算法应该具有线性时间复杂度。 你可以不使用额外空间来实现吗？
示例 1:
输入: [2,2,1]
输出: 1
示例 2:
输入: [4,1,2,1,2]
输出: 4
https://leetcode-cn.com/explore/interview/card/top-interview-questions/261/before-you-start/1106/
*/

/*
*
解释：利用 异或 操作的特性 相同的数 异或 等于 0, 而 0 与 任何数 异或 等于 任何数
*/
func singleNumber(nums []int) int {
	var tmp int
	for _, v := range nums {
		tmp ^= v
	}

	return tmp
}

func TestSingleNumber(t *testing.T) {
	nums := []int{2, 2, 1}
	res := singleNumber(nums)
	fmt.Println(res)

	nums = []int{4, 1, 2, 1, 2}
	res = singleNumber(nums)
	fmt.Println(res)
}
