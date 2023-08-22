# Stage 1: Build the Go binary
FROM golang:1.21-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

ENV GOPROXY=https://proxy.golang.org,direct
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]

# Stage 2: Create a minimal runtime image
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
