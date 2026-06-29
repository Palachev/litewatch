# --- Сборка ---
FROM golang:1.22-alpine AS build
WORKDIR /app
COPY agent/ .
RUN go mod download && go build -o litewatch .

# --- Финальный минимальный образ ---
FROM alpine:latest
COPY --from=build /app/litewatch /litewatch
COPY config.yaml /config.yaml
ENV LITEWATCH_CONFIG=/config.yaml
ENTRYPOINT ["/litewatch"]
