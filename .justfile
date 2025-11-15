# Global variable for invocation directory
invocation_dir := invocation_directory()

# List all available commands
default:
    @just --list

# Run the application
run:
    cd {{invocation_dir}} && go run ./

# Format code
fmt:
    cd {{invocation_dir}} && go fmt ./...

# Run go vet
vet:
    cd {{invocation_dir}} && go vet ./...