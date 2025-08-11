package main

import (
    "fmt"
    "github.com/atotto/clipboard"
    "time"
)

func main() {
    for {
        time.Sleep(1 * time.Second)
        text, err := clipboard.ReadAll()
        if err != nil {
            fmt.Println("Error reading clipboard:", err)
            continue
        }
        fmt.Println("Captured:", text)
    }
}