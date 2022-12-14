FROM golang:1.16-alpine as builder
WORKDIR /app
ADD go.mod ./
RUN go mod download
ADD ./ ./
RUN CGO_ENABLED=0 go build -o server main.go
ENTRYPOINT "/app/server"
