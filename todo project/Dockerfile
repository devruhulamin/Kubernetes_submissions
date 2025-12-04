FROM golang:1.25-alpine AS builder

WORKDIR /usr/src/app

COPY *.mod ./
RUN go mod download

COPY . .

RUN go build -o todoApp ./cmd

FROM alpine:latest

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/todoApp .

CMD ["./todoApp"]
