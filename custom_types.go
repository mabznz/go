package main

import "fmt"

const(
    KB = 1024
    MB = KB * 1024
    GB = MB * 1024
    TB = GB * 1024
    PB = TB * 1024
)

// New custom type that starts as a float
type ByteSize float64

func (b ByteSize) String() string {
    switch {
    case b >= PB:
        return "Too Big"
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b / TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b / GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b / MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b / KB)
    }
    return fmt.Sprintf("%dB", b)
}

func main() {
    fmt.Println(ByteSize(123456))
    fmt.Println(ByteSize(123683458275632))
    x := ByteSize(123456)
    fmt.Println(x)
}
