package main

import (
    "fmt"
    "flag"
)

func main(){
    name := flag.String("name", "jinsoo", "This is for name")
    flag.Parse()
   
    fmt.Println("Hello ", *name)
}
