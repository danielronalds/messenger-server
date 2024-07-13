# Builder
FROM golang:1.22.3-bookworm as builder

WORKDIR /app

COPY go.mod ./go.mod
COPY main.go ./main.go

RUN go mod tidy
RUN go build -o ./main .

# Runner

FROM debian:stable-slim as runner
WORKDIR /app
COPY --from=builder /app/main .

EXPOSE 8080

CMD [ "./main" ]
