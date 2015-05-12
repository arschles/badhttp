#!/bin/bash
docker run -v $GOPATH:/go -w /go/src/github.com/arschles/badhttp golang:1.4.2 go build
