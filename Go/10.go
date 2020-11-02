/*
Project Euler - Problem 10
Summation of primes

https://projecteuler.net/problem=10
*/

package main

import "fmt"

func isPrime(n int) bool{
	for k := 0; k < len(primes); k++ {
		prime := primes[k]
		if prime > n / 2 {
			return true
		}
		if n % prime == 0 {
			return false
		}
	}
	return true
}

var primes = []int{2}

func problem10() {
	sum := 2
	for i := 3; i < 2000000; i++ {
		if isPrime(i) {
			sum += i
			primes = append(primes, i)
		}
	}
	fmt.Println(sum)
	fmt.Println(len(primes))
}

