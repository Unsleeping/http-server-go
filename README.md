# Go HTTP Server Implementation

[![progress-banner](https://backend.codecrafters.io/progress/http-server/ecc78c2d-3aba-48b3-b3e4-3e08af08235f)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

## Project Description

This project is an implementation of an HTTP/1.1 server in Go, developed as part of the "Build Your Own HTTP server" challenge by CodeCrafters. The server is capable of handling multiple clients and implements various HTTP features including GET and POST requests, file serving, and gzip compression.

## Features

- Basic HTTP/1.1 server implementation
- Handling of GET and POST requests
- File serving capabilities
- Echo API endpoint
- User-Agent API endpoint
- Gzip compression support
- Concurrent client handling

## Technologies Used

- Go 1.19+
- Standard Go libraries (net, http, io, os, etc.)

## Getting Started

1. Clone this repository
2. Ensure you have Go 1.19 or later installed
3. Run the server using: `go run ./app/server.go`

## Usage

The server runs on port 4221 by default.
