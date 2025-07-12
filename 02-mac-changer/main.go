package main

import (
    "fmt"
    
    "os/exec"
)

func changeMAC(interfaceName, newMAC string) error {
    cmd := exec.Command("ifconfig", interfaceName, "down")
    if err := cmd.Run(); err != nil {
        return err
    }

    cmd = exec.Command("ifconfig", interfaceName, "hw", "ether", newMAC)
    if err := cmd.Run(); err != nil {
        return err
    }

    cmd = exec.Command("ifconfig", interfaceName, "up")
    if err := cmd.Run(); err != nil {
        return err
    }

    return nil
}

func main() {
    if err := changeMAC("eth0", "00:11:22:33:44:55"); err != nil {
        fmt.Println("Error changing MAC address:", err)
    } else {
        fmt.Println("MAC address changed successfully")
    }
}