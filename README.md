# udpproxy

A simple UDP Proxy Server in Golang.

## Features
- [x] one source, multi target based on copy.

## Build
* docker: `docker build -t udpproxy .`
* `go build main.go -o udpproxy`

## Run
* `--source`: data source, default source is `:2203`.
* `--target`: data target, e.g. `ip:port`.
* `--quiet`: whether to print logging info or not.
* `--buffer`: default is 10240.
