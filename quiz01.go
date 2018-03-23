package main

import "fmt"
import "flag"

func main()  {
    var depart = flag.String("name", "FirstParty", "Department name")
    flag.Parse()
    fmt.Println(*depart)
}
