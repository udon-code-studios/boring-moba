# file: Dockerfile
#

FROM golang:1.13.8

ADD src src

RUN go get github.com/gorilla/websocket 
RUN go build -o main src/*.go

CMD ["./main"]

EXPOSE 80

#
# end of file
