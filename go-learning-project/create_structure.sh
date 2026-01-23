#!/bin/bash

# Go Learning Project Structure Generator
# This script creates a complete directory structure for learning Go progressively

set -e

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

PROJECT_ROOT="${1:-.}"

echo -e "${BLUE}Creating Go Learning Project Structure${NC}"
echo "Root directory: $PROJECT_ROOT"
echo ""

# Create main directories
echo -e "${YELLOW}Creating main directories...${NC}"

DIRS=(
    "01-fundamentals/exercises/solutions"
    "02-functions/exercises"
    "03-data-structures/exercises"
    "04-interfaces/exercises"
    "05-concurrency/exercises/benchmarks"
    "06-error-handling/exercises"
    "07-packages-modules/package_basics"
    "07-packages-modules/internal_packages/cmd/app"
    "07-packages-modules/internal_packages/internal/config"
    "07-packages-modules/module_example"
    "07-packages-modules/vendor_example"
    "08-io-files/exercises/data"
    "09-testing/exercises/coverage"
    "10-web-basics/exercises/tests"
    "11-databases/exercises/migrations"
    "11-databases/tests"
    "12-reflection/exercises"
    "13-advanced-concurrency/exercises/benchmarks"
    "14-systems-programming/exercises"
    "15-performance/exercises/profiles/scripts"
    "16-distributed-systems/exercises/proto"
    "16-distributed-systems/tests"
    "17-real-world-projects/project-1-cli-tool/cmd"
    "17-real-world-projects/project-1-cli-tool/internal/storage"
    "17-real-world-projects/project-1-cli-tool/internal/cli"
    "17-real-world-projects/project-2-web-api/internal/api"
    "17-real-world-projects/project-2-web-api/internal/db"
    "17-real-world-projects/project-2-web-api/internal/middleware"
    "17-real-world-projects/project-2-web-api/migrations"
    "17-real-world-projects/project-2-web-api/tests"
    "17-real-world-projects/project-3-system-monitor/internal/metrics"
    "17-real-world-projects/project-3-system-monitor/internal/output"
    "17-real-world-projects/project-4-distributed-cache/cmd"
    "17-real-world-projects/project-4-distributed-cache/proto"
    "17-real-world-projects/project-4-distributed-cache/internal/cache"
    "17-real-world-projects/project-4-distributed-cache/internal/replication"
    "17-real-world-projects/project-4-distributed-cache/internal/rpc"
    "17-real-world-projects/project-5-task-scheduler/internal/scheduler"
    "17-real-world-projects/project-5-task-scheduler/internal/storage"
    "17-real-world-projects/project-5-task-scheduler/internal/api"
)

for dir in "${DIRS[@]}"; do
    mkdir -p "$PROJECT_ROOT/$dir"
    echo -e "${GREEN}✓${NC} $dir"
done

echo ""
echo -e "${YELLOW}Creating placeholder files...${NC}"

# Create placeholder README files for each major directory
for i in {01..17}; do
    dir=$(ls -d "$PROJECT_ROOT"/${i}-* 2>/dev/null | head -1)
    if [ -d "$dir" ]; then
        cat > "$dir/README.md" << 'EOF'
# Topic: $(basename $(pwd))

## Learning Objectives

- Objective 1
- Objective 2
- Objective 3

## Key Concepts

### Concept 1
Explanation here.

### Concept 2
Explanation here.

## Exercises

See `exercises/` directory for practice problems.

## Further Reading

