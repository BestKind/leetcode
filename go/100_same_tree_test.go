package main

import "testing"

/**
给定两个二叉树，编写一个函数来检验它们是否相同。
如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。
示例 1:
输入:       1         1
          / \       / \
         2   3     2   3
        [1,2,3],   [1,2,3]
输出: true
示例 2:
输入:      1          1
          /           \
         2             2
        [1,2],     [1,null,2]
输出: false
示例 3:
输入:       1         1
          / \       / \
         2   1     1   2
        [1,2,1],   [1,1,2]
输出: false
链接：https://leetcode-cn.com/problems/same-tree
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if nil == p && nil == q {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	if !isSameTree(p.Left, q.Left) {
		return false
	}
	if !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

func TestIsSameTree(t *testing.T) {
	l1 := &TreeNode{Val: 1}
	l1.Left = &TreeNode{Val: 2}
	l1.Right = &TreeNode{Val: 3}
	r1 := &TreeNode{Val: 1}
	r1.Left = &TreeNode{Val: 2}
	r1.Right = &TreeNode{Val: 3}
	t.Log(isSameTree(l1, r1))

	l2 := &TreeNode{Val: 1}
	l1.Left = &TreeNode{Val: 2}
	r2 := &TreeNode{Val: 1}
	r2.Right = &TreeNode{Val: 2}
	t.Log(isSameTree(l2, r2))

	l3 := &TreeNode{Val: 1}
	l3.Left = &TreeNode{Val: 2}
	l3.Right = &TreeNode{Val: 1}
	r3 := &TreeNode{Val: 1}
	r3.Left = &TreeNode{Val: 1}
	r3.Right = &TreeNode{Val: 2}
	t.Log(isSameTree(l3, r3))
}
