# FROM alpine:3.20 AS base
# RUN adduser -u 1001 edc-proxy-user --disabled-password
# RUN adduser -u 1001 edc-proxy-user -s /bin/sh --disabled-password

FROM ubuntu:24.04 as base
RUN useradd -u 1001 edc-proxy-user


FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/edc-proxy .

FROM scratch
COPY --from=base /etc/passwd /etc/passwd
COPY --from=builder /app/bin/edc-proxy /edc-proxy
USER edc-proxy-user
EXPOSE 8080
CMD ["/edc-proxy"]



