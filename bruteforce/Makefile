# Makefile for building all Go commands in cmd/*

# Variables
GO_CMD := go
CMD_DIR := cmd
BIN_DIR := bin

# Targets
.PHONY: all clean build

build:
	@echo "Building all sources in cmd..."
	@mkdir -p $(BIN_DIR)
	@go build -o bin/bruteforce_combination cmd/combination/main.go
	@go build -o bin/bruteforce_combination_async cmd/combinationAsync/main.go
	@go build -o bin/bruteforce_async cmd/async/main.go

benchmarks:
	@echo "Running benchmarks..."
	@hyperfine -N ./bin/bruteforce_async  ./bin/bruteforce_combination  ./bin/bruteforce_combination_async

all: build benchmarks

clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)