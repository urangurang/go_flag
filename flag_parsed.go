package main

import "fmt"
import "flag"

func main() {
    fmt.Println(flag.Parsed())
    flag.Parse()
    fmt.Println(flag.Parsed())
}
