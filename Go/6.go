/*
Project Euler - Problem 6
Sum square difference

https://projecteuler.net/problem=6
*/

package main

import "fmt"

func problem6() {
	sumSquares := 0
	squareSum := 0
	for i := 1; i <= 100; i++ {
		sumSquares += i * i
		squareSum += i
	}
	squareSum = squareSum * squareSum
	difference := squareSum - sumSquares
	fmt.Println(difference)
}
