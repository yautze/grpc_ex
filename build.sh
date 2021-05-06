#!/bin/bash

dir=$1

cd $dir

protoc -I=. -I=$GOPATH/src --gogofaster_out=plugins=grpc:. *.proto
