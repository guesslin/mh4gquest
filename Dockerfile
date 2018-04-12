FROM golang:1.10
MAINTAINER guesslin1986@gmail.com

RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/guesslin/mh4gquest/vendor
COPY vendor/vendor.json /go/src/github.com/guesslin/mh4gquest/vendor/
WORKDIR /go/src/github.com/guesslin/mh4gquest
RUN govendor sync

ADD . /go/src/github.com/guesslin/mh4gquest

RUN go test ./...

RUN go install

EXPOSE 8080

CMD mh4gquest -http
