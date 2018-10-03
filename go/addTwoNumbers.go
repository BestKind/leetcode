package main

import "fmt"

/**
 * 解题思路：a + b = target => a = target - b
 * 所以用一个 map[value] = index 就可以查询到 a 对应的 index
 */
func twoSum(num []int, target int) []int{
	maps := make(map[int]int)

	for i, value := range num {
		member := target - value

		if j, ok := maps[member]; ok {
			return []int{j, i}
		} else {
			maps[value] = i
		}
	}

	return []int{}
}

func main() {
	num := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(num, target)
	fmt.Println(res)
}
