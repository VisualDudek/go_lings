# List all available commands
default:
    @just --list

# Run the application
run dir=invocation_directory():
    cd {{dir}} && go run ./

# Format code
fmt dir=invocation_directory():
    cd {{dir}} && go fmt ./...

# Run go vet
vet dir=invocation_directory():
    cd {{dir}} && go vet ./...