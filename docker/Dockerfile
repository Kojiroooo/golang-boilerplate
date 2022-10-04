FROM golang:1.17-alpine

WORKDIR /go/src/app

COPY . .

RUN apk update  && \
  apk add git gcc libc-dev && \
  go get github.com/oxequa/realize && \
  go get -u gorm.io/gorm && \
  go get -u gorm.io/driver/sqlite && \
  go build

EXPOSE 3000
