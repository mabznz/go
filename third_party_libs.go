package main

import (
    "fmt"
    "github.com/russross/blackfriday"
)

func main() {
    markdown := []byte(`
# Header
* list
* more list

* Another list
* and again
    `)

    html := blackfriday.MarkdownBasic(markdown)
    fmt.Println(string(html))
}
