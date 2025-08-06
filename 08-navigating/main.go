package main

import (
    "fmt"
    "io/ioutil"
    "path/filepath"
)

func listFiles(dir string) {
    files, err := ioutil.ReadDir(dir)
    if err != nil {
        fmt.Println("Error reading directory:", err)
        return
    }
    for _, file := range files {
        fmt.Println(filepath.Join(dir, file.Name()))
    }
}

func main() {
    listFiles("/home/user")
}