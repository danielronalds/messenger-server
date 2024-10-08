# Builder
FROM golang:1.22.7-bookworm as builder

WORKDIR /app

COPY go.mod ./go.mod
COPY main.go ./main.go
COPY db/ ./db/
COPY resources/ ./resources/
COPY utils/ ./utils/
COPY stores/ ./stores/
COPY security/ ./security/

RUN go mod tidy
RUN go build -o ./main .

# Runner

FROM debian:stable-slim as runner
WORKDIR /app
COPY --from=builder /app/main .
COPY .env ./.env

EXPOSE 8080

CMD [ "./main" ]
