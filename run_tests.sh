#!/bin/bash

echo "Running unit tests..."
go test -v

echo ""
echo "Running tests with coverage..."
go test -cover
