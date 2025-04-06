# 🧠 Go TP

A collection of small projects and technical exercises written in Go, covering topics like concurrency, gRPC, memory analysis, benchmarks, and more.

## 📦 Project Structure

Each folder contains an isolated example or technical exercise:

- `benchmark/` – performance benchmarking
- `bruteforce/` – simple brute force algorithms
- `channel/`, `channels/`, `waitgroup/`, `semaphore/`, `pipeline/`, `select/` – Go concurrency patterns
- `context/` – usage of Go's `context` package
- `escape-analysis/` – understanding memory allocations
- `faninfanout/`, `goroutine/`, `tradingWaitGroup/` – concurrency fan-in/fan-out patterns
- `grpc/` – basic gRPC example
- `matrix/` – matrix manipulation and computation
- `memoization/`, `memory-layout/` – optimization and memory study
- `pprof/` – performance profiling
- `rle/` – run-length encoding
- `syraucuse/` – Syracuse (Collatz) sequence
- `tp/` – main TP logic or wrapper
- `unittest/` – unit testing examples

## 🚀 Getting Started

### Install dependencies

```bash
go mod tidy
```

### Run the game in the projet folder

```bash
go run main.go
```

### Run benchmarks (if available)

```bash
go test -bench=. -benchmem
```

### Build any project

```bash
go build -o mybinary
```
