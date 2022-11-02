package main

import (
	"fmt"
	"time"
)

/**
 * 编写一个函数来查找字符串数组中的最长公共前缀
 * 如果不存在公共前缀，返回空字符串 ""
 * 示例 1:
 * 	输入: ["flower","flow","flight"]
 * 	输出: "fl"
 * 示例 2:
 * 	输入: ["dog","racecar","car"]
 * 	输出: ""
 * 	解释: 输入不存在公共前缀。
 * 说明:
 * 	所有输入只包含小写字母 a-z
 * 方法二 优于 方法一
 */

/** 方法一 */
func longestCommonPrefix(strs []string) string {
	cnt := len(strs)
	if cnt <= 0 {
		return ""
	}
	str := ""
	flag := true
	for site, b := range strs[0] {
		for i := 1; i < cnt; i++ {
			if len(strs[i]) < site+1 || uint8(b) != strs[i][site] {
				flag = false
				break
			}
		}
		if !flag {
			break
		}
		str += string(b)
	}
	return str
}

/** 方法二 */
func longestCommonPrefix1(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 1; i < count; i++ {
		prefix = lcp(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func lcp(str1, str2 string) string {
	length := myMin(len(str1), len(str2))
	index := 0
	for index < length && str1[index] == str2[index] {
		index++
	}
	return str1[:index]
}

func myMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	start := time.Now()
	var strs = []string{"flower", "flow", "flight"}
	res := longestCommonPrefix(strs)
	fmt.Println(res)

	strs = []string{"dog", "racecar", "car"}
	res = longestCommonPrefix(strs)
	fmt.Println(res)

	strs = []string{"aa", "a"}
	res = longestCommonPrefix(strs)
	fmt.Println(res)
	fmt.Println("end: ", time.Now().Sub(start))

	start = time.Now()
	res = longestCommonPrefix1(strs)
	fmt.Println(res)
	fmt.Println("end: ", time.Now().Sub(start))
}
