FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o logApp .

FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/logApp .

CMD ["./logApp"]