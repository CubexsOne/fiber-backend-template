FROM --platform=linux/amd64 golang:1.20.5-alpine3.17 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY src src
RUN go build -o backend ./src/main.go

FROM alpine:3.17.0 as runner

WORKDIR /app

COPY --from=builder /app/backend /app/backend

ENTRYPOINT ["sh", "-c", "./backend"]