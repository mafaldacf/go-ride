# Go Ride

Just a simple project implemented with Go and gRPC

First project using Go

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

Compile project (optional):

    ./start.sh compile

Run server:

    ./start.sh run server

Run client:

    ./start.sh run client