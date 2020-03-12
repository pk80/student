#!/bin/sh

# remove folder gen
rm -rf gen
# make a new folder gen
mkdir gen
# generate server without main in gen folder
swagger generate server -A person -t gen -f ./api.yml
# get all necessary files
go get ./...
# format the code
go fmt ./...

# go build cmd

# go run cmd/main.go

