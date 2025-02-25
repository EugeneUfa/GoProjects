FROM golang:1.21 AS builder

WORKDIR /app

# Копируем весь код (если нет go.mod)
COPY . .

# Отключаем модули и компилируем
RUN go env -w GO111MODULE=off && go build -o server .

# Минимальный образ
FROM debian:bullseye-slim

WORKDIR /app

# Копируем бинарник из builder
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]