
// Project Euler - Problem 2
// Even Fibonacci Numbers

fun main() {

    var a = 1 // last fibonacci number
    var b = 1 // current fibonacci number
    var c: Int // helper variable
    var sum = 0

    while (b < 4000000) {
        c = a // keep track of a
        a = b // redefine a
        b = c + b // redefine b using "old a"
        if (b % 2 == 0) {
            sum += b
        }
    }

    println(sum)

}