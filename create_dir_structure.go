package main

import (
        "os"
        "fmt"
        "time"
)

func main() {
        t := time.Now()
        dir := "/tmp/" + t.Format("20060102")

        x := []string{"AAA","BBB"}
        for _, value := range x {
                dir := dir + "/" + value
                if !(checkPath(dir)) {
                        err := os.MkdirAll(dir, os.ModePerm)
                        if err != nil {
                                panic("")
                        }
                }
        }

        fmt.Println("End")
}

func checkPath(path string) (exists bool) {
        _, err := os.Stat(path)

        if os.IsNotExist(err) {
                if err != nil {
                        return false
                }
        }
        return true
}
