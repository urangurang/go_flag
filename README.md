# ECFP Go Study - Flag, Park Jinsoo

## Go by Example: Command-Line Flags
[link](https://gobyexample.com/command-line-flags)


## How to load package in Go

```go
import "flag"
import "fmt"
```

```go
import (
    "flag"
    "fmt"
)
```


## Flag Example


**flag_test.go**

```go
package main

import "fmt"
import "flag"

func main() {
	name := flag.String("name", "Jinsoo", "This is for name.")
	flag.Parse()
	fmt.Println(name)
}
```

### go build 


```
$ go build flag_test.go
go build command-line-arguments: no non-test Go files in /Users/jinsoo.park/Desktop

```

### Error ?!
In go lang, the file name xxx_test.go means that this file is test case.
so please set another name for it.

```
$ mv flag_test.go flag_example.go
$ go build flag_example.go
$ ls
flag_example flag_example.go

$ ls -al
-rwxr-xr-x    1 jinsoo.park  INTRA\Domain Users  2044800 Mar 22 15:02 flag_example
-rw-r--r--@   1 jinsoo.park  INTRA\Domain Users      160 Mar 22 15:01 flag_example.go
```

### flag_example is executable

```
$ ./flag_example
0xc4200421d0

$ ./flag_example -name=man
0xc4200421d0
```

### Why memory address?

```go
package main

import "fmt"
import "flag"

func main() {
	name := flag.String("name", "Jinsoo", "This is for name.")
	flag.Parse()
	fmt.Println(*name)
}
```

### How to use go Flag
```
$ go build flag_example.go

$ ./flag_example
jinsoo

$ ./flag_example -name=man
man

$ ./flag_example -help
$ ./flag_example -h
Usage of ./flag_example:
  -name string
    	This is for name. (default "Jinsoo")

$ ./flag_example -dsa
flag provided but not defined: -dsa
Usage of ./flag_example:
  -name string
    	This is for name. (default "Jinsoo")
```


### flag.Parse()
To parse the command line into the defined flags.

```
func main() {
    name := flag.String("name", "Jinsoo", "This is for name.")
    flag.Parse()
    fmt.Println(*name)

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    fmt.Println(*svar)
}
```

### Build fail
```
$ go build flag_example.go
# command-line-arguments
./flag_example.go:14:17: invalid indirect of svar (type string)
```


### Edit flag.StringVar
```
func main() {
    name := flag.String("name", "Jinsoo", "This is for name.")
    flag.Parse()
    fmt.Println(*name)

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    fmt.Println(svar)
}
```

### Build Success
```
$ go build flag_example.go
$ ./flag_example
Jinsoo
bar
```

### Use case
```
./flag_example -name=books -svar=bic
flag provided but not defined: -svar
Usage of ./flag_example:
  -name string
    	This is for name. (default "Jinsoo")
```

### Move flag.Parse()

```
func main() {
    name := flag.String("name", "Jinsoo", "This is for name."
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")
    
    flag.Parse()
    
    fmt.Println(*name)
    fmt.Println(svar)
}
```

```
$ ./flag_example -name=books -svar=bic
books
bic
```


## What's the difference?

[Official Document flag.go](https://golang.org/src/flag/flag.go)

### flag.String

String defines a string flag with specified name, default value, and usage string. 

The return value is the **address** of a string variable that stores the value of the flag.

```go
func (*FlagSet) String
func (f *FlagSet) String(name string, value string, usage string) *string
```


### flag.StringVar

StringVar defines a string flag with specified name, default value, and usage string.

The argument p points to a string variable in which to store the value of the flag.

```go
func (f *FlagSet) StringVar(p *string, name string, value string, usage string) {
	f.Var(newStringValue(value, p), name, usage)
}
```


### When Flag Redefine

```
package main

import "fmt"
import "flag"

func main()  {
    var ip = flag.Int("flagname", 321, "first define")
    var flagVar int
    flag.IntVar(&flagVar, "flagname", 2910, "redefine")
    fmt.Println(ip)
    fmt.Println(flagVar)
}
```

.

.

.

.

.

.

.

### Redefine Error

```
/private/var/folders/ph/tf09cxp10t32w2z853hk1lqc0000gn/T/___go_build_argument_example_go flag redefined: flagname
panic: /private/var/folders/ph/tf09cxp10t32w2z853hk1lqc0000gn/T/___go_build_argument_example_go flag redefined: flagname

goroutine 1 [running]:
```


## Quiz 1. What will be printed when execute last line?

```go
package main

import "fmt"
import "flag"

func main()  {
    var depart = flag.String("name", "FirstParty", "Department name")
    flag.Parse()
    fmt.Println(depart)
}
```

```
$ ./quiz01 -name=Books -name=Core Solution Group
```

### Please select right one.

1. Books
2. Core Solution Group
3. Core
4. 0xc42004e1e0(Memory address)
5. Redefine Error!?




### flag method - func (*FlagSet) Args
Args returns the non-flag arguments.
```go
func (f *FlagSet) Args() []string 

```



.

.

.

.

## Quiz 2. What will be printed when execute last line?

```go
package main

import "fmt"
import "flag"

func main()  {
    var spring = flag.Int("spring", 8080, "spring port")
    var flask = flag.Int("flask", 5000, "flask port")

    // parse flag arguments
    flag.Parse()
    fmt.Println(*spring)
    fmt.Println(*flask)

    fmt.Println(flag.Args())
    fmt.Println(flag.Args()[2])
}

$ ./quiz02 a b c d
[a b c d]
c

$ ./quiz02 -flask=8080 -spring=5000
```

.

.

.

.

.

.

.

.

#### panic: runtime error: index out of range





