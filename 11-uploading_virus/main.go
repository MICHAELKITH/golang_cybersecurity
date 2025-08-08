package main

import (
    "fmt"
    "io"
    "net"
    "os"
)

func receiveFile(conn net.Conn, filePath string) {
    file, err := os.Create(filePath)
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    _, err = io.Copy(file, conn)
    if err != nil {
        fmt.Println("Error receiving file:", err)
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

    receiveFile(conn, "/tmp/virus.exe")
}