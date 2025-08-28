FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o main .


FROM alpine:latest

RUN apk add --no-cache netcat-openbsd

WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/docs ./docs

COPY .env.docker .env

COPY docker-entrypoint.sh .
RUN chmod +x docker-entrypoint.sh

EXPOSE 8080

ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["./main"]