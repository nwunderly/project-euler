
// Project Euler - Problem 1
// Multiples of 3 and 5

fun main() {
    var sum = 0
    for (i in 0..999) {
        if (i % 3 == 0 || i % 5 == 0) {
            sum += i
        }
    }
    println(sum)
}
