package main

import (
    "fmt"
    "net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
    fmt.Fprint(res, "hello marshmellow")
}

func main() {
    http.HandleFunc("/", hello)
    http.ListenAndServe("localhost:4848", nil)
}
