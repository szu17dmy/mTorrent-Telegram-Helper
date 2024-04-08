FROM golang:1.21.9-bookworm AS builder

WORKDIR /app
ADD / /app

RUN go build -o ./build/target/helper github.com/szu17dmy/mtorrent-telegram-helper/cmd

FROM golang:1.21.9-bookworm

WORKDIR /app
COPY --from=builder /app/build/target/helper /app/helper

CMD ["/app/helper"]
