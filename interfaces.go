package main

import "fmt"

type Fruit interface {
    String() string
}

type Apple struct {
    Variety string
}

func (a Apple) String() string {
    return fmt.Sprintf("A %s apple", a.Variety)
}

type Orange struct {
    Size string
}

func (o Orange) String() string {
    return fmt.Sprintf("A %s orange", o.Size)
}

func PrintFruit(fruit Fruit) {
    fmt.Println("I have this fruit:", fruit.String())
}

func main() {
    // create
    apple := Apple{"Pink Lady"}
    orange := Orange{'"large"}

    PrintFruit(apple)
    PrintFruit(orange)
}
