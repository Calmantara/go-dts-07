#!/bin/bash

program="go-dts-user"

# ensure all is up to date
go mod tidy

# make artifact for program 
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $program

# make docker
docker build --no-cache	 -t $program -f deployment/build/Dockerfile . 

# clean artifact
rm -rf $program