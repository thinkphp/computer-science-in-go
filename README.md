## Computer Science in Go language

Go is a new language. Although it borrows ideas from existing languages, it has unusual properties that make effective Go programs different in character from programs written in its relatives. A straightforward translation of a C++ or Java program into Go is unlikely to produce a satisfactory result—Java programs are written in Java, not Go. On the other hand, thinking about the problem from a Go perspective could produce a successful but quite different program. In other words, to write Go well, it's important to understand its properties and idioms. It's also important to know the established conventions for programming in Go, such as naming, formatting, program construction, and so on, so that programs you write will be easy for other Go programmers to understand.

```
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}

$ go run hello-world.go
hello world

$ go build hello-world.go
$ ls
hello-world    hello-world.go

$ ./hello-world
hello world
```


https://golang.org/

https://en.wikipedia.org/wiki/Go_(programming_language)

## References

https://go.dev/doc/effective_go

https://www.programming-books.io/essential/go/

https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
