# conway-life-go
### An implementation of Conway's Game of Life written in Golang utilizing Ebitengine.

[![Go Version](https://img.shields.io/badge/go-%3E%3D1.20-blue)](https://golang.org)
[![Ebiten Version](https://img.shields.io/badge/ebiten-%3E%3Dv2.8-green)](https://ebiten.org/)
[![License](https://img.shields.io/badge/license-MIT-lightgrey)](/LICENSE)


### Features:
- [x] WASM Compilation
- [x] Fast Ticks
- [x] Hardware Acceleration

## Prerequisites

Before building make sure you have **Go** installed. 
Install from [https://go.dev/dl](https://go.dev/dl).

To verify that Go is installed:
```bash
go version
```

## Getting the Code

Clone the repo and enter it's directory:
```bash
git clone https://github.com/somewhat9/conway-life-go.git
cd conway-life-go
```

## Build

Fetch dependencies:
```bash
go mod tidy
```

Compile:
```bash
go build -o bin/conway-life-go ./cmd/conway-life-go
```

## Run

Run directly (without binary):
```bash
go run ./cmd/conway-life-go/main.go
```

Run from the built binary:
```bash
./bin/conway-life-go
```