#!/bin/bash

echo "Compiling..."
go build httpServer.go
echo "Running Server..."
./httpServer
