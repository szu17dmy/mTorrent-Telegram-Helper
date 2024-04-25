FROM golang:1.21.9-bookworm AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

ADD / /app

RUN go build -o ./build/target/helper github.com/szu17dmy/mtorrent-telegram-helper/cmd

FROM golang:1.21.9-bookworm

ENV DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Shanghai
RUN ln -fs /usr/share/zoneinfo/$TZ /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

WORKDIR /app
COPY --from=builder /app/build/target/helper /app/helper

CMD ["/app/helper"]
