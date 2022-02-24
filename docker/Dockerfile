FROM golang:1.17-alpine

WORKDIR /go/src/app

COPY . .

RUN apk update  && \
  apk add git && \
  go get github.com/oxequa/realize && \
  go build


EXPOSE 3000