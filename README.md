# Go Ride

A simple project implemented in Go and using gRPC

## Requirements

- **Golang** programming language
- **gRPC**: framework of remote procedure calls that supports client and server communication
- **Protobuf**: cross-platform data used to serialize structured data

## Getting Started

Install tool for managing Go source code:

    sudo apt install -y golang-go
    sudo go version

Install Protocol Buffer Compiler ([Installation](https://grpc.io/docs/protoc-installation/)):

    apt install -y protobuf-compiler
    protoc --version

## Running the Project

Generate new keys and certificates for TLS (done by default):

    ./genKeys.sh

Compile project (done by default):

    ./start.sh compile

Run server:

    ./start.sh run server

Run client:

    ./start.sh run client