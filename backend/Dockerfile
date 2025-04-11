# ./backend/Dockerfile
# FROM golang:1.24-alpine3.18 AS builder
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -o /go-todo-app

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go-todo-app /app/go-todo-app
EXPOSE 8080
CMD ["/app/go-todo-app"]