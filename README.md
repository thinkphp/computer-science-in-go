# Essential Go

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

## Variables

* var declares 1 or more variables.
* You can declare multiple variables at once.
* Go will infer the type of initialized variables
* Variables declared without a corresponding initialization are zero-valued. For example, the zero value for an int is 0.
* The := syntax is shorthand for declaring and initializing a variable, e.g. for var f string = "apple" in this case. This syntax is only available inside functions.
```
package main

import "fmt"

func main() {

    var a = "initial"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)

    var d = true
    fmt.Println(d)

    var e int
    fmt.Println(e)

    f := "apple"
    fmt.Println(f)
}
```


## Examples

```

package main

import "fmt"

func bubblesort(arr []int) {

     done := true
 
     var changes, i int 
 
     for done {

         changes = 0
         
         for i = 0; i < len(arr) - 1; i++ {

             if arr[ i ] > arr[ i + 1 ] {

                arr[ i ], arr[ i + 1 ] = arr[ i + 1 ], arr[ i ]

                changes++ 
             }   
         } 

         if changes == 0 {

            done = false
         }
     }
}

func main() {

     arr := []int{9,8,7,6,5,4,3,2,1}

     bubblesort( arr )

     for _, v := range arr {

         fmt.Printf("%d ", v)
     }   
}

```




## References

https://golang.org/

https://go.dev/doc/effective_go

https://gobyexample.com/

https://www.programming-books.io/essential/go/

https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
