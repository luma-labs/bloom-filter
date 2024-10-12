# Bloom Filter Implementation
Paper: https://crystal.uta.edu/~mcguigan/cse6350/papers/Bloom.pdf

This project is a Bloom Filter implementation in Go, featuring customizable false positive rates, optimal hashing, and memory-efficient bitset management. The project includes a command-line interface (CLI) and REST/gRPC APIs for interacting with the Bloom filter.

## Table of Contents
- [Bloom Filter Implementation](#bloom-filter-implementation)
  - [Table of Contents](#table-of-contents)
  - [Project Overview](#project-overview)
  - [Features](#features)
  - [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Clone the Repository](#clone-the-repository)
    - [Install Dependenies \& Build Project](#install-dependenies--build-project)
  - [Usage](#usage)
    - [CLI](#cli)
    - [API](#api)
    - [Library Usage](#library-usage)
    - [Benchmarking](#benchmarking)
    - [Testing](#testing)
    - [Project Structure](#project-structure)

---

## Project Overview

A **Bloom Filter** is a space-efficient probabilistic data structure used to test whether an element is a member of a set. It may yield false positives but guarantees no false negatives. This project implements a Bloom Filter in Go with the ability to configure:
- The number of items to store.
- The desired false positive rate.
- Command-line and API-based interfaces for usage.

---

## Features

- **Optimal Bloom Filter Size**: Automatically adjusts the bitset size and number of hash functions based on the number of elements and desired error rate.
- **Command-Line Interface (CLI)**: Interact with the Bloom filter through simple commands.
- **REST & gRPC APIs**: APIs to interact with the Bloom filter remotely.
- **Benchmarking**: Performance benchmarking for testing efficiency under different workloads.
- **Unit and Integration Testing**: Comprehensive testing suite for robustness.

---
## Installation

### Prerequisites
- **Go 1.16+**
- **Make**
- **Protobuf Compiler** (for gRPC API)
- **curl** or **HTTP client** (for testing REST API)

### Clone the Repository
```bash
git clone https://github.com/luma-labs/bloom-filter.git

cd bloom-filter
```

### Install Dependenies & Build Project 
```bash
go mod download 

make build
```
## Usage

### CLI 
```bash
🚧 TODO 
```
### API 
```bash
🚧 TODO 
```
### Library Usage

If you want to use the Bloom Filter as a library in your own Go projects:
1. Import the Package

```bash
import "github.com/luma-labs/bloom-filter/pkg/bloomfilter"
```
2. Example Usage

```go
filter := bloomfilter.Create(100, 0.01) // 100 elements, 1% error rate
filter.Add([]byte("alice"))
filter.Add([]byte("bob"))
filter.Add([]byte("charlie"))

filter.Has([]byte("alice")) // true
filter.Has([]byte("dave")) // false

```

### Benchmarking
```bash
🚧 TODO
```
### Testing
```bash
🚧 TODO
```

### Project Structure
```
api/ (todo)
   - grpc/
      -- bloomfilter.proto    # gRPC protocol definitions
      -- server.go            # gRPC server implementation
   - rest/
      -- handler.go           # REST API handlers
      -- router.go            # REST router setup

benchmarks/ (todo)
  - bloomfilter_bench_test.go  # Benchmark tests for Bloom filter

cmd/
  - bloomfilter/               # Main application entry point
    -- main.go
  - bloomfilter-cli/           # CLI-specific entry point
    -- main.go

docs/                          # Project documentation

internal/
   - bits/
      -- bitset.go             # Bitset implementation
   - hash/
     -- hash.go                # Hashing functions
   - utils/
      -- util.go               # Utility functions

pkg/
   - bloomfilter/
     -- bloomfilter.go         # Bloom filter logic
     -- bloomfilter_test.go    # Bloom filter unit tests

tests/                         # Integration and e2e tests (todo)

Makefile                       # Build commands and tasks
go.mod                         # Go modules definition
go.sum                         # Dependency checksum

```