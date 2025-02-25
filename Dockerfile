# Используем минимальный образ Golang
FROM golang:1.21 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем файлы проекта
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем бинарный файл
RUN go build -o server .

# Минимальный образ для запуска
FROM debian:bullseye-slim

WORKDIR /app

# Копируем бинарный файл из builder
COPY --from=builder /app/server .

# Открываем порт
EXPOSE 8080

# Запускаем сервер
CMD ["./server"]
