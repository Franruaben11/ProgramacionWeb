# Etapa 1: Compilación
FROM golang:1.27-alpine AS builder
WORKDIR /app
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o /app/api .

# Etapa 2: Imagen final (solo el binario y estáticos)
FROM alpine:3.24
RUN adduser -D -u 1000 app
USER app
WORKDIR /app
COPY --from=builder /app/api ./api
COPY --from=builder /app/static ./static
EXPOSE 8080
CMD ["./api"]