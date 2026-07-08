FROM golang:1.26-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /app/client_api .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/client_api .
COPY --from=builder /app/.env .env
EXPOSE 8000
CMD ["./client_api"]
