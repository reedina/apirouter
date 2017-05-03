FROM golang:1.7.5-alpine
RUN mkdir -p /go/src/github.com/mikerap/apirouter
ADD . /go/src/github.com/mikerap/apirouter
RUN apk update && apk upgrade && apk add git
RUN go install github.com/mikerap/apirouter
