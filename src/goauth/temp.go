package main

import "strings"
import "fmt"

func main() {
    initial := "<h1>Hello World!</h1>"

    out := strings.TrimLeft(strings.TrimRight(initial,"</h1>"),"<h1>")
    fmt.Println(out)
}