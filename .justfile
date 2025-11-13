# List all available commands
default:
    @just --list

# Run the application
run dir=invocation_directory():
    cd {{dir}} && go run ./

# Format code
fmt:
    go fmt ./...

# Run go vet
vet:
    go vet ./...