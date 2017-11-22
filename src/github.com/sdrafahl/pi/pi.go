package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    numOfIterations, err := strconv.Atoi(os.Args[1])
    if(err != nil) {
        fmt.Println("Error Converting Argument")
    }
    ch := make(chan float64)
    go countPositives(numOfIterations, ch)
    go countNegatives(numOfIterations, ch)
    negative, positive := <- ch, <- ch
    pi := positive - negative
    fmt.Println(pi * 4.0)
}

func countNegatives(iterations int, ch chan float64) {
    sum := 0.0
    bottom := 3.0
    for i := 0; i < iterations; i++ {
        val := 1.0 / bottom
        sum = sum + val
        bottom = bottom + 4.0
    }
    ch <- sum
}

func countPositives(iterations int, ch chan float64) {
    sum := 0.0
    bottom := 1.0
    for i := 0; i < iterations; i++ {
        val := 1.0 / bottom
        sum = sum + val
        bottom = bottom + 4
    }
    ch <- sum
}
