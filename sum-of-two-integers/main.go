package main

import (
	"fmt"
	"os"
)

func main() {
	expect("#1", getSum(1, 1), 2)
	expect("#2", getSum(4, 7), 11)
	expect("#3", getSum(1, 2), 3)
	expect("#4", getSum(20, 30), 50)
	expect("#5", getSum(-1, 1), 0)
	expect("#6", getSum(1, -2), -1)
	expect("#7", getSum(-1, -1), -2)
}

func expect[T comparable](message string, got T, expected T) {
	if got != expected {
		fmt.Printf("\033[31m> fail: %s (expected: %v, got: %v)\033[0m\n", message, expected, got)
		os.Exit(1)
	}
	fmt.Printf("\033[92m> pass: %s (got: %v)\033[0m\n", message, got)
}

func getSum(a int, b int) int {
	result := 0
	carry := 0
	for i := range 32 {
		aBit, bBit := (a>>i)&1, (b>>i)&1
		cBit := (carry ^ aBit ^ bBit)
		result |= cBit << i
		carry = ((carry & 1) & aBit) | ((carry & 1) & bBit) | (aBit & bBit)
	}
	if result > 0x7FFFFFFF {
		result = ^(result ^ 0xFFFFFFFF)
	}
	return result
}
