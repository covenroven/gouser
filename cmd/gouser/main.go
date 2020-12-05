package main

import (
    "fmt"
    "net/http"

    "github.com/covenroven/gouser/internal/router"
)

func main() {
    r, err := router.Init()
    if (err != nil) {
        log.Fatal("Failed to initialize router")
    }

    http.ListenAndServe(":3000", r)
}
