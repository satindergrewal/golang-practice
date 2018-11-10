package main

import "strings"
import "fmt"
import "kmdutil"

func main() {
    initial := "<h1>Hello World!</h1>"

    out := strings.TrimLeft(strings.TrimRight(initial,"</h1>"),"<h1>")
    fmt.Println(out)

    dir := kmdutil.AppDataDir("komodo", false)
    fmt.Println(dir)
}