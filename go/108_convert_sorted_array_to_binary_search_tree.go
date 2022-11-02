package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
链接来源:https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
 */
func sortedArrayToBST(nums []int) *TreeNode {
	lens := len(nums)
	if lens <= 0 {
		return nil
	}
	mid := lens / 2
	if 0 == lens%2 {
		mid -= 1
	}
	t := &TreeNode{}
	if mid < 0 || mid > lens {
		return t
	}
	t.Val = nums[mid]
	if len(nums[:mid]) > 0 {
		t.Left = sortedArrayToBST(nums[:mid])
	}
	if len(nums[(mid+1):]) > 0 {
		t.Right = sortedArrayToBST(nums[(mid + 1):])
	}

	return t
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	res := sortedArrayToBST(nums)
	travel(res)
	nums = []int{}
	res = sortedArrayToBST(nums)
	travel(res)
}

func travel(node *TreeNode) {
	if nil == node {
		return
	}
	travel(node.Left)
	fmt.Println(node.Val)
	travel(node.Right)
}
