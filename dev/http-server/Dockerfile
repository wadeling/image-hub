FROM golang:alpine3.16 as builder
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
WORKDIR /go/src/github.com/wadeling/http-server
COPY . ./
RUN go build -v -o server .

FROM alpine:3.16
COPY --from=builder /go/src/github.com/wadeling/http-server/server /usr/local/bin
CMD ["server"]
