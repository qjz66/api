FROM ubuntu:latest
LABEL authors="qjz"

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o api .

FROM scratch

COPY ./config.json /config.json
COPY ./file /file

COPY --from=builder /build/api /

ENTRYPOINT ["api", "-b"]