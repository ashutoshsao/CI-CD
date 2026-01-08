package main

import (
    "fmt"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "hello from GoLang from aws")
    })

    fmt.Println("Server is running on http://localhost:8080")
    
    // This 'log.Fatal' is the critical part!
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
