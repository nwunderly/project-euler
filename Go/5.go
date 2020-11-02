// Project Euler - Problem 5
// Smallest multiple

package main

import "fmt"

func makeDivisibleBy(number int, factor int) int {
	for i := 1; i <= factor; i++ {
		if (number*i)%factor == 0 {
			return i
		}
	}
	return 0
}

func problem5() {
	var n = 1
	for i := 1; i <= 20; i++ {
		n *= makeDivisibleBy(n, i)
		fmt.Println(i, n)
	}
	fmt.Println(n)
}
