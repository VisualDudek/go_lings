# Quick Start Examples for Each Chapter

These are minimal examples showing the core idea of each chapter.

## 01: Fundamentals - Hello Go

```go
package main

import "fmt"

func main() {
    // Variables
    var name string = "Go"
    age := 15  // Type inference
    
    // Constants
    const version = "1.21"
    
    // Basic operations
    fmt.Printf("Hello, %s! Version: %s, Age: %d\n", name, version, age)
    
    // Control flow
    if age > 10 {
        fmt.Println("Go is mature!")
    }
    
    // Loops
    for i := 0; i < 3; i++ {
        fmt.Printf("Iteration %d\n", i)
    }
}
```

**Concepts**: var, const, := short declaration, if/else, for loops

---

## 02: Functions - Parameter Passing & Returns

```go
package main

import "fmt"

// Multiple return values (idiomatic for error handling)
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("division by zero")
    }
    return a / b, nil
}

// Variadic function
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// Function as value
func applyTwice(fn func(int) int, x int) int {
    return fn(fn(x))
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("Result: %.2f\n", result)
    
    fmt.Println("Sum:", sum(1, 2, 3, 4, 5))
    
    double := func(x int) int { return x * 2 }
    fmt.Println("Apply twice:", applyTwice(double, 3))
}
```

**Concepts**: Multiple returns, error handling, variadic functions, closures, higher-order functions

---

## 03: Data Structures - Slices & Maps

```go
package main

import "fmt"

type Person struct {
    Name string
    Age  int
}

func main() {
    // Slices - dynamic arrays
    numbers := []int{1, 2, 3}
    numbers = append(numbers, 4, 5)
    fmt.Println("Slice:", numbers)
    fmt.Println("Length:", len(numbers), "Capacity:", cap(numbers))
    
    // Maps - key-value storage
    scores := map[string]int{
        "Alice": 95,
        "Bob":   87,
    }
    scores["Charlie"] = 92
    fmt.Println("Scores:", scores)
    
    // Structs - group related data
    people := []Person{
        {Name: "Alice", Age: 30},
        {Name: "Bob", Age: 25},
    }
    for _, p := range people {
        fmt.Printf("%s is %d\n", p.Name, p.Age)
    }
}
```

**Concepts**: Slices, append, maps, structs, ranges

---

## 04: Interfaces - Abstraction

```go
package main

import "fmt"

// Interface - defines behavior, not implementation
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

type Cat struct {
    Name string
}

// Dog implements Animal
func (d Dog) Speak() string {
    return fmt.Sprintf("%s says Woof!", d.Name)
}

// Cat implements Animal
func (c Cat) Speak() string {
    return fmt.Sprintf("%s says Meow!", c.Name)
}

func makeAnimalSpeak(a Animal) {
    fmt.Println(a.Speak())
}

func main() {
    animals := []Animal{
        Dog{Name: "Rex"},
        Cat{Name: "Whiskers"},
    }
    
    for _, animal := range animals {
        makeAnimalSpeak(animal)
    }
}
```

**Concepts**: Interface definition, implicit satisfaction, type assertions, polymorphism

---

## 05: Concurrency - Goroutines & Channels

```go
package main

import (
    "fmt"
    "time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job)
        time.Sleep(100 * time.Millisecond)
        results <- job * 2
    }
}

func main() {
    jobs := make(chan int, 10)
    results := make(chan int)
    
    // Start 3 workers
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // Send jobs
    go func() {
        for j := 1; j <= 5; j++ {
            jobs <- j
        }
        close(jobs)
    }()
    
    // Collect results
    for r := 0; r < 5; r++ {
        fmt.Printf("Result: %d\n", <-results)
    }
}
```

**Concepts**: Goroutines (go keyword), channels, send/receive, closing channels, worker pattern

---

## 06: Error Handling - Custom Errors

```go
package main

import (
    "errors"
    "fmt"
)

// Custom error type
type ValidationError struct {
    Field   string
    Message string
}

func (e ValidationError) Error() string {
    return fmt.Sprintf("validation error in %s: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
    if email == "" {
        return ValidationError{Field: "email", Message: "empty"}
    }
    return nil
}

func main() {
    err := validateEmail("")
    if err != nil {
        // Type assertion to handle custom error
        if ve, ok := err.(ValidationError); ok {
            fmt.Printf("Field: %s, Message: %s\n", ve.Field, ve.Message)
        }
    }
    
    // Error wrapping with context
    err = fmt.Errorf("failed to validate: %w", validateEmail(""))
    if errors.Is(err, fmt.Errorf("")) {
        fmt.Println("Wrapped error detected")
    }
}
```

**Concepts**: Error interface, custom error types, error assertions, error wrapping

---

## 07: Packages - Module Organization

```
myapp/
├── go.mod
├── main.go
├── math/
│   └── calc.go
└── utils/
    └── helpers.go
```

```go
// main.go
package main

import (
    "fmt"
    "myapp/math"
)

func main() {
    result := math.Add(5, 3)
    fmt.Println("Result:", result)
}

// math/calc.go
package math

func Add(a, b int) int {
    return a + b
}

// go.mod
module myapp

go 1.21
```

**Concepts**: Package structure, imports, public/private (capitalization), go.mod, dependency management

---

## 08: I/O - Reading & Writing

```go
package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

func main() {
    // File operations
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer file.Close()
    
    // Read with buffering
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    
    // Work with io.Reader interface
    reader := strings.NewReader("Hello, World!")
    io.Copy(os.Stdout, reader)
}
```

**Concepts**: os.Open, defer, bufio.Scanner, io.Reader/Writer interfaces, buffering

---

## 09: Testing - Unit Tests & Benchmarks

