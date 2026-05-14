# Etapa 1: Compilación
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

# Etapa 2: Ejecución (Imagen mínima)
FROM scratch
WORKDIR /
COPY --from=builder /app/main .
COPY --from=builder /app/static ./static
EXPOSE 9000
CMD ["./main"]