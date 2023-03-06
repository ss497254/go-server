FROM golang:1.20-alpine as builder
LABEL stage=go-builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o go_server main.go

RUN chmod +x /app/go_server

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/go_server /app

CMD ["/app/go_server"]