```go
package math

import "testing"

// Unit test
func TestAdd(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive", 1, 2, 3},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Add(%d, %d) = %d, want %d", 
                    tt.a, tt.b, result, tt.expected)
            }
        })
    }
}

// Benchmark test
func BenchmarkAdd(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Add(5, 3)
    }
}
```

**Concepts**: Table-driven tests, t.Run subtests, benchmarks, test organization

---

## 10: Web - HTTP Server

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Message struct {
    Text string `json:"text"`
}

func handleHello(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(Message{Text: "Hello, World!"})
}

func handleEcho(w http.ResponseWriter, r *http.Request) {
    msg := r.URL.Query().Get("msg")
    w.Header().Set("Content-Type", "text/plain")
    fmt.Fprintf(w, "Echo: %s", msg)
}

func main() {
    http.HandleFunc("/hello", handleHello)
    http.HandleFunc("/echo", handleEcho)
    
    fmt.Println("Server starting on :8080")
    http.ListenAndServe(":8080", nil)
}
```

**Concepts**: http.HandleFunc, handlers, routing, JSON encoding, headers

---

## 11: Databases - SQL Operations

```go
package main

import (
    "database/sql"
    "fmt"
    _ "github.com/sqlite3"
)

type User struct {
    ID   int
    Name string
    Age  int
}

func main() {
    db, err := sql.Open("sqlite3", ":memory:")
    if err != nil {
        panic(err)
    }
    defer db.Close()
    
    // Create table
    db.Exec(`CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)`)
    
    // Insert
    result, _ := db.Exec("INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 30)
    id, _ := result.LastInsertId()
    
    // Query
    var user User
    db.QueryRow("SELECT id, name, age FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Name, &user.Age)
    
    fmt.Printf("User: %+v\n", user)
}
```

**Concepts**: database/sql, connections, prepared statements, scanning results, transactions

---

## 12: Reflection - Type Inspection

```go
package main

import (
    "fmt"
    "reflect"
)

type Person struct {
    Name string `field:"name"`
    Age  int    `field:"age"`
}

func main() {
    p := Person{Name: "Alice", Age: 30}
    t := reflect.TypeOf(p)
    v := reflect.ValueOf(p)
    
    fmt.Printf("Type: %v, Kind: %v\n", t.Name(), t.Kind())
    
    // Iterate struct fields
    for i := 0; i < t.NumField(); i++ {
        field := t.Field(i)
        value := v.Field(i)
        tag := field.Tag.Get("field")
        fmt.Printf("%s (%s): %v\n", field.Name, tag, value.Interface())
    }
}
```

**Concepts**: reflect.Type, reflect.Value, struct tags, introspection

---

## 13: Advanced Concurrency - Synchronization

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

func main() {
    // WaitGroup - wait for multiple goroutines
    var wg sync.WaitGroup
    var mu sync.Mutex
    counter := 0
    
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            mu.Lock()
            counter++
            mu.Unlock()
        }()
    }
    
    wg.Wait()
    fmt.Printf("Counter: %d\n", counter)
    
    // Context for cancellation
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()
    
    select {
    case <-time.After(3 * time.Second):
        fmt.Println("Done!")
    case <-ctx.Done():
        fmt.Println("Timeout!")
    }
}
```

**Concepts**: WaitGroup, Mutex, Once, Context, timeouts, cancellation

---

## 14: Systems Programming - Signal Handling

```go
package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    // Create channel for signals
    sigChan := make(chan os.Signal, 1)
    signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
    
    fmt.Println("Application running. Press Ctrl+C to exit.")
    
    // Wait for signal
    sig := <-sigChan
    fmt.Printf("Received signal: %v\n", sig)
    fmt.Println("Cleaning up...")
    os.Exit(0)
}
```

**Concepts**: Signal handling, os.Signal, syscall, graceful shutdown

---

## 15: Performance - Benchmarking

```go
package main

import (
    "fmt"
    "testing"
)

func concatenate(strings []string) string {
    result := ""
    for _, s := range strings {
        result += s
    }
    return result
}

func BenchmarkConcatenate(b *testing.B) {
    data := []string{"a", "b", "c", "d", "e"}
    b.ResetTimer()
    
    for i := 0; i < b.N; i++ {
        concatenate(data)
    }
}

// Run with: go test -bench=. -benchmem
```

**Concepts**: Benchmarks, b.ResetTimer, -benchmem flag, profiling

---

## 16: Distributed Systems - gRPC Example

```
# service.proto
syntax = "proto3";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

```go
// server.go
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// client.go
client := pb.NewGreeterClient(conn)
resp, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "World"})
fmt.Println(resp.Message)
```

**Concepts**: Protocol Buffers, gRPC service definition, server/client implementation

---

## 17: Real-World Project - Simple Task Manager

```go
// main.go - combines packages, testing, error handling, concurrency, etc.
package main

import (
    "flag"
    "fmt"
    "mytasks/internal/storage"
    "mytasks/internal/api"
)

func main() {
    var port string
    flag.StringVar(&port, "port", "8080", "Server port")
    flag.Parse()
    
    // Initialize storage
    store := storage.NewMemoryStore()
    
    // Start API server
    server := api.NewServer(store)
    fmt.Printf("Starting server on :%s\n", port)
    server.Start(port)
}
```

**Concepts**: Project structure, multiple packages, testing, concurrency, error handling, HTTP APIs

---

## How to Use These Examples

1. **Read** the example code
2. **Type** it out (don't copy-paste)
3. **Run** it: `go run example.go`
4. **Modify** it: Change values, add features
5. **Break** it intentionally to understand error messages
6. **Test** it: Add assertions, benchmarks

Each example demonstrates the core concept without unnecessary complexity. Use them as:
- Learning reference
- Code patterns to follow
- Starting points for exercises
- Basis for combining concepts

