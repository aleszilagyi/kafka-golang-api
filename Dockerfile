FROM golang:1.17
WORKDIR /go/src

ENV PATH="go/bin:${PATH}"

RUN apt update && apt install build-essential librdkafka-dev -y \
    go install github.com/golang/mock/mockgen@v1.5.0

CMD ["tail", "-f", "/dev/null"]