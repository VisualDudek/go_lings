# go_lings
My Go learning exercises. Inspired by ZigLings

- Check [ Claude go learning project ](./go-learning-project/README.md) for more structured Go learning path.
- [ shortlist of Go libraries that are both widely used and excellent for learning practical Go skills](./GO_LIBS.md)
- [ your cheat sheet is Go by Example ](https://gobyexample.com/)

## Anki
Flashcards from various Go learning resources. See md files in [Anki](./Anki) directory.
- [Go Fundamentals - Flashcards](./Anki/golang_tour_flashcards.md)
- [Go Tour More Types - Flashcards](./Anki/001c_tour_more_types_flashcards.md)
- [Go Methods and Interfaces - Flashcards](./Anki/go_methods_interfaces_flashcards.md)

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
- [ go101 memory layout ](https://go101.org/article/memory-layout.html)
- [ Go by example ](https://gobyexample.com/)

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