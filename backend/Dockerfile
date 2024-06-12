#BuildStage
FROM golang:1.22.4-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main ./cmd/api/main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/4.17.1/migrate.linux-amd64.tar.gz | tar xvz
#Run Stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 /usr/bin/migrate
# COPY app.env .
COPY start.sh .
COPY ./internal/db/migrations ./migrations

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]