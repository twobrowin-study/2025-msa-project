FROM golang:1.25.5-alpine3.23 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN apk update --no-cache && apk add --no-cache tzdata

USER 1000:1000

WORKDIR /build
ENV GOCACHE=/build/.cache

COPY --chown=1000:1000 go.mod .
COPY --chown=1000:1000 go.sum .

RUN go mod download

COPY --chown=1000:1000 src src

RUN go build -ldflags="-s -w" -o /build/main ./src

FROM alpine:3.23

RUN apk update --no-cache && apk add --no-cache ca-certificates

COPY --from=builder /usr/share/zoneinfo/Europe/Moscow /usr/share/zoneinfo/Europe/Moscow
ENV TZ=Europe/Moscow

USER 1000:1000

WORKDIR /app

COPY --from=builder /build/main /app/main

EXPOSE 8000

CMD ["/app/main"]