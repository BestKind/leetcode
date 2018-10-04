package main

import "fmt"

/**
	给定一个字符串，找出不含有重复字符的最长子串的长度。

	示例 1:
	输入: "abcabcbb"
	输出: 3
	解释: 无重复字符的最长子串是 "abc"，其长度为 3。

	示例 2:
	输入: "bbbbb"
	输出: 1
	解释: 无重复字符的最长子串是 "b"，其长度为 1。

	示例 3:
	输入: "pwwkew"
	输出: 3
	解释: 无重复字符的最长子串是 "wke"，其长度为 3。
		 请注意，答案必须是一个子串，"pwke" 是一个子序列 而不是子串
 */
func lengthOfLongestSubstring(s string) int {
	location := [256]int{}
	left, maxLen := 0, 0

	for i := range location {
		location[i] = -1
	}

	for i := 0; i < len(s); i++ {
		if location[s[i]] >= left {
			left = location[s[i]] + 1
		} else if i + 1 - left > maxLen {
			maxLen = i + 1 - left
		}
		location[s[i]] = i
	}

	return maxLen
}

func main() {
	str := "abcabcabc"
	res := lengthOfLongestSubstring(str)
	fmt.Println(res)
}
