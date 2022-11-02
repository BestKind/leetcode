package main

import (
	"fmt"
	"math/bits"
)

/**
通过位运算 来统计 每行 每列 以及每个宫 之内都出现了哪些数字
然后通过递归的方式来计算出每个待填写的空格数字
 */
func solveSudoku(board [][]byte) [][]byte {
	var line, column [9]int
	var block [3][3]int
	var spaces [][2]int

	flip := func(i, j int, digit byte) {
		line[i] ^= 1 << digit
		column[j] ^= 1 << digit
		block[i/3][j/3] ^= 1 << digit
	}
	for i, row := range board {
		for j, b := range row {
			if '.' == b {
				spaces = append(spaces, [2]int{i, j})
			} else {
				// - 1 是因为 数组下标从 0 开始
				digit := b - '1'
				flip(i, j, digit)
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1]
		mask := 0x1ff &^ uint(line[i]|column[j]|block[i/3][j/3])
		for ; mask > 0; mask &= mask - 1 {
			digit := byte(bits.TrailingZeros(mask))
			flip(i, j, digit)
			// + 1 是因为数组下标从 0 开始
			board[i][j] = digit + '1'
			if dfs(pos + 1) {
				return true
			}
			flip(i, j, digit)
		}
		return false
	}
	dfs(0)
	return board
}

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	res := solveSudoku(board)
	for _, row := range res {
		for _, v := range row {
			fmt.Print(int(v-'0'), " ")
		}
		fmt.Println()
	}
}
