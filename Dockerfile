# builder
FROM golang:alpine AS builder
RUN apk add --no-cache gcc g++ make ca-certificates git

WORKDIR /go/src/github.com/allenjoseph/go-cqrs
COPY db db
COPY messaging messaging
COPY model model
COPY util util
COPY woof-service woof-service
COPY query-service query-service
COPY client client

WORKDIR /go/src/github.com/allenjoseph/go-cqrs
RUN go get -d -v ./...
RUN go install -v ./...

# runner
FROM alpine
WORKDIR /usr/bin
COPY --from=builder /go/bin .

