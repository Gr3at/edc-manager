FROM ubuntu:24.04 as base
# create non root user
RUN useradd -u 1001 edc-proxy-user
# add custom CA to trust store
# COPY custom-ca-cert.pem /usr/local/share/ca-certificates/custom-ca-cert.crt
RUN apt update && apt install -y ca-certificates && update-ca-certificates

FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/edc-proxy .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/edc-proxy-migrate ./migrate/migrate.go

FROM scratch
COPY --from=base /etc/passwd /etc/passwd
# reason why certs are copied in that location: https://go.dev/src/crypto/x509/root_linux.go
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/bin/edc-proxy /edc-proxy
COPY --from=builder /app/bin/edc-proxy-migrate /edc-proxy-migrate
USER edc-proxy-user
EXPOSE 8080
CMD ["/edc-proxy"]



