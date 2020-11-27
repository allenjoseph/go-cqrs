# builder
FROM golang:alpine AS builder
RUN apk add --no-cache gcc g++ make ca-certificates git

WORKDIR /go/src/go-cqrs
COPY db db
COPY messaging messaging
COPY model model
COPY util util
COPY woof-service woof-service
COPY query-service query-service
COPY client client

WORKDIR /go/src/go-cqrs
RUN go install -v ./util ./model
RUN go get -d -v ./db ./messaging
RUN go install -v ./db ./messaging
RUN go get -d -v ./woof-service ./query-service ./client
RUN go install -v ./woof-service ./query-service ./client

# runner
FROM alpine
WORKDIR /usr/bin
COPY --from=builder /go/bin .

