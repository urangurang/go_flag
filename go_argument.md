# ECFP Go Study - Flag, Park Jinsoo

## Go by Example: Command-Line Arguments

Command-line arguments are a common way to parameterize execution of programs. For example, go run hello.go uses run and hello.go arguments to the go program.

os.Args provides access to raw command-line arguments. Note that the first value in this slice is the path to the program, and os.Args[1:] holds the arguments to the program.

You can get individual args with normal indexing.

```
package main

import "os"
import "fmt"

func main() {
    argsWithProg := os.Args
    argsWithoutProg := os.Args[1:]
    arg := os.Args[3]

    fmt.Println(argsWithProg)
    fmt.Println(argsWithoutProg)
    fmt.Println(arg)
}
```

Build go file

```
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c
```

Run without arguments

```
$ ./command-line-arguments
panic: runtime error: index out of range

goroutine 1 [running]:
main.main()
	/Users/Urang/Desktop/go_study/command-line-arguments.go:11 +0x210
```
