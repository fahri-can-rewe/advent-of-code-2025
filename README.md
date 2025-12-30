# Advent of Code 2025

This repository contains my solutions for [Advent of Code 2025](https://adventofcode.com/2025) implemented in Go.

## Project Structure

The project is organized as follows:

```text
├── cmd/
│   ├── day01/
│   │   ├── main.go      # Entry point for Day 1
│   │   └── input.txt    # Your puzzle input
│   ├── day02/
│   │   ├── main.go
│   │   └── input.txt
│   └── ...
├── internal/
│   └── util/
│       └── util.go      # Common logic (e.g., ReadInput)
├── go.mod
└── README.md
```

## How to Run

To run a specific day's solution, navigate to the `cmd/dayXX` directory and run:

```bash
go run main.go
```

Or from the root:

```bash
go run cmd/day01/main.go
```
