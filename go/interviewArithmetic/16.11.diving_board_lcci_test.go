package main

import (
	"fmt"
	"testing"
)

/**
你正在使用一堆木板建造跳水板。有两种类型的木板，其中长度较短的木板长度为shorter，长度较长的木板长度为longer。你必须正好使用k块木板。编写一个方法，生成跳水板所有可能的长度。
返回的长度需要从小到大排列。
示例：
输入：
shorter = 1
longer = 2
k = 3
输出： {3,4,5,6}
提示：
0 < shorter <= longer
0 <= k <= 100000
链接：https://leetcode-cn.com/problems/diving-board-lcci
*/

func divingBoard(shorter int, longer int, k int) []int {
	var res []int
	if k <= 0 {
		return res
	}
	if shorter == longer {
		res = append(res, shorter*k)
		return res
	}
	for i := k; i >= 0; i-- {
		res = append(res, shorter*i+longer*(k-i))
	}
	return res
}

func TestDivingBoard(t *testing.T) {
	res := divingBoard(1, 1, 0)
	fmt.Println(res)

	res = divingBoard(1, 2, 3)
	fmt.Println(res)
}
