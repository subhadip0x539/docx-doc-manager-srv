FROM golang:1.22.1 AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux make build

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/build/bin ./build/bin

ENV GIN_MODE=release

ENTRYPOINT /app/build/bin