FROM golang:1.18-alpine AS build-stage

ENV CGO_ENABLED 0
ENV GOOS linux

WORKDIR /go/src/go-server-exmaple/
COPY . .

RUN go mod download
RUN go build -a -installsuffix cgo -o server cmd/server/server.go  

FROM alpine:latest AS production-stage

ENV TZ=Asia/Bangkok

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /root/
COPY --from=build-stage  /go/src/go-server-exmaple/server /bin/server

ENTRYPOINT [ "server" ]