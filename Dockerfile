FROM golang:1.15.8-alpine

RUN apk add --no-cache git \
    && go get -u golang.org/x/tour/wc
