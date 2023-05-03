package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world!")
}

func main() {
    http.HandleFunc("/", helloHandler)

    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        panic(err)
    }
}
