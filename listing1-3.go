package main

import (
    "net"
    "fmt"
    "bufio"
)

func main() {
    conn, _ := net.Dial("tcp", "golang.org:80")
    fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
    s, _ := bufio.NewReader(conn).ReadString('\n')

    fmt.Println(s)
}

/*
HTTP/1.0 200 OK

*/