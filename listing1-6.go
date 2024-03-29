package main

import (
    "fmt"
    "time"
)

func count() {
    for i := 0; i < 5; i++ {
        fmt.Println(i)
        time.Sleep(time.Millisecond)
    }
}

func main() {
    go count()
    time.Sleep(time.Millisecond * 2)
    fmt.Println("Hello world")
    time.Sleep(time.Millisecond * 5)
}


/*
0
1
Hello world
2
3
4
*/
