FROM golang:1.25.5-alpine3.23 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY src src

RUN mkdir -p /app && \
    go build -ldflags="-s -w" -o /app/main ./src

FROM alpine:3.23

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ=Europe/Moscow

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY config config

EXPOSE 8000

CMD ["/app/main"]