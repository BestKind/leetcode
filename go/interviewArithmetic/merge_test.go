package main

import (
	"fmt"
	"testing"
)

/**
给你两个有序整数数组 nums1 和 nums2，请你将 nums2 合并到 nums1 中，使 nums1 成为一个有序数组。
说明:
初始化 nums1 和 nums2 的元素数量分别为 m 和 n 。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。
示例:
输入:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3
输出: [1,2,2,3,5,6]
https://leetcode-cn.com/explore/interview/card/top-interview-questions/261/before-you-start/1109/
*/

func merge(nums1 []int, m int, nums2 []int, n int) []int {
	var num []int
	for i, j, k := 0, 0, 0; k < m+n; k++ {
		if i >= m {
			num = append(num, nums2[j])
			j++
			continue
		}
		if j >= n {
			num = append(num, nums1[i])
			i++
			continue
		}
		if nums1[i] <= nums2[j] {
			num = append(num, nums1[i])
			i++
		} else {
			num = append(num, nums2[j])
			j++
		}
	}
	return num
}

func TestMain(t *testing.T) {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	res := merge(nums1, 3, nums2, 3)
	fmt.Println(res)

	nums1 = []int{}
	nums2 = []int{1}
	res = merge(nums1, 0, nums2, 1)
	fmt.Println(res)

	nums1 = []int{1}
	nums2 = []int{}
	res = merge(nums1, 1, nums2, 0)
	fmt.Println(res)

	nums1 = []int{2, 0}
	nums2 = []int{1}
	res = merge(nums1, 1, nums2, 1)
	fmt.Println(res)
}
