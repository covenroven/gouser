package main

import (
    "net/http"
    // "fmt"
    "log"

    "github.com/covenroven/gouser/internal/router"
)

func main() {
    // Initialize router
    r, err := router.Init()
    if err != nil {
        log.Fatal("Failed to initialize router", err)
    }

    http.ListenAndServe(":3000", r)
}
