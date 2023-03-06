#!/bin/bash

if [ "$#" -eq 1 ] && [ $1 = "compile" ]
  then
    echo "generating protoc files from proto/ride.proto..."
    protoc --go_out=. --go-grpc_out=. proto/ride.proto
    echo "done!"

elif [ "$#" -eq 2 ] && [ $1 = "run" ] && [ $2 = "server" ]
  then
    cd server
    go build
    ./server

elif [ "$#" -eq 2 ] && [ $1 = "run" ] && [ $2 = "client" ]
  then
    cd client
    go build
    ./client

else
    echo "Invalid arguments!"
    echo "(1) ./start.sh compile"
    echo "(2) ./start.sh run server|client"
    exit 1
fi