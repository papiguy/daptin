#!/usr/bin/env bash

export GOPATH=/media/artpar/ddrive/workspace/newgocode
rice embed-go
go build  -ldflags '-linkmode external -extldflags -static -w' main.go