package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("vim-go")
	answer := make([]string, 15+1)
	for i := 1; i <= 15; i++ {
		if i%3 == 0 && i%5 == 0 {
			answer[i] = "FizzBuzz"
		} else if i%5 == 0 {
			answer[i] = "Buzz"
		} else if i%3 == 0 {
			answer[i] = "Fizz"
		} else {
			answer[i] = strconv.Itoa(i)
		}
	}
	fmt.Println(answer)
}
