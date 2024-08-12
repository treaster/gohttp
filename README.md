# gohttp

## Overview
`gohttp` is a VERY dumb HTTP server, implemented in Go around its built-in HTTP package. It's intended for serving directories of static content during development. Additionally, it includes a couple of bash scripts for common use cases: generating TLS certs suitable (only) for a local development server, and wrapping the gohttp binary in a very simple bash script to cause it to serve a specified directory with HTTPS, auto-generating certs if necessary.

## Usage
This will serve HTML files in `some_dir` on https://localhost:4443.
```
./simple_serve.sh [some_dir]
```
