package main

import (
    "fmt"
    "os"
    "strconv"
    //"sync"
)


func main() {
    //message := make(chan int)
    //var wg sync.WaitGroup
    //wg.Add(2)
    numOfIterations, err := strconv.Atoi(os.Args[1])
    if(err != nil) {
        fmt.Println("Error Converting Argument")
    }
    positive := countPositives(numOfIterations)
    negative := countNegatives(numOfIterations)
    //wg.Wait()
    pi := positive - negative
    fmt.Println(pi * 4.0)
}

func countNegatives(iterations int) float64 {
    //defer wg.Done()
    sum := 0.0
    bottom := 3.0
    for i := 0; i < iterations; i++ {
        val := 1.0 / bottom
        sum = sum + val
        bottom = bottom + 4.0
    }
    //message <- 2
    return sum
}

func countPositives(iterations int) float64 {
    //defer wg.Done()
    sum := 0.0
    bottom := 1.0
    for i := 0; i < iterations; i++ {
        val := 1.0 / bottom
        sum = sum + val
        bottom = bottom + 4
    }
    //message <- 1
    return sum
}
