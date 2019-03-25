FROM golang:1.10.2-alpine3.7 AS build
RUN apk --no-cache add gcc g++ make ca-certificates
WORKDIR /go/src/github.com/tattsum/meower

COPY Gopkg.lock Gopkg.toml ./
COPY vender vender
COPY util util
COPY event event
COPY db db
COPY search search
COPY schema schema
COPY meow-service meow-service
COPY query-service query-service
COPY pusher-service pusher-service
COPY retry retry

RUN go install ./...

FROM alpine:3.7
WORKDIR /usr/bin
COPY --form=build /go/bin
