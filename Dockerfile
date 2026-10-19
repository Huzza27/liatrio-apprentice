FROM golang:1.27 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG APP_VERSION=dev

RUN CGO_ENABLED=0 GOOS=linux go build -o app .

FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/app ./app

ARG APP_VERSION=dev
ENV APP_VERSION=${APP_VERSION}

CMD ["./app"]
