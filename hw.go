package main

import "fmt"

func main() {
    fmt.Println("Hello")
    mySlice := []int{4,7,9,3,8}
    fmt.Println(mySlice[0:1])
    for index, number := range mySlice {
        fmt.Println("Index:" + string(index) + " Value:" + string(number))
        fmt.Println(index,number)
    }

    sites := map[string]int {
        "Taita": 4,
        "Naenae": 2,
        "Avalon": 3,
    }

    for site, quantity := range sites {
        fmt.Println(site, "has", quantity, "sites.")
    }
    myString, myInt := threeParamsTwoReturns(2, 3.4, mySlice)
    fmt.Println(myString, myInt)

    point := "Not Pointing"
    fmt.Println(point);
    demoPointer(&point)
    fmt.Println(point);
}

func threeParamsTwoReturns(myInt int, myFloat float64, mySlice []int) (string, int) {
    fmt.Println(mySlice[0:5])
    return "Return", myInt + 100
}

func demoPointer(pointer *string) {
    *pointer = "Now pointing. Function scope has changed point variable."
}
