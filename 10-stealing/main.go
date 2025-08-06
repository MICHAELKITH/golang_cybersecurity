package main

import (
    "fmt"
    "io"
    "net"
    "os"
    // "path/filepath"
)

func sendFile(conn net.Conn, filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    _, err = io.Copy(conn, file)
    if err != nil {
        fmt.Println("Error sending file:", err)
    }
}

func main() {
    listener, err := net.Listen("tcp", ":4444")
    if err != nil {
        fmt.Println("Error starting server:", err)
        return
    }
    defer listener.Close()

    conn, err := listener.Accept()
    if err != nil {
        fmt.Println("Error accepting connection:", err)
        return
    }
    defer conn.Close()

    sendFile(conn, "/home/user/sensitive_file.txt")
}