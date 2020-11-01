// Project Euler - Problem 1
// Multiples of 3 and 5

package main

import "fmt"

func init() {
	registerFunction("1", problem1)
}

func problem1() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if (i%3 == 0) || (i%5 == 0) {
			sum += i
		}
	}
	fmt.Println(sum)
}
