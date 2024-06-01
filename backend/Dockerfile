FROM golang:1.22.0

WORKDIR /go/src

COPY go* ./
RUN go mod download

COPY . .
