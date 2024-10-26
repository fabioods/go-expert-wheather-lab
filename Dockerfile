# Etapa de build
FROM golang:1.23.2-alpine AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o weather ./cmd

# Etapa de execução
FROM alpine:latest
WORKDIR /app
COPY --from=build /app/weather .
ENTRYPOINT ["./weather"]
