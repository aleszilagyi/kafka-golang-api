FROM golang:1.19-alpine
WORKDIR /go/src

ENV PATH="go/bin:${PATH}"

RUN apk update && apk add build-base librdkafka-dev  \
    && go install github.com/golang/mock/mockgen@v1.5.0

CMD ["tail", "-f", "/dev/null"]