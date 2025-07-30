# syntax=docker/dockerfile:1
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o forgebase ./cmd/forgebase

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/forgebase ./forgebase
COPY .env.example .
EXPOSE 8080
CMD ["./forgebase"]
