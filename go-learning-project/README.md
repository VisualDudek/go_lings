# Go Learning Project Structure

A progressive learning path from Go fundamentals to advanced systems programming concepts.

- [ Structure Guide ](./STRUCTURE_GUIDE.md)
- [ Learning Path Overview ](./LEARNING_PATH.md)
- [ Project Overview ](./PROJECT_OVERVIEW.md)

## Directory Organization

```
go-learning-project/
├── 01-fundamentals/          # Basic syntax, variables, control flow
├── 02-functions/             # Function basics, defer, panic/recover
├── 03-data-structures/       # Arrays, slices, maps, structs
├── 04-interfaces/            # Interface design, composition, polymorphism
├── 05-concurrency/           # Goroutines, channels, synchronization primitives
├── 06-error-handling/        # Error types, custom errors, error wrapping
├── 07-packages-modules/      # Package organization, imports, module management
├── 08-io-files/              # File I/O, buffers, readers/writers
├── 09-testing/               # Unit tests, benchmarks, table-driven testing
├── 10-web-basics/            # HTTP handlers, routing, middleware
├── 11-databases/             # SQL, ORM patterns, transaction handling
├── 12-reflection/            # Type inspection, dynamic behavior
├── 13-advanced-concurrency/  # Worker pools, context patterns, synchronization
├── 14-systems-programming/   # Unix socket, signals, system calls
├── 15-performance/           # Profiling, optimization, memory management
├── 16-distributed-systems/   # RPC, gRPC, load balancing
├── 17-real-world-projects/   # Capstone projects combining all concepts
└── README.md
```

## Learning Path Overview

### Foundation (01-03)
Get comfortable with Go syntax, functions, and basic data structures. These form the building blocks for everything else.

### Core Language Features (04-09)
Master Go's unique features: interfaces, goroutines, error handling, and testing. This is where Go's philosophy becomes clear.

### Practical Development (10-12)
Build actual applications using HTTP, databases, and reflection. Learn how Go concepts apply in real systems.

### Advanced Topics (13-17)
Dive into concurrent systems design, Unix-level programming, performance optimization, and distributed systems architectures.

## How to Use This Structure

1. **Start sequentially** - Each directory builds on previous concepts
2. **Complete exercises** in each directory before moving forward
3. **Reference implementation** - See provided examples for each topic
4. **Experiment** - Modify examples and create variations
5. **Benchmark** - Use performance tools to understand trade-offs

Each directory contains:
- `README.md` - Topic explanation and concepts
- `*.go` - Working examples and exercises
- `_test.go` - Tests demonstrating expected behavior
- `exercises/` - Practice problems with solutions
