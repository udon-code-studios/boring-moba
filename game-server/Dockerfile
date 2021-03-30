# file: Dockerfile
#

FROM golang:1.13.8

ADD src src

# install non-standard go packages
RUN go get github.com/gorilla/websocket
RUN go get github.com/antonholmquist/jason

RUN go build -o main src/*.go

CMD ["./main"]

EXPOSE 80

#
# end of file
