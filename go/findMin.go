package main

import "fmt"

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
请找出其中最小的元素。
你可以假设数组中不存在重复元素。
示例 1:
输入: [3,4,5,1,2]
输出: 1
示例 2:
输入: [4,5,6,7,0,1,2]
输出: 0
链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
*/

func findMin(nums []int) int {
	low := 0
	high := len(nums) - 1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] {
			high = pivot
		} else if nums[pivot] > nums[high] {
			low = pivot + 1
		} else {
			high--
		}
	}
	return nums[low]
}

func main() {
	numbers := []int{3, 4, 5, 1, 2}
	res := findMin(numbers)
	fmt.Println(res)

	numbers = []int{2, 2, 2, 0, 1}
	res = findMin(numbers)
	fmt.Println(res)

	numbers = []int{4, 5, 6, 7, 0, 1, 2}
	res = findMin(numbers)
	fmt.Println(res)
}
