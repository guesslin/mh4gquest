FROM golang:1.10-alpine AS builder
MAINTAINER guesslin1986@gmail.com
RUN apk --no-cache add ca-certificates && update-ca-certificates && apk --no-cache add git

RUN go get -u github.com/kardianos/govendor

RUN mkdir -p /go/src/github.com/guesslin/mh4gquest/vendor
COPY ./vendor /go/src/github.com/guesslin/mh4gquest/vendor
WORKDIR /go/src/github.com/guesslin/mh4gquest
RUN govendor sync -v
ADD . /go/src/github.com/guesslin/mh4gquest
RUN go test -cover -failfast ./...
RUN go build -v

FROM golang:1.10-alpine AS final
RUN apk --no-cache add ca-certificates && update-ca-certificates
RUN mkdir -p /opt/app
WORKDIR /opt/app
COPY --from=builder /go/src/github.com/guesslin/mh4gquest/mh4gquest /opt/app/
COPY ./quests.json /opt/app/
ENV SRV=/opt/app/mh4gquest

EXPOSE 8080

CMD ${SRV} -http
