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

# Run in the highest numbered subfolder (e.g., 014_type_inference)
run-latest:
    cd {{invocation_dir}} && cd $(ls -1 | grep '^[0-9][0-9][0-9]_' | sort -rV | head -1) && go run ./

tour:
    ~/go/bin/tour
