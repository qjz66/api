FROM golang:alpine AS builder
LABEL authors="qjz"

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o test .

FROM debian:bullseye-slim

COPY ./wait-for.sh /

COPY ./config.json /
COPY ./file /file

COPY --from=builder /build/test /

#RUN sed -i 's/deb.debian.org/archive.debian.org/' /etc/apt/sources.list

RUN set -eux; \
	apt-get update; \
	#apt-get install -y \
    DEBIAN_FRONTEND=noninteractive;
RUN apt-get install -y \
		--no-install-recommends \
        #netcat-traditional\
		netcat; 
RUN chmod 755 wait-for.sh
ENTRYPOINT ["./test"]