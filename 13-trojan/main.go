package main

import (
    "archive/zip"
    "fmt"
    "io"
    "os"
    "path/filepath"
)

func createTrojan(source, destination string) error {
    zipFile, err := os.Create(destination)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    zipWriter := zip.NewWriter(zipFile)
    defer zipWriter.Close()

    file, err := os.Open(source)
    if err != nil {
        return err
    }
    defer file.Close()

    info, err := file.Stat()
    if err != nil {
        return err
    }

    header, err := zip.FileInfoHeader(info)
    if err != nil {
        return err
    }

    header.Method = zip.Deflate

    writer, err := zipWriter.CreateHeader(header)
    if err != nil {
        return err
    }

    _, err = io.Copy(writer, file)
    return err
}

func main() {
    if err := createTrojan("malware.exe", "legitimate_application.zip"); err != nil {
        fmt.Println("Error creating trojan:", err)
    } else {
        fmt.Println("Trojan created successfully")
    }
}