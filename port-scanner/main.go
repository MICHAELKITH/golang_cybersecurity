package main

import (
    "fmt"
    "net"
    "time"
)

func scanPort(target string, port int) {
    address := fmt.Sprintf("%s:%d", target, port)
    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err != nil {
        return
    }
    conn.Close()
    fmt.Printf("Port %d is open\n", port)
}

func main() {
    target := "facebook.com"
    for port := 1; port <= 1024; port++ {
        go scanPort(target, port)
    }
    time.Sleep(2 * time.Second)
}