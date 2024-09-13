#!/usr/bin/env bash

# Build Server
cd server || exit
go test -v ./...
