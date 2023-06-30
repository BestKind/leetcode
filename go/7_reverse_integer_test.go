package main

import (
	"fmt"
	"math"
	"testing"
)

/*
*
给定一个 32 位有符号整数，将整数中的数字进行反转。

示例 1:

输入: 123
输出: 321

	示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。根据这个假设，如果反转后的整数溢出，则返回 0。
*/
func reverse(x int) int {
	sign := 1

	if x < 0 {
		sign = -1
		x = -1 * x
	}

	res := 0
	for x > 0 {
		temp := x % 10
		res = res*10 + temp
		x = x / 10
	}

	res = sign * res

	if res > math.MaxInt32 || res < math.MinInt32 {
		res = 0
	}

	return res
}

func TestReverse(t *testing.T) {
	x := -123
	res := reverse(x)
	fmt.Println(res)
}
