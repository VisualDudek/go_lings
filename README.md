# go_lings
My Go learning exercises. Inspired by ZigLings

Check [ Claude go learning project ](./go-learning-project/README.md) for more structured Go learning path.

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
    
    

## SOP
- Use `.justfile`, grab inspiration from `.justfile_go_dev` (AI generated).

### Directory Structure Rules
```
XXX_topic_name/
```
- `XXX`: 3-digit zero-padded number (001, 002, 003...)
- `topic_name`: lowercase with underscores

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