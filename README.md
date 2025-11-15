# go_lings
My Go learning exercises. Inspired by ZigLings

## Q
- How to ignore staticcheck in VS Code and `go vet`? Is it even possible?
- What is `./...` pattern in Go?
    - `...` is a wildcard pattern in Go package matching syntax
    

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