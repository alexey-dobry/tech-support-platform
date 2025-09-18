FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED=0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app/service/build

COPY internal/services/bot_service/go.mod .
COPY internal/services/bot_service/go.sum .
RUN go mod download

COPY internal/services/bot_service/configs/config.toml ./config
COPY internal/services/bot_service .
RUN go build -ldflags="-s -w" -o /app/bot_service ./cmd


FROM scratch

WORKDIR /app/manager_bot_service

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/bot_service /app/bot_service

CMD ["./bot_service"]