package main

import "fmt"

/**
 * 给定 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。
 * 在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。
 * 找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
 */
func maxArea(height []int) int {
    i, j := 0, len(height) - 1
    max := 0

    for i < j {
	a, b := height[i], height[j]
	h := min(a, b)

	area := h * (j - i)
	if max < area {
	    max = area
	}

	if a < b {
	    i++
	} else {
	    j--
	}
    }

    return max
}

func min(i int, j int) int {
    if i < j {
	return i
    }
    return j
}

func main() {
    test := []int{1, 3, 6, 4, 3, 5, 6, 7, 8, 9, 7, 5, 4, 3, 2, 1}
    fmt.Println(maxArea(test))
    temp := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
    fmt.Println(maxArea(temp))
}
