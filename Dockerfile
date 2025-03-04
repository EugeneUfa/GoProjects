FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

# Отключаем модули и компилируем
RUN go env -w GO111MODULE=off && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build>

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]