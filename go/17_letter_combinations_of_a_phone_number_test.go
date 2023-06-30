package main

import (
	"fmt"
	"testing"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	numToChar := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	results := []string{""}

	for _, digit := range digits {
		tempResults := []string{}

		for _, chr := range numToChar[string(digit)] {
			for _, res := range results {
				tempResults = append(tempResults, string(res)+string(chr))
			}
		}

		results = tempResults
	}

	return results
}

func TestLetterCombinations(t *testing.T) {
	fmt.Println(letterCombinations("234"))
}
