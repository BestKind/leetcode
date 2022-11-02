package main

import "fmt"

/**
 * 给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。
 *
 * 你可以假设除了数字 0 之外，这两个数字都不会以零开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
 *
 * 解题思路：主要是把整数换成了另一种表示方法，实现加法
 * 注意进位问题
 */
type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resPre := &ListNode{}
	cur := resPre
	carry := 0

	for  l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}

	return resPre.Next
}

func main() {
	l1 := &ListNode{Val: 2}
	l1 = addNode(l1, 4)
	l1 = addNode(l1, 3)

	l2 := &ListNode{Val: 5}
	l2 = addNode(l2, 6)
	l2 = addNode(l2, 4)

	res := addTwoNumbers(l1, l2)
	showList(res)
}

/**
 * 插入节点（尾插法）
 */
func addNode(l1 *ListNode, val int) *ListNode {
	node := &ListNode{Val: val}
	res := l1
	for res.Next != nil  {
		res = res.Next
	}
	res.Next = node
	return l1
}

func showList(l1 *ListNode) {
	for l1.Next != nil {
		fmt.Println(l1.Val)
		l1 = l1.Next
	}
	fmt.Println(l1.Val)
}
