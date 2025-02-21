FROM golang:1.24 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

COPY .env .env

RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 \
     go build -ldflags="-linkmode external -extldflags -static" -o /app/main ./cmd/web/

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

RUN apk add --no-cache libc6-compat gcompat
RUN chmod +x main

EXPOSE 8080

CMD ["./main"]