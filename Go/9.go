/*
Project Euler - Problem 9
Special Pythagorean triplet

https://projecteuler.net/problem=9
*/

package main

import "fmt"

func isPythagorean(a int, b int, c int) bool {
	return (a*a + b*b == c*c) || (a*a + c*c == b*b) || (b*b + c*c == a*a)
}

func problem9() {
	n := 0
	for i := 1; i < 1000; i++ {
		for j := 0; j < 1000 - i; j++ {
			n++
			k := 1000 - (i + j)
			if isPythagorean(i, j, k) {
				fmt.Println("numbers:", i, j, k)
				fmt.Println("product:", i*j*k)
				fmt.Println("loops:", n)
				return
			}
		}
	}
}