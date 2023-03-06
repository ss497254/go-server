FROM alpine:latest

WORKDIR /app

ARG file_name="go_server"

COPY /dist/$file_name /app/go_server

EXPOSE 8080

CMD ["./go_server"]
