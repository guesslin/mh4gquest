FROM golang:1.10 AS builder
MAINTAINER guesslin1986@gmail.com

RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/guesslin/mh4gquest/vendor
COPY ./vendor /go/src/github.com/guesslin/mh4gquest/vendor
WORKDIR /go/src/github.com/guesslin/mh4gquest
RUN govendor sync -v
ADD . /go/src/github.com/guesslin/mh4gquest
RUN go test ./...
RUN go build -v

FROM alpine:3.6 AS final
RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN mkdir -p /opt/app
WORKDIR /opt/app
COPY --from=builder /go/src/github.com/guesslin/mh4gquest/mh4gquest /opt/app/

EXPOSE 8080

CMD ["/opt/app/mh4gquest", "-http"]
