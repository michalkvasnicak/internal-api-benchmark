#!/usr/bin/env bash

protoc --go_out=plugins=grpc:. grpc/*.proto
protoc --go_out=. protobuf/*.proto