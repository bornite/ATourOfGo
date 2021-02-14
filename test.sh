#!/bin/sh

docker build -t a-tour-of-go .

for s in `ls *_test.go`
do
    echo "-- ${s} --"
    docker run --rm -v `pwd`:/root -w /root a-tour-of-go go test -v ${s} ${s%_test.go}.go
done
