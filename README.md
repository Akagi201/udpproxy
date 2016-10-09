# udpproxy

A simple UDP Proxy Server in Golang.

## Build
* docker: `docker build -t udpproxy .`
* `go build main.go -o udpproxy`

## Run
* `--source`: default source is `:2203`.
* `--target`: default target is `127.0.0.1:2202`.
* `--quiet`: whether to print logging info or not.
