#!/bin/bash

echo "Testing Go Backend Setup..."

# Test compilation
echo "1. Testing compilation..."
go build -o tmp/test cmd/server/main.go
if [ $? -eq 0 ]; then
    echo "✓ Compilation successful"
    rm -f tmp/test
else
    echo "✗ Compilation failed"
    exit 1
fi

# Test go mod
echo "2. Testing go modules..."
go mod tidy
if [ $? -eq 0 ]; then
    echo "✓ Go modules OK"
else
    echo "✗ Go modules failed"
    exit 1
fi

# Test basic structure
echo "3. Testing project structure..."
required_dirs=("internal/config" "internal/models" "internal/middleware" "cmd/server" "migrations")
for dir in "${required_dirs[@]}"; do
    if [ -d "$dir" ]; then
        echo "✓ $dir exists"
    else
        echo "✗ $dir missing"
        exit 1
    fi
done

echo "✓ All tests passed! Backend setup is ready."