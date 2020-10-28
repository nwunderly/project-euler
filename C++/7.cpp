
/*
Project Euler - Problem 7
10001st prime

https://projecteuler.net/problem=7
*/

#include <iostream>

using namespace std;

int testIfPrime(int n) {
    for (int i = 2; i <= n/2 + 1; i++) {
        if (n % i == 0) return false;
    }
    return true;
}

int main(void) {
    int lastPrime = 2; // first prime
    for (int i = 2; i <= 10001; i++) { // begin by looking for 2nd prime
        bool isPrime = false;
        int nextPrime = lastPrime;
        while (!isPrime) {
            nextPrime++;
            if (testIfPrime(nextPrime)) {
                isPrime = true;
                lastPrime = nextPrime;
            }
        }
    }
    cout << lastPrime << endl;
    return 0;
}