# Essential Go

### Course Description
This course introduces students to the Go programming language, focusing on its unique features, syntax, and practical applications. Students will learn the fundamentals of Go and gain hands-on experience through various projects and exercises.

### Course Objectives
By the end of this course, students will be able to:
1. Understand Go's syntax and basic concepts
2. Write and run Go programs
3. Implement Go's concurrency features
4. Work with Go packages and tools
5. Develop small to medium-sized applications in Go

### Week-by-Week Schedule

#### Week 1: Introduction to Go
- History and philosophy of Go
- Setting up the Go development environment
- "Hello, World!" in Go
  ```go
  package main

  import "fmt"

  func main() {
      fmt.Println("Hello, Harvard!")
  }
  ```
- Basic syntax and data types
  ```go
  var i int = 42
  f := 3.14
  s := "Go is awesome"
  b := true
  ```

#### Week 2: Control Structures and Functions
- Conditionals (if, else, switch)
  ```go
  if x > 0 {
      fmt.Println("Positive")
  } else if x < 0 {
      fmt.Println("Negative")
  } else {
      fmt.Println("Zero")
  }

  switch day {
  case "Monday":
      fmt.Println("Start of the week")
  case "Friday":
      fmt.Println("TGIF!")
  default:
      fmt.Println("Regular day")
  }
  ```
- Loops (for, range)
  ```go
  for i := 0; i < 5; i++ {
      fmt.Println(i)
  }

  numbers := []int{1, 2, 3, 4, 5}
  for index, value := range numbers {
      fmt.Printf("Index: %d, Value: %d\n", index, value)
  }
  ```
- Functions and methods
  ```go
  func add(a, b int) int {
      return a + b
  }

  type Rectangle struct {
      width, height float64
  }

  func (r Rectangle) Area() float64 {
      return r.width * r.height
  }
  ```
- Error handling
  ```go
  func divide(a, b float64) (float64, error) {
      if b == 0 {
          return 0, errors.New("division by zero")
      }
      return a / b, nil
  }
  ```

#### Week 3: Data Structures
- Arrays and slices
  ```go
  var arr [5]int = [5]int{1, 2, 3, 4, 5}
  slice := []int{1, 2, 3, 4, 5}
  slice = append(slice, 6)
  ```
- Maps
  ```go
  studentGrades := map[string]int{
      "Alice": 92,
      "Bob":   87,
      "Charlie": 95,
  }
  ```
- Structs
  ```go
  type Person struct {
      Name string
      Age  int
  }
  p := Person{Name: "Alice", Age: 30}
  ```
- Pointers
  ```go
  x := 42
  ptr := &x
  fmt.Println(*ptr) // Prints 42
  ```

#### Week 4: Object-Oriented Programming in Go
- Interfaces
  ```go
  type Shape interface {
      Area() float64
  }

  type Circle struct {
      Radius float64
  }

  func (c Circle) Area() float64 {
      return math.Pi * c.Radius * c.Radius
  }
  ```
- Embedding
  ```go
  type Animal struct {
      Name string
  }

  type Dog struct {
      Animal
      Breed string
  }
  ```
- Composition vs. inheritance
  ```go
  type Writer interface {
      Write([]byte) (int, error)
  }

  type ConsoleWriter struct{}

  func (cw ConsoleWriter) Write(data []byte) (int, error) {
      return fmt.Println(string(data))
  }
  ```

#### Week 5: Concurrency
- Goroutines
  ```go
  go func() {
      fmt.Println("Hello from a goroutine!")
  }()
  ```
- Channels
  ```go
  ch := make(chan int)
  go func() {
      ch <- 42
  }()
  value := <-ch
  ```
- Select statement
  ```go
  select {
  case msg1 := <-ch1:
      fmt.Println("Received from ch1:", msg1)
  case msg2 := <-ch2:
      fmt.Println("Received from ch2:", msg2)
  case <-time.After(time.Second):
      fmt.Println("Timeout")
  }
  ```
- Synchronization primitives
  ```go
  var mu sync.Mutex
  var wg sync.WaitGroup
  ```

#### Week 6: Packages and Modules
- Creating and using packages
  ```go
  // math/math.go
  package math

  func Add(a, b int) int {
      return a + b
  }

  // main.go
  import "myproject/math"

  result := math.Add(3, 4)
  ```
- Go modules
  ```
  go mod init github.com/username/project
  go get github.com/somepackage
  ```
- Third-party package management
  ```
  go get -u github.com/gin-gonic/gin
  ```

#### Week 7: Testing and Debugging
- Unit testing with the testing package
  ```go
  func TestAdd(t *testing.T) {
      result := Add(2, 3)
      if result != 5 {
          t.Errorf("Add(2, 3) = %d; want 5", result)
      }
  }
  ```
- Benchmarking
  ```go
  func BenchmarkFibonacci(b *testing.B) {
      for i := 0; i < b.N; i++ {
          Fibonacci(20)
      }
  }
  ```
- Debugging tools and techniques
  ```go
  import "log"
  log.Printf("Debug: x = %d", x)
  ```

#### Week 8: Web Development with Go
- Introduction to net/http package
  ```go
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello, Harvard!")
  })
  http.ListenAndServe(":8080", nil)
  ```
- RESTful API development
  ```go
  // Using Gin framework
  r := gin.Default()
  r.GET("/api/users", func(c *gin.Context) {
      c.JSON(200, gin.H{
          "users": []string{"Alice", "Bob", "Charlie"},
      })
  })
  ```
- Working with databases (SQL and NoSQL)
  ```go
  // SQL example
  db, _ := sql.Open("mysql", "user:password@/dbname")
  rows, _ := db.Query("SELECT * FROM users")
  ```

#### Week 9: Advanced Topics
- Reflection
  ```go
  t := reflect.TypeOf(3)
  fmt.Println(t.String()) // "int"
  ```
- Context package
  ```go
  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  ```
- Best practices and common patterns
  ```go
  // Error wrapping
  if err != nil {
      return fmt.Errorf("failed to process data: %w", err)
  }
  ```

#### Week 10: Final Project
- Students will work on a final project implementing a small application or service using Go

### Assessment
- Weekly coding assignments (40%)
- Midterm exam (20%)
- Final project (30%)
- Class participation (10%)

### Required Materials
- "The Go Programming Language" by Alan A. A. Donovan and Brian W. Kernighan
- Access to a computer with Go installed (version 1.16 or later)

### Online Resources
- [Go Official Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)


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

- https://golang.org/

- https://cs.lmu.edu/~ray/notes/introgo/

- https://go.dev/doc/effective_go

- https://gobyexample.com/

- https://www.programming-books.io/essential/go/

- https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs
