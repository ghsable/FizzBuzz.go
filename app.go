package main

import (
    "strconv"
)

func getPrint(n int) string {
    if (n%15 == 0) {
        return "Fizz Buzz"
    } else if (n%5 == 0) {
        return "Buzz"
    } else if (n%3 == 0) {
        return "Fizz"
    } else {
        return strconv.Itoa(n)
    }
}
