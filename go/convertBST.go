package main


/**
给定一个二叉搜索树（Binary Search Tree），把它转换成为累加树（Greater Tree)，使得每个节点的值是原来的节点值加上所有大于它的节点值之和。
例如：
输入: 原始二叉搜索树:
              5
            /   \
           2     13
输出: 转换为累加树:
             18
            /   \
          20     13
链接：https://leetcode-cn.com/problems/convert-bst-to-greater-tree
 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
解题思路：
	1. 因为是二叉树，所以 右子树节点值 > 根节点值 > 左子树节点值
	2. 由此可使用中序遍历(反过来用)
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Right)
			sum += node.Val
			node.Val = sum
			dfs(node.Left)
		}
	}
	dfs(root)
	return root
}

func main() {
	
}
