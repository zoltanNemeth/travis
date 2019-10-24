FROM golang

WORKDIR /travis

COPY . ./

RUN go test
