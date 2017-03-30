FROM golang:1.8

MAINTAINER Christopher EnyTC <chris@enytc.com>

RUN mkdir -p /go/src/github.com/chrisenytc/ullli

WORKDIR /go/src/github.com/chrisenytc/ullli

COPY . $APP_HOME

RUN script/install
