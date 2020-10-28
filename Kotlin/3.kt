
// Project Euler - Problem 3
// Largest prime factor

fun is_prime(n: Long): Boolean {
    for (i in 2L..n/2L) {
        if (n % i == 0L) {
            return false
        }
    }
    return true
}

fun is_factor(big: Long, small: Long): Boolean {
    return (big % small == 0L)
}

fun get_largest_factor(number: Long): Long {
    var i = 2L
    while (i < number/2L) {
        if (is_factor(number, i)) {
            var possible_prime_factor = number / i
            if (is_factor(number, possible_prime_factor) && is_prime(possible_prime_factor)) {
                return possible_prime_factor
            }
        }
        i++
    }
    return -1
}

fun main() {
    val number = 600851475143L
    println(get_largest_factor(number))
}