- [Official Go Documentation](#)
- [Effective Go](#)

EOF
        echo -e "${GREEN}✓${NC} README.md in $(basename $dir)"
    fi
done

# Create example .gitignore
cat > "$PROJECT_ROOT/.gitignore" << 'EOF'
# Binaries
*.o
*.a
*.so
.o

# Test binaries
*.test

# Output files
*.out
*.log
*.prof
*.pprof

# Dependency directories
vendor/

# Go module files (optional, usually commit these)
# go.sum
# go.mod

# IDE
.vscode/
.idea/
*.swp
*.swo
*~
.DS_Store

# Build directories
build/
dist/
bin/

# SQLite databases (if using)
*.db
*.sqlite
*.sqlite3
EOF
echo -e "${GREEN}✓${NC} .gitignore created"

# Create top-level Makefile for convenience
cat > "$PROJECT_ROOT/Makefile" << 'EOF'
.PHONY: help test bench clean structure

help:
	@echo "Go Learning Project"
	@echo "=================="
	@echo "make test      - Run all tests"
	@echo "make bench     - Run benchmarks"
	@echo "make clean     - Clean generated files"
	@echo "make structure - Display project structure"

test:
	@echo "Running tests..."
	@for dir in {01..16}-*; do \
		if [ -d "$$dir" ]; then \
			echo "Testing $$dir..."; \
			(cd $$dir && go test -v ./... 2>/dev/null || true); \
		fi \
	done

bench:
	@echo "Running benchmarks..."
	@for dir in {01..16}-*/*bench* {13..16}-*; do \
		if [ -f "$$dir"/*_test.go ]; then \
			echo "Benchmarking $$dir..."; \
			(cd $$(dirname $$dir) && go test -bench=. -benchmem ./... 2>/dev/null || true); \
		fi \
	done

clean:
	@find . -name "*.out" -delete
	@find . -name "*.test" -delete
	@find . -name "*.prof" -delete
	@rm -rf vendor/
	@echo "Cleaned build artifacts"

structure:
	@echo "Project Structure:"
	@echo ""
	@tree -L 2 -I '__pycache__|*.pyc' . || find . -maxdepth 2 -type d | sort

fmt:
	@echo "Formatting Go code..."
	@gofmt -s -w .
	@goimports -w .

lint:
	@echo "Linting Go code..."
	@golangci-lint run ./...
EOF
echo -e "${GREEN}✓${NC} Makefile created"

# Create a comprehensive learning guide
cat > "$PROJECT_ROOT/LEARNING_GUIDE.md" << 'EOF'
# Go Learning Guide

## How to Use This Project

### 1. Read First
- Start with the README.md in each directory
- Understand the concepts before writing code
- Review existing examples

### 2. Run Examples
```bash
cd 01-fundamentals
go run variables_constants.go
```

### 3. Modify & Experiment
- Change example code
- Add print statements
- Break things intentionally to understand behavior

### 4. Complete Exercises
- Follow the exercise descriptions
- Write your solution
- Check solutions in `solutions/` directory

### 5. Test Your Code
```bash
go test -v ./...
```

### 6. Run Benchmarks (for performance chapters)
```bash
go test -bench=. -benchmem ./...
```

## Tips for Success

### Understanding Errors
Go's error messages are very specific. Read them carefully:
```
undefined: fmt
```
This tells you exactly what's missing.

### Debugging
Use:
- `fmt.Printf()` for quick debugging
- `log` package for structured logging
- `pprof` for performance analysis (chapter 15)

### Testing as Learning
Look at the `_test.go` files in each directory:
- They show how the code is meant to be used
- They demonstrate edge cases
- They serve as documentation

### Concurrency Visualization
For chapters 5 and 13, use:
```bash
# Create a trace
go test -trace=trace.out

# View trace (in browser)
go tool trace trace.out
```

### Performance Analysis
Use pprof to understand your code:
```bash
go test -cpuprofile=cpu.prof -memprofile=mem.prof
go tool pprof cpu.prof
```

## Common Mistakes to Avoid

1. **Goroutine Leaks**: Always ensure goroutines exit
2. **Unbuffered Channel Deadlocks**: Understand channel blocking
3. **Nil Pointer Dereferences**: Check for nil before using
4. **Race Conditions**: Use `go test -race` to detect them
5. **Forgetting Defer**: Use defer for cleanup in database/file operations

## Recommended Reading Order

```
Week 1-2:   01-fundamentals, 02-functions, 03-data-structures
Week 3-4:   04-interfaces, 05-concurrency
Week 5-6:   06-error-handling, 07-packages-modules
Week 7-8:   08-io-files, 09-testing
Week 9-10:  10-web-basics, 11-databases
Week 11-12: 12-reflection, 13-advanced-concurrency
Week 13-14: 14-systems-programming, 15-performance
Week 15-16: 16-distributed-systems
Week 17+:   17-real-world-projects
```

## Setting Up Your Environment

### Required
```bash
# Install Go 1.21 or later
go version

# Initialize module (from project root)
go mod init github.com/yourusername/go-learning
```

### Optional but Recommended
```bash
# Code formatting
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Import organization
go install golang.org/x/tools/cmd/goimports@latest

# Go doc server
go install golang.org/x/tools/cmd/godoc@latest
```

## Quick Reference

### Running Code
```bash
go run main.go                    # Run directly
go build                          # Create binary
go build -o myapp                 # Named output
```

### Testing
```bash
go test ./...                     # All tests
go test -v ./...                  # Verbose
go test -race ./...               # Check for race conditions
go test -cover ./...              # Coverage report
go test -coverprofile=c.out ./... # Detailed coverage
go tool cover -html=c.out         # View HTML coverage
```

### Formatting & Linting
```bash
gofmt -s -w .                     # Format (simple)
goimports -w .                    # Organize imports
golangci-lint run ./...           # Comprehensive lint
```

### Performance
```bash
go test -bench=. ./...            # Run benchmarks
go test -cpuprofile=cpu.prof ./...  # CPU profile
go tool pprof cpu.prof            # Analyze profile
```

### Debugging
```bash
GODEBUG=gctrace=1 go run main.go  # GC info
go run -race main.go              # Race detector
```

## Projects Integration

When you reach chapter 17, build complete applications combining:
- Proper package structure (07)
- Concurrency patterns (05, 13)
- Error handling (06)
- Testing (09)
- Database/persistence (11)
- HTTP APIs or CLI (10, 14)
- Performance considerations (15)

Each project teaches you how concepts work together in real applications.

EOF
echo -e "${GREEN}✓${NC} LEARNING_GUIDE.md created"

echo ""
echo -e "${BLUE}Project structure created successfully!${NC}"
echo ""
echo -e "${YELLOW}Next steps:${NC}"
echo "1. cd $PROJECT_ROOT"
echo "2. go mod init github.com/yourusername/go-learning"
echo "3. Read LEARNING_GUIDE.md"
echo "4. Start with 01-fundamentals/README.md"
echo ""
