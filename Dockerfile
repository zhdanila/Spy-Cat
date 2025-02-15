FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux GOARCH=arm64 go build -o main cmd/server/main.go

FROM alpine:3.18.0 AS final

RUN apk add --no-cache wget bash ca-certificates

WORKDIR /app

COPY --from=builder /app/main ./
COPY --from=builder /app/migrations /app/migrations
COPY --from=builder /app/entrypoint.sh /app/entrypoint.sh

RUN chmod +x /app/entrypoint.sh

ARG TARGETARCH
RUN case "$TARGETARCH" in \
      "amd64"|"") GOOSE_ARCH="amd64";; \
      "arm64") GOOSE_ARCH="arm64";; \
      *) echo "Unsupported architecture: $TARGETARCH; switching to amd64."; GOOSE_ARCH="amd64";; \
    esac && \
    wget -O /bin/goose https://github.com/pressly/goose/releases/download/v3.24.1/goose_linux_${GOOSE_ARCH} && \
    chmod +x /bin/goose

EXPOSE 8080

ENTRYPOINT ["/app/entrypoint.sh"]

CMD ["./main"]