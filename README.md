# tricks
Golang compilation of tricks I picked up in software development. 

# Running 
```bash
    go run src/main.go
```

# Testing
This should run all tests in current directory and all of its subdirectories:
```bash
    go test ./...
```

This should run all tests for given specific directories:
```bash
    go test ./tests/... ./unit-tests/... ./my-packages/...
```

Generate Test Coverage
```bash
    go tool cover -html=<output_filename>
```