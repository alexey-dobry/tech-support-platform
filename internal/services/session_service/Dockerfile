FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/service/build

COPY internal/services/session_service/go.mod .
COPY internal/services/session_service/go.sum .
RUN go mod download

COPY internal/services/session_service/configs/config.toml ./config
COPY internal/services/session_service .
RUN go build -ldflags="-s -w" -o /app/session_service ./cmd


FROM scratch

WORKDIR /app/req_manager_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/session_service /app/session_service

CMD ["./session_service"]