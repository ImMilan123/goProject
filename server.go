package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Hello World :)")
    })

    fmt.Println("running")
    http.ListenAndServe(":8080", nil)
}