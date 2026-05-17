FROM golang:1.22-alpine AS builder
WORKDIR /app

COPY . .

RUN go build -o service-binary .

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/service-binary .

EXPOSE 8080

CMD ["./service-binary"]