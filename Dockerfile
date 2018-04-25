FROM golang:latest

RUN export GOPATH=/go
RUN export PATH=$PATH:/go/bin

