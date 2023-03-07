FROM golang:1.20-alpine as builder
LABEL stage=go-builder

# Create unprivileged user
ENV USER=appuser
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o go_server main.go

RUN chmod +x /app/go_server

FROM scratch

COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable
COPY --from=builder /app/go_server /app/go_server

# Use unprivileged userd user
USER appuser:appuser

CMD ["/app/go_server"]