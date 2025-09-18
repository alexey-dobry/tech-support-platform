FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/service/build

COPY internal/services/auth_service/go.mod .
COPY internal/services/auth_service/go.sum .
RUN go mod download

COPY internal/services/auth_service/configs/config.toml ./config
COPY internal/services/auth_service .
RUN go build -ldflags="-s -w" -o /app/auth_service ./cmd


FROM scratch

WORKDIR /app/auth_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/auth_service /app/auth_service

CMD ["./auth_service"]