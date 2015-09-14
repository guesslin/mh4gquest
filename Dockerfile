FROM golang:1.5-onbuild
MAINTAINER guesslin1986@gmail.com

EXPOSE 8080

ENTRYPOINT ["/go/bin/app", "-http"]
