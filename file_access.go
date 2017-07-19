package main

import (
        "os"
        "log"
        "fmt"
)

func main() {

        file, err := os.Open("hw.go")
        if err != nil {
                log.Fatal(err)
        }

        data := make([]byte, 1000)
        count, err := file.Read(data)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Printf("read %d bytes: %q\n", count, data[900:count])

        fmt.Println(len(data))
}
