# go_lings
My Go learning exercises. Inspired by ZigLings

- Check [ Claude go learning project ](./go-learning-project/README.md) for more structured Go learning path.
- [ **Go mindset and learning strategies based on "Make It Stick"** ](./MINDSETUP.md)
- [ shortlist of Go libraries that are both widely used and excellent for learning practical Go skills](./GO_LIBS.md)
- [ your cheat sheet is Go by Example ](https://gobyexample.com/)

## Anki
Flashcards from various Go learning resources. See md files in [Anki](./Anki) directory.
- [Go Fundamentals - Flashcards](./Anki/golang_tour_flashcards.md)
- [Go Tour More Types - Flashcards](./Anki/001c_tour_more_types_flashcards.md)
- [Go Methods and Interfaces - Flashcards](./Anki/go_methods_interfaces_flashcards.md)
- [Go Concurrency Flashcards](./Anki/golang_concurrency.md)

## Pitfalls
- [Go Struct Pitfalls (Bug-Oriented)](./go-learning-project/03-data-structures/PITFALLS.md)

## Q
- How to ignore staticcheck in VS Code and `go vet`? Is it even possible?
- What is `./...` pattern in Go?
    - `...` is a wildcard pattern in Go package matching syntax
- does go have formatter? How to setup max. line length?
    - `gofmt` is the standard code formatter for Go. It does not have a max line length setting.
    - try `golangci-lint` with `lll` (line length limit) linter enabled. [ Golangci-lint ](https://golangci-lint.run/)
- What is `go.mod` file?
- Can set Go version with `go.mod` file?
    - Yes, you can specify the Go version in the `go.mod` file using the `go` directive.
- Underlying type for `time.Weekday`?
    - The underlying type for `time.Weekday` is `int`.
- What is diff between `fmt.Println` and `println()`?
    - `fmt.Println` is a formatted output function from the `fmt` package, while `println()` is a built-in function in Go for simple printing. `fmt.Println` provides more formatting options and is generally preferred for production code.
    - `fmt.Println` return number of bytes written and error, while `println()` does not return any value.
    - `fmt.Println` support Interface types.
- `var f func(func(int,int) int, int) int` - what is this declaration?
    - This declaration defines a variable `f` that is a function. This function takes two parameters: the first parameter is itself a function that takes two `int` arguments and returns an `int`, and the second parameter is an `int`. The outer function returns an `int`.
- can I use "named return" in func. signature with multiple return values?
    - use "naked return" only for small functions, otherwise it reduces code readability.
- What is "guard clauses"?
- `fmt.Sprint` very interesting: "Spaces are added between operands when neither is a string." (docs)
- Anonymous Functions
- Closures return function, so you need first call outer function to get inner function, then call inner function.
    - aha-moment: EXAMPLE: Fibonacci closure
- Why do I need nil slices ?
- I need `make()` for Pre-allocating memory for slices, without it I need to append to nil slice which cause multiple allocations -> wasteful bc. of copying data to new memory locations multiple times.
- `range` over int., one value is returned, iterate over slice, two values are returned (index, value). You can skip second var with `_` or omit it.
- before using `map` need to initialize it with `make()`, otherwise it is nil and will cause runtime panic on write.
- curly braces `{ }` for literals.
- mind bending, using short var declaration `:=` inside 'for' block loop
```go
for _, v := range slice {
    _, ok := myMap[v] // mind bending
}
```
- Function Literals
- You can declare a method on non-struct types, too. (!!!)
- In Go you pass by value not by reference. This has consequences when passing by value structs that you expect to be modified inside the method you get copy of struct. **And that is why pointer receivers are used to modify the original struct.**
- Check memory layout of struct with reflect package.
```go
typ := reflect.TypeOf(stats{})
fmt.Printf("Struct is %d bytes\n", typ.Size())
```
- What is `unsafe` package?
- Empty struct
- sugar syntax logic: you can call method with pointer receiver on value type, Go will take address automatically.
- interface variable
    - interface variable can hold any value that implements the interface
    - the deal is that interface variable holds two things: the actual value and the type information about that value which is different from the value itself.
- think about interface under the hood as a tuple of (type, value)
- empty interface `interface{}` act as a generic type that can hold any value aka any type.
- `any` is type alias for `interface{}`
- if you want to emphasize that an interface relates to a type or method, add `-er` suffix to the interface name. e.g. `Reader`, `Writer`, `Stringer`
- [ EPIC `Stringer()` ](./001d_tour_methods_interfaces/018_epic_bug/main.go)
- Best practice: Use a pointer receiver for error types to avoid unnecessary copying and because errors often need to be returned as pointers from functions.
    - if you use value receiver for error type, you may encounter issues when trying to compare errors to `nil` bc. it will always be non-nil due to the way interfaces work in Go. -> zero value e.g. `string` will be `""` and not `nil`.
- Why `strings.NewReader` return pointer to struct `*strings.Reader` instead of value `strings.Reader`?
    - because `strings.Reader` has methods with pointer receivers, so to use those methods you need a pointer to the struct.
- Return bytes by args slice, not by return - this is mind bending and very efficient.
```go
Read([]byte) (int, error) // Note that return value to not contain []byte
```
- based on `Reader` **so brilliant design**, 
    - using named return values to return number of bytes read and error only, while the actual data is returned via the byte slice passed as argument. 
    - by using named return values in branch where no error occurs, just init value and do naked return, which implicit return nil for error. (**brilliant!**)
- if a struct has a field that is `io.Reader`:
    - then the struct automatically implements the `io.Reader` interface. This is because the struct can delegate the `Read` method call to its embedded `io.Reader` field. (???)
    - you can use method `Read` on this field
- Empty interface `interface{}` is always implemented by every type because it has no methods requirements.
- Name your interfaces parameters:
```go
type Copier interface {
    Copy(dst string, src string) (written int64, err error)
}
// is better than
type Copier interface {
    Copy(string, string) (n int64, err error)
```
- Keep interfaces small.
- chain error checks with `if err != nil { return err }` this is brilliant design.
```go
func doSomething() (int, error) {
    result, err := doAnotherThing()
    if err != nil {
        return 0, err
    }
    return result, nil
}
```
- custom error with `fmt.Errorf("custom error: %v", int)` or better `errors.New("custom error")` 
- Anti go pattern: using defer/recover for regular error handling is an anti-pattern in Go. Defer/recover should be reserved for truly exceptional situations, not for routine error handling.
- heal your future slice pain with rule of thumb:
    - append on the same slice the result is assigned to.
- TODO: add slice pitfalls: when you iterate over slice of structs with `for _, v := range slice` you get copy of struct, so modifying `v` does not modify original struct in slice -> need to use index to modify original struct in slice.
- good to know `strings.Contains(s, substr string) bool` and `strings.ToLower(s string) string`
- maps are passed by reference into functions.
- Go idiom: `if _, ok := myMap[key]; ok { // key is found }` use it with `ok` or `!ok` depending on your logic.
- create set with map using bool values: `mySet := make(map[string]bool)` and Python`in` logic: `if mySet[element] { // element is in the set }`
- or use empty struct as value to save memory: `mySet := make(map[string]struct{})` and check existence with: `if _, ok := mySet[element]; ok { // element is in the set }`
- A dir of Go code can have st most one  package. All `.go` files in that dir must declare the same package name.
- `go install` in your repo go module will install binaries to `$GOPATH/bin` or `$GOBIN` if set. -> will build and install the package with dir name.
- if there is not `main.go` file in the package, `go build` will build the package but not produce an executable binary, instead it will produce file which is in local cache that can be imported and used by other packages. Where is this cache located?
    - The compiled package files are stored in the Go build cache, which is typically located in the `$GOPATH/pkg/mod/cache` directory. (???)
- in `go.mod` `replace` directive for local module replacement during development.
- different dir == different package
- Clean package:
    - Hide Internal Logic -> design package API to expose only what is necessary.
    - Do not change API
    - Do not export func. from the Main package
    - Package should know about dependencies
- channels data flow notation `<-`:
    - `<- channel` receive from channel
    - `channel <- value` send to channel
    - just check at which side of notation is channel to understand direction.
- ohohoh `time.Tick` returns a channel that delivers the time.
- Web solution for "Binary Trees" are sooo different to each other (!!!)
- inside `func Walk()` using anonymous fn. assigned to var. named `walk` is brilliant design.
- What is idiomatic way to wait in main fn. for goroutines to finish?
    - Use `sync.WaitGroup` to wait for multiple goroutines to finish before exiting the main function. This allows you to coordinate the completion of concurrent tasks without blocking indefinitely or risking a deadlock.
- "Receiving from a closed channel returns the zero value immediately" but how it is useful? I need example ...
- receive zero value from closed channel is very useful for concurrent tree walk, it allows to signal the end of traversal without needing additional synchronization mechanisms. (???)
- How can I use `time.After()` for timeouts?
    - in select statement, use `case <-time.After(duration):` to trigger a timeout if no other case is ready within the specified duration. 
- Channels are **reference types**, so when you pass a channel to a function, you are passing a reference to the channel, not a copy of it. This means that multiple goroutines can share the same channel and communicate through it without needing to worry about copying or synchronization issues that arise with value types. 
- Channels can be used as signals to coordinate goroutines, for example by closing a channel to signal that no more values will be sent, allowing receiving goroutines to exit gracefully. `c <- struct{}{}` is a common pattern for sending a signal without any data. Read "signal channel" `<-c` to wait for the signal.
- How the hell does this work?
    - Why I can pass into processData() the result of downloadData() which is a channel that is created inside downloadData() and used inside the goroutine in downloadData()?
```go
func downloadData() chan struct{} {
	downloadDoneCh := make(chan struct{})

	go func() {
		fmt.Println("Downloading data file...")
		time.Sleep(2 * time.Second) // simulate download time

		// after the download is done, send a "signal" to the channel
		downloadDoneCh <- struct{}{}
	}()

	return downloadDoneCh
}

func processData(downloadDoneCh chan struct{}) {
	// any code here can run normally
	fmt.Println("Preparing to process data...")

	// block until `downloadData` sends the signal that it's done
	<-downloadDoneCh

	// any code here can assume that data download is complete
	fmt.Println("Data download complete, starting data processing...")
}

processData(downloadData())
// Preparing to process data...
// Downloading data file...
// Data download complete, starting data processing...
```
- What does it mens ?"sending on a buffered channel blocks only when the buffer is full, and receiving blocks only when the buffer is empty." 
    - For a buffered channel, you can send values into the channel without blocking until the buffer reaches its capacity. Once the buffer is full, any further sends will block until space is available (i.e., until a value is received from the channel). Conversely, receiving from a buffered channel will block only if the buffer is empty, meaning there are no values to receive. This allows for more flexible communication between goroutines, as they can proceed without blocking as long as the buffer has space or values to receive.
- Read-only, write-only channels: `ch <-chan int` is a read-only channel, `ch chan<- int` is a write-only channel. This is useful for function parameters to indicate the intended use of the channel and to prevent accidental misuse.
- `sync.Mutex` is NOT a reference type, do not copy it, pass pointer to it. 
    - check Usage Patterns.
- for read-intensive processes, use `sync.RWMutex` which allows multiple readers but only one writer at a time.
- why generics exist in Go? to avoid following code duplication:
```go
func splitIntSlice(s []int) ([]int, []int) {
    mid := len(s) / 2
    return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
    mid := len(s) / 2
    return s[:mid], s[mid:]
}

// with generics you can have:
func splitSlice[T any](s []T) ([]T, []T) {
    mid := len(s) / 2
    return s[:mid], s[mid:]
}
```
- Zero value of a Generic type `var myZero T`
- huge leverage of generics to use as type interface such as `func concat[T stringer](vals []T) string {}`
- "syntax sugar" for type constraints -> use interface type list (type set) to specify multiple types that satisfy the constraint, instead of defining separate interfaces for each type.
```go
type Ordered interface {
    ~int | ~float64 | ~string
}

func Min[T Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}
```
- Interface definition can accept type parameters as well.
- What Golang is missing?
    - Enums e.g. Rust Result type
    - union types
- Golang "union type" broken solution:
```go
type perm string

const (
    Read perm = "read"
    Write perm = "write"
    Execute perm = "execute"
)

func checkPermission(p perm) {}
```
- (???) `cURL` is a cmd tool for transferring data using **various protocols**.
- `curl https://jsonplaceholder.typicode.com/users/1` 
- `jq` Basics, get:
 - name field: `jq '.name'`
 - field from each element in array: `jq '.[].name'`
 - multiple fields: `jq '.name, .email'`

### encoding/json
- `json.NewDecoder(r io.Reader) *json.Decoder` creates a new JSON decoder that reads from the provided `io.Reader`. This allows you to decode JSON data from various sources, such as files, network connections, or in-memory buffers, by passing the appropriate `io.Reader` implementation.
- `func (dec *Decoder) Decode(v interface{}) error` is a method of the `json.Decoder` type that decodes JSON data from the underlying `io.Reader` and stores it in the value pointed to by `v`. The `v` parameter is typically a pointer to a struct or a map where the decoded JSON data will be stored. The method returns an error if the decoding process fails, such as when the JSON data is malformed or does not match the structure of `v`. 
- `func json.Marshal(v interface{}) ([]byte, error)` is a function that takes an input value `v` and encodes it into JSON format. The function returns a byte slice containing the JSON representation.

### net/http
- `type Client struct` is a struct that represents an HTTP client. It has methods for making HTTP requests, such as `Get`, `Post`, etc. The `Client` struct also has fields for configuring the client's behavior, such as timeouts, transport settings, etc. By using a `Client`, you can customize how your HTTP requests are made and handle responses more effectively.
- `func (c *Client) Do(req *http.Request) (*http.Response, error)` 
- `type http.Response struct` have fields such as `Header` and `Header` have methods such as `Get` to access specific header values. 
- simple `http.Get` vs. verbose and more powerful `http.Client` with `Do` method.
- (???) why do I need to call `resp.Body.Close()` after reading the response body? 
- package provide loots of useful constants for HTTP status codes, such as `http.StatusOK`, `http.StatusNotFound`, etc. 

### net/url
- `url.Parse(rawURL string) (*url.URL, error)` 
- `type URL struct` have fields such as `Scheme`, `Host`, `Path`, etc. and methods such as `String()` to get the string representation of the URL.

### bad vs. good
```go
// BAD
package main

func fetchTasks(baseURL, availability string) []Issue {
	query := "?" + "sort=estimate" + "&limit="

	var limit string
	
	switch availability {
	case "Low":
		limit = "1"
	case "Medium":
		limit = "3"
	case "High":
		limit = "5"
	}

	fullURL := baseURL + query + limit
	return getIssues(fullURL)
}
```
```go
// GOOD
package main

import (
	"fmt"
)

func fetchTasks(baseURL, availability string) []Issue {
	amountofIssues := map[string]int{
		"Low":    1,
		"Medium": 3,
		"High":   5,
	}

	fullURL := fmt.Sprintf("%s?sort=estimate&limit=%d", baseURL, amountofIssues[availability])
	return getIssues(fullURL)
}

```

### Boot.dev ideas
- Embedded Structs
- Anonymous Structs
- memory layout of structs
    
    

## SOP
- Use `.justfile`, grab inspiration from `.justfile_go_dev` (AI generated).

### Directory Structure Rules
```
XXX_topic_name/
```
- `XXX`: 3-digit zero-padded number (001, 002, 003...)
- `topic_name`: lowercase with underscores

### reading materials
- [ Go's Declaration Syntax ](https://go.dev/blog/declaration-syntax)
- [ Go Slices: usage and internals ](https://go.dev/blog/slices-intro)
- [ Composite literals in Go ](https://go.dev/ref/spec#Composite_literals)
- [ Go specification ](https://go.dev/ref/spec)
- [ wiki: Closure ](https://en.wikipedia.org/wiki/Closure_(computer_programming))
- [ Effective Go ](https://go.dev/doc/effective_go)
    - maps
- [ go101 memory layout ](https://go101.org/article/memory-layout.html)
- [ Go by example ](https://gobyexample.com/)
- [ Go maps in action ](https://go.dev/blog/maps)
- [ Stack vs. Heap ](https://go.dev/doc/faq#stack_or_heap)
- [ Strings, bytes, runes and characters in Go ](https://go.dev/blog/strings)
- [ Go FAQ ](https://go.dev/doc/faq)
- [ Channel Axioms ](https://dave.cheney.net/2014/03/19/channel-axioms)
- [ The Grug Brained Developer ](https://grugbrain.dev/)
- [ Go Proverbs ](https://go-proverbs.github.io/)
- [ Gopherfest Go Proverbs ](https://www.youtube.com/watch?v=PAAkCSZUG1c)
- [ curl tutorial ](https://curl.se/docs/tutorial.html)
- [ Tutorials ](https://go.dev/doc/tutorial/)
- [ TDD done correctly ](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests) -> table-driven tests

## Postmortem

Why when running `justfile vet` recipe I got output with error:
```
cd /home/m/workspace/go_lings/001a_vet/001_usecase && go vet ./...
# go_lings/001a_vet/001_usecase
# [go_lings/001a_vet/001_usecase]
./main.go:25:8: assignment copies lock value to c2: go_lings/001a_vet/001_usecase.Counter contains sync.Mutex
./main.go:21:20: fmt.Printf format %d has arg name of wrong type string
./main.go:30:2: unreachable code
error: Recipe `vet` failed on line 18 with exit code 1
```
while running `go vet ./...` directly from terminal does not throw this last line with `error`?

Takeaway: justfile reports vet recipe exit code which is nonzero.