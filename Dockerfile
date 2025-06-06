# Используем официальный образ Go как базовый
FROM golang:1.24.3-alpine

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum отдельно — так Docker будет кэшировать зависимости
COPY go.mod ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Собираем приложение
WORKDIR /app/internal
RUN go build -o main .

# Указываем команду, которую нужно выполнить при запуске контейнера
CMD ["./main"]
