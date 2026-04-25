# udpproxy

A simple UDP Proxy Server in Golang.

## Features
- [x] one source, multi target based on copy.

## Build

**Note:** This project uses Go Modules and requires Go 1.11 or higher.

* Using Go Modules (recommended):
  ```bash
  git clone https://github.com/Akagi201/udpproxy.git
  cd udpproxy
  go build
  ```
* Docker: `docker build -t udpproxy .`
* Legacy GOPATH mode: `go build main.go -o udpproxy`

## Run
* `--source`: data source, default source is `:2203`.
* `--target`: data target, e.g. `ip:port`.
* `--quiet`: whether to print logging info or not.
* `--buffer`: default is 10240.
