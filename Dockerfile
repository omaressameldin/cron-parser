FROM golang:1.16-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
COPY commands/go.mod commands/go.sum ./commands/
COPY parser/go.mod ./parser/
COPY utils/go.mod ./utils/

RUN go mod download

COPY . .

RUN go build -o cron-parser .

FROM alpine:3.14

WORKDIR /app

COPY --from=build /app/cron-parser .

ENTRYPOINT ["./cron-parser"]
