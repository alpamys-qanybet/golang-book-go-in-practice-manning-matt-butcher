package main

import (
    "fmt"
    "time"
)

func printCount(c chan int) {
    n := 0
    for n >= 0 {
        n = <- c
        fmt.Print(n, " ")
    }
}

func main() {
    c := make(chan int)
    a := []int{8,6,7,5,3,0,9,-1}
    // a := []int{8,6,7,-3,5,3,0,9,-1} // it will throw error: 8 6 7 -3 fatal error: all goroutines are asleep - deadlock!

    go printCount(c)

    for _, v := range a {
        c <- v
    }

    time.Sleep(time.Millisecond)
    fmt.Println("End of main")
}

// 8 6 7 5 3 0 9 -1 End of main