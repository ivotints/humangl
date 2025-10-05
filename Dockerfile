# Multi-stage build для оптимизации размера образа
FROM golang:1.21-alpine AS builder

# Установка необходимых пакетов для OpenGL и GLFW
RUN apk add --no-cache \
    gcc \
    musl-dev \
    libx11-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxcursor-dev \
    libxi-dev \
    mesa-dev \
    xorg-server-dev

# Установка рабочей директории
WORKDIR /app

# Копирование go.mod и go.sum для кэширования зависимостей
COPY go.mod go.sum ./

# Загрузка зависимостей
RUN go mod download

# Копирование исходного кода
COPY . .

# Сборка приложения DO go run, not go build
RUN CGO_ENABLED=1 GOOS=linux go build -o humangl ./cmd/humangl 

# Финальный образ для запуска
FROM alpine:latest

# Установка библиотек для работы с OpenGL и X11
RUN apk add --no-cache \
    mesa-dri-gallium \
    libx11 \
    libxrandr \
    libxinerama \
    libxcursor \
    libxi \
    mesa-gl

# Создание пользователя для безопасности
RUN adduser -D -s /bin/sh humangl

# Копирование скомпилированного приложения
COPY --from=builder /app/humangl /usr/local/bin/humangl

# Переключение на пользователя
USER humangl

# Запуск приложения
CMD ["humangl"]