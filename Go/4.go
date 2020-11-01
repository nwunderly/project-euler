// Project Euler - Problem 4
// Largest palindrome product

package main

import (
	"fmt"
	"strconv"
)

func init() {
	registerFunction("4", problem4)
}

func is_palindrome(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func problem4() {

	var largest int

	for i := 999; i > 100; i-- {
		for j := 999; j > 100; j-- {
			n := i * j
			s := strconv.Itoa(n)
			if is_palindrome(s) {
				if n > largest {
					largest = n
				}
			}
		}
	}
	fmt.Println(largest)
}
